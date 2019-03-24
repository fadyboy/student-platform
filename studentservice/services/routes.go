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
		endpoints.GetStudent,
	},
	Route{
		"CreateStudent",
		"POST",
		"/students",
		endpoints.CreateStudent,
	},
	Route{
		"ListStudents",
		"GET",
		"/students",
		endpoints.ListAllStudents,
	},
}
