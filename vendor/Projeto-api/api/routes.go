package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"GetCourses", "GET", "/courses", GetCourses},
	Route{"GetQuestions", "GET", "/courses/questions", GetQuestions},
	Route{"IndicateCourse", "POST", "/course/indicate", IndicateCourse},
	Route{"Feedback", "POST", "/course/feedback", Feedback},
	Route{"GetUser", "GET", "/users", GetUser},
	Route{"CreateUser", "POST", "/users", CreateUser},
	Route{"UpdateUser", "POST", "/users/update", UpdateUser},
	Route{"GetTypeCourses", "GET", "/types/courses", GetTypeCourses},
	Route{"IndicateCourse", "POST", "/indicate/course", IndicateCourse},
	Route{"SaveUserPreferences", "POST", "/users/profile", SaveUserPreferences},
}
