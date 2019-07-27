package endpoints

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/student-platform/studentservice/dbclient"
	"github.com/student-platform/studentservice/models"
)

// CreateStudent - function that creates a new student
func CreateStudent(res http.ResponseWriter, req *http.Request) {
	// connect to database
	db := dbclient.ConnectToDatabase()
	defer db.Close()

	res.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if err := req.ParseForm(); err != nil {
		http.Error(res, "Bad request. Error processing form data", 400)
		return
	}
	firstname := req.FormValue("firstname")
	lastname := req.PostFormValue("lastname")
	gender := req.PostFormValue("gender")
	dateOfBirth, err := convertDateOfBirth(req.PostFormValue("dateofbirth"))
	if err != nil {
		http.Error(res, "Error processing date of birth", 400)
		return
	}

	parentorguardian := req.PostFormValue("parentorguardian")
	address := req.PostFormValue("address")
	contactNumber := req.PostFormValue("contactnumber")

	student := &models.Student{
		Firstname:        firstname,
		Lastname:         lastname,
		Gender:           gender,
		DateOfBirth:      dateOfBirth,
		ParentOrGuardian: parentorguardian,
		Address:          address,
		ContactNumber:    contactNumber,
	}
	db.Save(&student)

	json.NewEncoder(res).Encode(student)

}

// Helper function to convert date of birth which is a string to a time.Time type
func convertDateOfBirth(dateOfBirth string) (time.Time, error) {
	convertedBirthDate, err := time.Parse("2006-01-02", dateOfBirth)
	if err != nil {
		log.Printf("Error converting date of birth - %s", err)
	}

	return convertedBirthDate, err
}
