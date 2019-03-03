package main

import (
	"fmt"

	"github.com/student-plaform/studentservice/services"
)

func main() {
	fmt.Println("Starting the student service")
	services.StartWebserver("8001")
}
