package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/student-platform/studentservice/dbclient"
	"github.com/student-platform/studentservice/models"
)

// CreateStudent - function that creates a new student
func CreateStudent(res http.ResponseWriter, req *http.Request) {
	// connect to database
	db := dbclient.ConnectToDatabase()
	defer db.Close()

	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	studentDetails := &models.Student{}
	err := json.NewDecoder(req.Body).Decode(&studentDetails)
	if err != nil {
		http.Error(res, "Bad request. Please check parameters of request", 400)
	}
	firstname := studentDetails.Firstname
	lastname := studentDetails.Lastname
	gender := studentDetails.Gender
	parentorguardian := studentDetails.ParentOrGuardian
	address := studentDetails.Address
	contactNumber := studentDetails.ContactNumber

	// vars := mux.Vars(req)
	// firstname := vars["firsname"]
	// fmt.Println("name ", firstname)
	// lastname := vars["lastname"]
	// gender := vars["gender"]
	// dateOfBirth, _ := date.Parse(vars["dateofbirth"], "yyyy-mm-dd")
	// parentGuardian := vars["parentorguardian"]
	// address := vars["address"]
	// contactNumber, _ := strconv.ParseInt(vars["contactnumber"], 10, 64)

	db.Create(&models.Student{
		Firstname: firstname,
		Lastname:  lastname,
		Gender:    gender,
		// DateOfBirth:      dateOfBirth,
		ParentOrGuardian: parentorguardian,
		Address:          address,
		ContactNumber:    contactNumber,
	})

	json.NewEncoder(res).Encode(studentDetails)

}
