package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net"

	dpb "github.com/TekClinic/Doctors-MicroService/doctors_protobuf"
	ms "github.com/TekClinic/MicroService-Lib"
	"github.com/go-playground/validator/v10"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// doctorsServer is an implementation of GRPC Doctor microservice. It provides access to a database via db field.
type doctorsServer struct {
	dpb.UnimplementedDoctorsServiceServer
	ms.BaseServiceServer
	db *bun.DB
	// use a single instance of Validate, it caches struct info
	validate *validator.Validate
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

// GetDoctor returns a Doctor that corresponds to the given id.
// Requires authentication. If authentication is not valid, codes.Unauthenticated is returned.
// Requires an admin role. If roles are insufficient, codes.PermissionDenied is returned.
// If a Doctor with a given id doesn't exist, codes.NotFound is returned.
func (server doctorsServer) GetDoctor(ctx context.Context, req *dpb.GetDoctorRequest) (*dpb.GetDoctorResponse, error) {
	claims, err := server.VerifyToken(ctx, req.GetToken())
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if !claims.HasRole("admin") {
		return nil, status.Error(codes.PermissionDenied, permissionDeniedMessage)
	}

	doctor := new(Doctor)
	err = server.db.NewSelect().Model(doctor).Where("? = ?", bun.Ident("id"), req.GetId()).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "Doctor is not found")
		}
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to fetch a doctor by id: %w", err).Error())
	}
	return &dpb.GetDoctorResponse{
		Doctor: doctor.toGRPC(),
	}, nil
}

// GetDoctorsIDs returns a list of doctors' ids with given filters and pagination.
// Requires authentication. If authentication is not valid, codes.Unauthenticated is returned.
// Requires an admin role. If roles are insufficient, codes.PermissionDenied is returned.
// Offset value is used for pagination. Required be a non-negative value.
// Limit value is used for pagination. Required to be a positive value.
func (server doctorsServer) GetDoctorsIDs(ctx context.Context,
	req *dpb.GetDoctorsIDsRequest) (*dpb.GetDoctorsIDsResponse, error) {
	claims, err := server.VerifyToken(ctx, req.GetToken())
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if !claims.HasRole("admin") {
		return nil, status.Error(codes.PermissionDenied, permissionDeniedMessage)
	}

	if req.GetOffset() < 0 {
		return nil, status.Error(codes.InvalidArgument, "offset has to be a non-negative integer")
	}
	if req.GetLimit() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "limit has to be a positive integer")
	}
	if req.GetLimit() > maxPaginationLimit {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("maximum allowed limit values is %d", maxPaginationLimit))
	}

	var ids []int32
	baseQuery := server.db.NewSelect().Model((*Doctor)(nil)).Column("id")

	if req.GetSearch() != "" {
		// Postgres specific code. Use full-text search to search for patients.
		baseQuery = baseQuery.
			TableExpr("replace(websearch_to_tsquery('simple', ?)::text || ' ',''' ',''':*') query", req.GetSearch()).
			Where("text_searchable @@ query::tsquery").
			OrderExpr("ts_rank(text_searchable, query::tsquery) DESC")
	}

	err = baseQuery.
		Offset(int(req.GetOffset())).
		Limit(int(req.GetLimit())).
		Scan(ctx, &ids)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to fetch doctors: %w", err).Error())
	}
	count, err := baseQuery.Count(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to count doctors: %w", err).Error())
	}

	return &dpb.GetDoctorsIDsResponse{
		Count:   int32(count),
		Results: ids,
	}, nil
}

// CreateDoctor creates a doctor with the given specifications.
// Requires authentication. If authentication is not valid, codes.Unauthenticated is returned.
// Requires an admin role. If roles are not sufficient, codes.PermissionDenied is returned.
// If some argument is missing or not valid, codes.InvalidArgument is returned.
func (server doctorsServer) CreateDoctor(ctx context.Context,
	req *dpb.CreateDoctorRequest) (*dpb.CreateDoctorResponse, error) {
	claims, err := server.VerifyToken(ctx, req.GetToken())
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if !claims.HasRole("admin") {
		return nil, status.Error(codes.PermissionDenied, permissionDeniedMessage)
	}

	doctor := Doctor{
		Active:       true,
		Name:         req.GetName(),
		Gender:       req.GetGender(),
		PhoneNumber:  req.GetPhoneNumber(),
		Specialities: req.GetSpecialities(),
		SpecialNote:  req.GetSpecialNote(),
	}
	if err = server.validate.Struct(doctor); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, err = server.db.NewInsert().Model(&doctor).Exec(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to create a doctor: %w", err).Error())
	}

	return &dpb.CreateDoctorResponse{Id: doctor.ID}, nil
}

// DeleteDoctor deletes a doctor with the given id.
// Requires authentication. If authentication is not valid, codes.Unauthenticated is returned.
// Requires an admin role. If roles are not sufficient, codes.PermissionDenied is returned.
// If the doctor with the given id doesn't exist, codes.NotFound is returned.
func (server doctorsServer) DeleteDoctor(ctx context.Context, req *dpb.DeleteDoctorRequest) (
	*dpb.DeleteDoctorResponse, error) {
	claims, err := server.VerifyToken(ctx, req.GetToken())
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if !claims.HasRole("admin") {
		return nil, status.Error(codes.PermissionDenied, permissionDeniedMessage)
	}

	res, err := server.db.NewDelete().Model((*Doctor)(nil)).Where("id = ?", req.GetId()).Exec(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to delete the doctor: %w", err).Error())
	}
	// if db supports affected rows count and no rows were affected, return not found
	rows, err := res.RowsAffected()
	if err == nil && rows == 0 {
		return nil, status.Error(codes.NotFound, "doctor is not found")
	}
	return &dpb.DeleteDoctorResponse{}, nil
}

// createDoctorsServer initializes a doctorsServer with all the necessary fields.
func createDoctorsServer() (*doctorsServer, error) {
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
	return &doctorsServer{
		BaseServiceServer: base,
		db:                db,
		validate:          validator.New(validator.WithRequiredStructEnabled())}, nil
}

func main() {
	service, err := createDoctorsServer()
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
	dpb.RegisterDoctorsServiceServer(srv, service)

	log.Println("Server listening on :" + service.GetPort())
	if err = srv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
