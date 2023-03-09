package application

import (
	"fmt"
	"time"
)

func Monitor(sites []string, tries int, delay int) {
	for i := 1; i <= tries; i++ {
		fmt.Println("Try number", i)
		monitorSites(sites, delay)
	}
}

func monitorSites(sites []string, delay int) {
	for _, site := range sites {
		fmt.Println("Monitoring site:", site)
	}
	fmt.Println("")

	time.Sleep(time.Duration(delay) * time.Second)
}
