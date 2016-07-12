package main

import (
	"encoding/json"
	"net/http"

	"github.com/tcc-denis-raul/Projeto-api/db"
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
	course := r.FormValue("course")

	if typ == "" || course == "" {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	questions, err := db.GetQuestions(typ, course, "")
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
