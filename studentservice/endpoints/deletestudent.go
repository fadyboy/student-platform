package endpoints

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/student-platform/studentservice/dbclient"
	"github.com/student-platform/studentservice/models"
)

// DeleteStudent - Method to delete student record
func DeleteStudent(res http.ResponseWriter, req *http.Request) {
	// Connect to db
	db := dbclient.ConnectToDatabase()
	defer db.Close()

	var student models.Student
	studentID := mux.Vars(req)["studentId"]
	// Check if student actually exists
	if db.Where("ID = ?", studentID).Find(&student).RecordNotFound() {
		http.Error(res, "Record not found", 404)
		return
	}

	db.Delete(&student)
	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	res.WriteHeader(http.StatusOK)
}
