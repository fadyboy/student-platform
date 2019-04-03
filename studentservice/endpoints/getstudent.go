package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/student-platform/studentservice/dbclient"
	"github.com/student-platform/studentservice/models"
)

// GetStudent - handler function to get a particular student from the student table
func GetStudent(res http.ResponseWriter, req *http.Request) {
	// connect to database
	db := dbclient.ConnectToDatabase()
	defer db.Close()

	// Get studentId from request param
	studentID := mux.Vars(req)["studentId"]
	var student models.Student
	// check if record found and return appropraite response
	if db.Where("ID = ?", studentID).Find(&student).RecordNotFound() {
		http.Error(res, "No record found", 404)
		return
	}
	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(student)
}
