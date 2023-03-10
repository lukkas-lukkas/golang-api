package application

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func Monitor(sites []string, tries int, delay int) {
	for i := 1; i <= tries; i++ {
		fmt.Println("Try number", i)
		monitorSites(sites, delay)
	}
}

func monitorSites(sites []string, delay int) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln("error to create file:", err)
	}

	for _, site := range sites {
		_, err := http.Get(site)

		if err != nil {
			fmt.Println("error to read site:", site, "Error:", err)
		}

		file.WriteString("site: " + site)
	}
	fmt.Println("")

	time.Sleep(time.Duration(delay) * time.Second)
}
