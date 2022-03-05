package controllers

import (
	"context"
	"log"
	"os"
	"testing"

	"kora.com/project/src/database"
)

func TestMain(m *testing.M) {
	client := database.DBConnect()
	code := m.Run()
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(code)
}
