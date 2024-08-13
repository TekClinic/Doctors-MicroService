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

// fromGRPC creates a Doctor from a GRPC version.
func fromGRPC(d *dpb.Doctor) Doctor {
	return Doctor{
		ID:           d.GetId(),
		Active:       d.GetActive(),
		Name:         d.GetName(),
		Gender:       d.GetGender(),
		PhoneNumber:  d.GetPhoneNumber(),
		Specialities: d.GetSpecialities(),
		SpecialNote:  d.GetSpecialNote(),
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

	// Postgres specific code. Add a text_searchable column for full-text search.
	_, err := db.NewRaw(
		"CREATE OR REPLACE FUNCTION immutable_array_to_string(text[]) " +
			"RETURNS text as $$ SELECT array_to_string($1, ','); $$ LANGUAGE sql IMMUTABLE;" +
			"ALTER TABLE doctors " +
			"ADD COLUMN IF NOT EXISTS text_searchable tsvector " +
			"GENERATED ALWAYS AS " +
			"(" +
			"setweight(to_tsvector('simple', coalesce(phone_number, '')), 'A')   || " +
			"setweight(to_tsvector('simple', coalesce(name, '')), 'B')           || " +
			"setweight(to_tsvector('simple', immutable_array_to_string(coalesce(specialities, '{}'))), 'C')   || " +
			"setweight(to_tsvector('simple', coalesce(special_note, '')), 'C')" +
			") STORED").Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
