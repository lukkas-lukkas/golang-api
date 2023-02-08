package main

import (
	"fmt"
	"github.com/lukkas-lukkas/go-api-rest/src/domain"
)

func main() {
	doctor1 := domain.NewDoctor("Lucas Lima")

	doctor2 := domain.NewDoctor("Lucas Oliveira Lima")

	list := domain.DoctorList{}

	list.Add(*doctor1)
	list.Add(*doctor2)

	fmt.Println(list)
}
