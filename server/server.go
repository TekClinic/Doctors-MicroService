package main

import (
	"context"
	"database/sql"
	"fmt"
	ms "github.com/TekClinic/MicroService-Lib"
	ppb "github.com/TekClinic/PHR_Doctors_MicroService/doctors_protobuf"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

// doctorsServer is an implementation of GRPC doctor microservice. It provides access to database via db field
type doctorsServer struct {
	ppb.UnimplementedDoctorServiceServer
	ms.BaseServiceServer
	db *bun.DB
}

const (
	envDBAddress     = "DB_ADDR"
	envDBUser        = "DB_USER"
	envDBDatabase    = "DB_DATABASE"
	envDBPassword    = "DB_PASSWORD"
	envBunDebugLevel = "BUN_DEBUG"

	applicationName = "doctors"

	permissionDeniedMessage = "You don't have enough permission to access this resource"

	maxPaginationLimit = 50
)

// Getdoctor returns a doctor that corresponds to the given id
// Requires authentication. If authentication is not valid, codes.Unauthenticated is returned
// Requires admin role. If roles is not sufficient, codes.PermissionDenied is returned
// If doctor with a given id doesn't exist, codes.NotFound is returned
func (server doctorsServer) Getdoctor(ctx context.Context, req *ppb.DoctorRequest) (*ppb.Doctor, error) {
	claims, err := server.VerifyToken(ctx, req.GetToken())
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if !claims.HasRole("admin") {
		return nil, status.Error(codes.PermissionDenied, permissionDeniedMessage)
	}

	doctor := new(doctor)
	err = server.db.NewSelect().Model(doctor).Where("? = ?", bun.Ident("id"), req.Id).Scan(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to fetch a user by id: %w", err).Error())
	}
	if doctor == nil {
		return nil, status.Error(codes.NotFound, "User is not found")
	}
	return doctor.toGRPC(), nil
}

// GetdoctorsIds returns a list of doctors' ids with given filters and pagination
// Requires authentication. If authentication is not valid, codes.Unauthenticated is returned
// Requires admin role. If roles is not sufficient, codes.PermissionDenied is returned
// Offset value is used for a pagination. Required be a non-negative value
// Limit value is used for a pagination. Required to be a positive value
func (server doctorsServer) GetdoctorsIds(ctx context.Context, req *ppb.DoctorsRequest) (*ppb.PaginatedResponse, error) {
	claims, err := server.VerifyToken(ctx, req.GetToken())
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if !claims.HasRole("admin") {
		return nil, status.Error(codes.PermissionDenied, permissionDeniedMessage)
	}

	if req.Offset < 0 {
		return nil, status.Error(codes.InvalidArgument, "offset has to be a non-negative integer")
	}
	if req.Limit <= 0 {
		return nil, status.Error(codes.InvalidArgument, "limit has to be a positive integer")
	}
	if req.Limit > maxPaginationLimit {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("maximum allowed limit values is %d", maxPaginationLimit))
	}

	var ids []int32
	baseQuery := server.db.NewSelect().Model((*doctor)(nil)).Column("id")
	err = baseQuery.
		Offset(int(req.Offset)).
		Limit(int(req.Limit)).
		Scan(ctx, &ids)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to fetch users: %w", err).Error())
	}
	count, err := baseQuery.Count(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to count users: %w", err).Error())
	}

	return &ppb.PaginatedResponse{
		Count:   int32(count),
		Results: ids,
	}, nil
}

// createdoctorsServer initializes a doctorsServer with all the necessary fields.
func createdoctorsServer() (*doctorsServer, error) {
	base, err := ms.CreateBaseServiceServer()
	if err != nil {
		return nil, err
	}
	addr, err := ms.GetRequiredEnv(envDBAddress)
	if err != nil {
		return nil, err
	}
	user, err := ms.GetRequiredEnv(envDBUser)
	if err != nil {
		return nil, err
	}
	password, err := ms.GetRequiredEnv(envDBPassword)
	if err != nil {
		return nil, err
	}
	database, err := ms.GetRequiredEnv(envDBDatabase)
	if err != nil {
		return nil, err
	}
	connector := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(addr),
		pgdriver.WithUser(user),
		pgdriver.WithPassword(password),
		pgdriver.WithDatabase(database),
		pgdriver.WithApplicationName(applicationName),
		pgdriver.WithInsecure(true),
	)
	db := bun.NewDB(sql.OpenDB(connector), pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv(envBunDebugLevel),
	))
	return &doctorsServer{BaseServiceServer: base, db: db}, nil
}

func main() {
	service, err := createdoctorsServer()
	if err != nil {
		log.Fatal(err)
	}

	err = createSchemaIfNotExists(context.Background(), service.db)
	if err != nil {
		log.Fatal(err)
	}

	listen, err := net.Listen("tcp", ":"+service.GetPort())
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	ppb.RegisterDoctorServiceServer(srv, service)

	log.Println("Server listening on :" + service.GetPort())
	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
