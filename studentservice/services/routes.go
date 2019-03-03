package services

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

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
}
