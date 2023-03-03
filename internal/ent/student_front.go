package ent

import (
	"github.com/google/uuid"
	"github.com/naofel1/api-golang-template/internal/primitive"
)

// StudentsFront is a parsable slice of StudentFront.
type StudentsFront []*StudentFront

// StudentDisplayFront is the model entity for the Student schema.
type StudentDisplayFront struct {
	FirstName string           `json:"firstname,omitempty"`
	LastName  string           `json:"lastname,omitempty"`
	Pseudo    string           `json:"pseudo,omitempty"`
	Gender    primitive.Gender `json:"gender,omitempty"`
	Birthday  string           `json:"birthday,omitempty"`
}

// StudentFront is the model entity returned to frontend
type StudentFront struct {
	Display      *StudentDisplayFront `json:"display,omitempty"`
	PasswordHash []byte               `json:"-"`
	ID           uuid.UUID            `json:"id,omitempty"`
}

// ToFront is the structure returned to the frontend
func (s *Student) ToFront() *StudentFront {
	return &StudentFront{
		ID: s.ID,
		Display: &StudentDisplayFront{
			FirstName: s.FirstName,
			LastName:  s.LastName,
			Birthday:  s.Birthday.Format("2006-01-02"),
			Pseudo:    s.Pseudo,
			Gender:    s.Gender,
		},
	}
}

// ToFront is the structure returned to the frontend
func (s Students) ToFront() []*StudentFront {
	domainStudents := make([]*StudentFront, 0, len(s))

	for _, stud := range s {
		domainStudents = append(domainStudents, stud.ToFront())
	}

	return domainStudents
}
