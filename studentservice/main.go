package main

import (
	"fmt"

	"github.com/student-platform/studentservice/dbclient"
	"github.com/student-platform/studentservice/services"
)

func main() {
	fmt.Println("Starting the student service")
	dbclient.InitializeDatabase()
	services.StartWebserver("8001")
}
