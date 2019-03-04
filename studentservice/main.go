package main

import (
	"fmt"

	"github.com/student-plaform/studentservice/dbclient"
	"github.com/student-plaform/studentservice/services"
)

func main() {
	fmt.Println("Starting the student service")
	dbclient.InitializeDatabase()
	services.StartWebserver("8001")
}
