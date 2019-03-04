package endpoints

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fxtlabs/date"
	"github.com/gorilla/mux"

	"github.com/student-plaform/studentservice/dbclient"
	"github.com/student-plaform/studentservice/models"
)

// CreateStudent - function that creates a new student
func CreateStudent(res http.ResponseWriter, req *http.Request) {
	// connect to database
	db := dbclient.ConnectToDatabase()
	defer db.Close()

	vars := mux.Vars(req)
	firstname := vars["firstname"]
	lastname := vars["lastname"]
	gender := vars["gender"]
	dateOfBirth, _ := date.Parse(vars["dateofbirth"], "yyyy-mm-dd")
	parentGuardian := vars["parentorguardian"]
	address := vars["address"]
	contactNumber, _ := strconv.ParseInt(vars["contactnumber"], 10, 64)

	db.Create(&models.Student{
		Firstname:        firstname,
		Lastname:         lastname,
		Gender:           gender,
		DateOfBirth:      dateOfBirth,
		ParentOrGuardian: parentGuardian,
		Address:          address,
		ContactNumber:    contactNumber,
	})

	fmt.Fprintf(res, "New student successfully created")

}
