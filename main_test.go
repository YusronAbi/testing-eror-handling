package main

import (
	"inventory-management/database"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	database.ConnectTestDB()
	code := m.Run()
	os.Exit(code)
}
