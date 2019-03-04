package services

import (
	"net/http"

	"github.com/student-plaform/studentservice/endpoints"
)

// Route - struct for the route object
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes - type for routes
type Routes []Route

var routes = Routes{
	Route{
		"GetStudent",
		"GET",
		"/students/{studentId}",
		func(res http.ResponseWriter, req *http.Request) {
			res.Header().Set("Content-Type", "application/json; charset=UTF-8")
			res.Write([]byte("{\"status\":\"success\"}"))
		},
	},
	Route{
		"CreateStudent",
		"POST",
		"/students",
		endpoints.CreateStudent,
	},
}
