package main

import (
	"context"

	dpb "github.com/TekClinic/Doctors-MicroService/doctors_protobuf"
	"github.com/uptrace/bun"
)

// Doctor defines a schema of doctors.
type Doctor struct {
	ID           int32             `bun:",pk,autoincrement"`
	Active       bool              ``
	Name         string            `validate:"required,min=1,max=100"`
	Gender       dpb.Doctor_Gender ``
	PhoneNumber  string            `validate:"required,e164"`
	Specialities []string          `bun:",array" validate:"max=30"`
	SpecialNote  string            `validate:"max=500"`
}

// toGRPC returns a GRPC version of Doctor.
func (doctor Doctor) toGRPC() *dpb.Doctor {
	return &dpb.Doctor{
		Id:           doctor.ID,
		Active:       doctor.Active,
		Name:         doctor.Name,
		Gender:       doctor.Gender,
		PhoneNumber:  doctor.PhoneNumber,
		Specialities: doctor.Specialities,
		SpecialNote:  doctor.SpecialNote,
	}
}

// createSchemaIfNotExists creates all required schemas for Doctor microservice.
func createSchemaIfNotExists(ctx context.Context, db *bun.DB) error {
	models := []interface{}{
		(*Doctor)(nil),
	}

	for _, model := range models {
		if _, err := db.NewCreateTable().IfNotExists().Model(model).Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}
