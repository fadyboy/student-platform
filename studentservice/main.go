package main

import (
	"fmt"

<<<<<<< HEAD
	"github.com/student-platform/studentservice/dbclient"
	"github.com/student-platform/studentservice/services"
=======
	"github.com/student-plaform/studentservice/dbclient"
	"github.com/student-plaform/studentservice/services"
>>>>>>> f62371150825e90e133ee51bd9d30eec152e53e6
)

func main() {
	fmt.Println("Starting the student service")
	dbclient.InitializeDatabase()
	services.StartWebserver("8001")
}
