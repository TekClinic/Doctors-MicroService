package main

import (
	"context"
	ppb "github.com/TekClinic/PHR_Doctors_MicroService/doctors_protobuf"
	"github.com/uptrace/bun"
	"time"
)

// PersonalId defines a schema of personal ids
type PersonalId struct {
	ID   string
	Type string
}

// doctor defines a schema of doctors
type doctor struct {
	ID          int32 `bun:",pk,autoincrement"`
	Active      bool
	Name        string
	PersonalId  PersonalId `bun:"embed:personal_id_"`
	Gender      ppb.Doctor_Gender
	PhoneNumber string
	Languages   []string `bun:",array"`
	BirthDate   time.Time
	ReferredBy  string
	SpecialNote string
}

// toGRPC returns a GRPC version of PersonalId
func (personalId PersonalId) toGRPC() *ppb.Doctor_PersonalId {
	return &ppb.Doctor_PersonalId{
		Id:   personalId.ID,
		Type: personalId.Type,
	}
}

// toGRPC returns a GRPC version of doctor
func (doctor doctor) toGRPC() *ppb.Doctor {
	return &ppb.Doctor{
		Id:          doctor.ID,
		Active:      doctor.Active,
		Name:        doctor.Name,
		PersonalId:  doctor.PersonalId.toGRPC(),
		Gender:      doctor.Gender,
		PhoneNumber: doctor.PhoneNumber,
		Languages:   doctor.Languages,
		BirthDate:   doctor.BirthDate.Format(time.DateOnly),
		Age:         int32(time.Now().Year() - doctor.BirthDate.Year()),
		ReferredBy:  doctor.ReferredBy,
		SpecialNote: doctor.SpecialNote,
	}
}

// createSchemaIfNotExists creates all required schemas for doctor microservice
func createSchemaIfNotExists(ctx context.Context, db *bun.DB) error {
	models := []interface{}{
		(*doctor)(nil),
	}

	for _, model := range models {
		if _, err := db.NewCreateTable().IfNotExists().Model(model).Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}
