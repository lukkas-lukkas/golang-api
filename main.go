package main

import "fmt"
import uuid "github.com/satori/go.uuid"

type Doctor struct {
	ID string
	Name string
}

func NewDoctor(name string) *Doctor {
	return &Doctor{
		ID: uuid.NewV4().String(),
		Name: name,
	}
}

type DoctorList struct {
	Doctors []Doctor
}

func (d *DoctorList) Add(doctor Doctor) {
	d.Doctors = append(d.Doctors, doctor)
}

func main() {
	doctor1 := NewDoctor("Lucas Lima")

	doctor2 := NewDoctor("Lucas Oliveira Lima")

	list := DoctorList{}
	
	list.Add(*doctor1)
	list.Add(*doctor2)

	fmt.Println(list)
}
