package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	s := echo.New()
	s.GET("/status", Handler)
	err := s.Start(":8080")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server running")
}

func Handler(ctx echo.Context) error {
	employees, _ := json.Marshal("True")
	err := ctx.String(http.StatusOK, string(employees))
	if err != nil {
		return err
	}
	return nil
}
