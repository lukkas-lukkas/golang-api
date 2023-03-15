package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type FileLogger struct{}

func NewFileLogger() *FileLogger {
	return &FileLogger{}
}

func (fl *FileLogger) Log(log map[string]string) {
	location, _ := time.LoadLocation("America/Sao_Paulo")
	file, err := os.OpenFile("logs.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error to create file:", err)
	}

	data := map[string]interface{}{
		"time":    time.Now().In(location).Format("2006-01-02 15:04:05"),
		"context": log,
	}

	json, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error to parse log")
		return
	}

	file.WriteString(string(json) + "\n")

	defer file.Close()
}
