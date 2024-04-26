package main

import (
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
	err := ctx.String(http.StatusOK, "test")
	if err != nil {
		return err
	}
	return nil
}
