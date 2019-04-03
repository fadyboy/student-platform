package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/student-platform/studentservice/dbclient"
	"github.com/student-platform/studentservice/models"
)

// ListAllStudents - function that lists all students in database
func ListAllStudents(res http.ResponseWriter, req *http.Request) {
	db := dbclient.ConnectToDatabase()
	defer db.Close()

	var students []models.Student
	db.Find(&students)
	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(students)
}
