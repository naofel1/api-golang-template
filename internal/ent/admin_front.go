// Package ent contains the schema for frontend entities.
package ent

import (
	"github.com/google/uuid"
)

// AdminDisplayFront is the model entity for the Admin schema.
type AdminDisplayFront struct {
	// Pseudo holds the value of the "pseudo" field.
	Pseudo string `json:"pseudo"`
}

// AdminFront is the model entity returned to the frontend
type AdminFront struct {
	Display *AdminDisplayFront `json:"display,omitempty"`
	ID      uuid.UUID          `json:"id,omitempty"`
}

// ToFront is the structure returned to the frontend
func (a *Admin) ToFront() *AdminFront {
	return &AdminFront{
		ID: a.ID,
		Display: &AdminDisplayFront{
			Pseudo: a.Pseudo,
		},
	}
}
