package application

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/lukkas-lukkas/go-api-rest/src/domain"
)

type MonitorService struct {
	Logger domain.Logger
}

func NewMonitorService(logger domain.Logger) *MonitorService {
	return &MonitorService{
		Logger: logger,
	}
}

func (ms *MonitorService) Monitor(sites []string, tries int, delay int) {
	for i := 1; i <= tries; i++ {
		fmt.Println("Try number", i)
		ms.monitorSites(sites, delay)
	}
}

func (ms *MonitorService) monitorSites(sites []string, delay int) {
	for _, site := range sites {
		response, err := http.Get(site)

		if err != nil {
			fmt.Println("error to read site:", site, "Error:", err)
		}

		ms.Logger.Log(map[string]string{
			"site":        site,
			"status_code": strconv.Itoa(response.StatusCode),
		})
	}

	time.Sleep(time.Duration(delay) * time.Second)
}
