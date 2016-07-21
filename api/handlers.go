package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"Projeto-api/db"
)

/*
title: get courses
path: /courses
method: GET
produce: application/json
response:
	200: list courses
	204: No Content
	400: Invalid data
	404: Not found
	500: Internal error
*/
func GetCourses(w http.ResponseWriter, r *http.Request) {
	typ := r.FormValue("type")
	course := r.FormValue("course")

	if typ == "" || course == "" {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	courses, err := db.GetCourses(typ, course, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if len(courses) == 0 {
		http.Error(w, "No Content", http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(courses); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
title: get questions
path: /courses/questions
method: GET
produce: application/json
response:
	200: list questions
	204: No Content
	400: Invalid data
	404: Not found
	500: Internal error
*/
func GetQuestions(w http.ResponseWriter, r *http.Request) {
	typ := r.FormValue("type")

	if typ == "" {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	questions, err := db.GetQuestions(typ, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if len(questions) == 0 {
		http.Error(w, "No Content", http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(questions); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

/*
title: create user
path: /users
method: POST
response:
	201: user created
	400: Invalid data
	409: User already exists
*/
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := db.User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
		Created:  time.Now(),
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err := user.CreateUser("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

/*
title: update user
path: /users/update
method: POST
response:
	200: user updated
	400: Invalid data
	404: user not found
*/
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := db.User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	if user.Name == "" || (user.Email == "" && user.Password == "") {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err := user.UpdateUser("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusOK)
}

/*
title: feedback
path: /feedback
method: POST
reponse:
	200: feedback ok
	400: invalid data
	404: course not found
*/
func Feedback(w http.ResponseWriter, r *http.Request) {
	var course db.Courses
	t := r.FormValue("type")
	c := r.FormValue("course")
	vote := r.FormValue("vote")
	name := r.FormValue("name")
	if t == "" || c == "" || vote == "" {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	rate, err := strconv.Atoi(vote)
	if err != nil {
		http.Error(w, "Invalid data, vote must be a integer", http.StatusBadRequest)
		return
	}
	course.Rate = rate
	course.Name = name
	err = course.Feedback(t, c, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusOK)
}
