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
	file, err := os.OpenFile("logs.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("error to create file:", err)
	}

	for _, site := range sites {
		response, err := http.Get(site)

		if err != nil {
			fmt.Println("error to read site:", site, "Error:", err)
		}

		time := time.Now().Format("2006-01-02 15:04:05")

		log := fmt.Sprintf("%s: site: %s, status_code: %d\n", time, site, response.StatusCode)

		file.WriteString(log)
	}
	defer file.Close()

	time.Sleep(time.Duration(delay) * time.Second)
}
