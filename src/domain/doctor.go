package domain

import uuid "github.com/satori/go.uuid"

func NewDoctor(name string) *Doctor {
	return &Doctor{
		ID: uuid.NewV4().String(),
		Name: name,
	}
}

type Doctor struct {
	ID string
	Name string
}

type DoctorList struct {
	Doctors []Doctor
}

func (d *DoctorList) Add(doctor Doctor) {
	d.Doctors = append(d.Doctors, doctor)
}