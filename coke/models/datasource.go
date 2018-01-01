package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type Datasource struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	Fqdn      string    `json:"fqdn" db:"fqdn"`
	Port      int       `json:"port" db:"port"`
	Type      int       `json:"type" db:"type"`
	Secured   bool      `json:"secured" db:"secured"`
}

// String is not required by pop and may be deleted
func (d Datasource) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Datasources is not required by pop and may be deleted
type Datasources []Datasource

// String is not required by pop and may be deleted
func (d Datasources) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *Datasource) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: d.Name, Name: "Name"},
		&validators.StringIsPresent{Field: d.Fqdn, Name: "Fqdn"},
		&validators.IntIsPresent{Field: d.Port, Name: "Port"},
		&validators.IntIsPresent{Field: d.Type, Name: "Type"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *Datasource) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *Datasource) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
