package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/student-platform/studentservice/dbclient"
	"github.com/student-platform/studentservice/models"
)

// UpdateStudent - function to update student record
func UpdateStudent(res http.ResponseWriter, req *http.Request) {
	// connect to database to get student
	db := dbclient.ConnectToDatabase()
	defer db.Close()

	studentID := mux.Vars(req)["studentId"]
	var student models.Student
	if db.Where("ID = ?", studentID).Find(&student).RecordNotFound() {
		http.Error(res, "Record not found", 404)
		return
	}

	// Decode json and map to student model and save changes
	err := json.NewDecoder(req.Body).Decode(&student)
	if err != nil {
		http.Error(res, "Bad request. Please check parameters in request", 400)
	}
	db.Save(&student)

	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(student)

}
