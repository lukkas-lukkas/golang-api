package main

import "fmt"

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

func main() {
	doctor := Doctor{}
	doctor.ID = "123456"
	doctor.Name = "Lucas Lima"

	list := DoctorList{}
	list.Add(doctor)

	fmt.Println(list)
}
