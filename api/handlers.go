package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tcc-denis-raul/Projeto-api/db"
)

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "HELLO")
}

/*
title: get type courses
path: /types/courses
method: GET
produce: application/json
response:
	200: list types
	204: No Content
	404: Not found
	500: Internal error
*/

func GetTypeCourses(c *gin.Context) {
	typesCourses, err := db.GetTypeCourses()
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	if len(typesCourses) == 0 {
		c.String(http.StatusNoContent, "No Content")
		return
	}
	c.JSON(http.StatusOK, typesCourses)
	return
}

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
func GetCourses(c *gin.Context) {
	leng, err := strconv.Atoi(c.Query("length"))
	if err != nil {
		leng = 0
	}
	filter := db.Filter{
		Type:     c.Query("type"),
		Course:   c.Query("course"),
		Based:    c.Query("based"),
		Dynamic:  c.Query("dynamic"),
		Platform: c.Query("platform"),
		Extra:    c.Query("extra"),
		Price:    c.Query("price"),
		Length:   leng,
	}

	if filter.Type == "" || filter.Course == "" {
		c.String(http.StatusBadRequest, "Invalid data")
		return
	}

	courses, err := filter.GetCourses()
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	if len(courses) == 0 {
		c.String(http.StatusNoContent, "No Content")
		return
	}
	c.JSON(http.StatusOK, courses)
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
func GetQuestions(c *gin.Context) {
	typ := c.Query("type")

	if typ == "" {
		c.String(http.StatusBadRequest, "Invalid data")
		return
	}

	questions, err := db.GetQuestions(typ)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	if len(questions) == 0 {
		c.String(http.StatusNoContent, "No Content")
		return
	}
	c.JSON(http.StatusOK, questions)

}

/*
title: details of course
path: /course/detail
method: GET
produce: application/json
response:
	200: detail course
	400: Invalid data
	404: Not found
	500: Internal error

*/
func GetDetailCourse(c *gin.Context) {
	filter := db.CourseDetail{
		Type:   c.Query("type"),
		Course: c.Query("course"),
		Name:   c.Query("name"),
	}
	if filter.Type == "" || filter.Course == "" || filter.Name == "" {
		c.String(http.StatusBadRequest, "Invalid data")
		return
	}
	course, err := filter.GetDetailCourse()
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, course)
}

/*
title: get user
path: /users
method: GET
response:
	200: information about a user
	400: Invalid data
	409: User already exists
*/

func GetUser(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.String(http.StatusBadRequest, "Invalid data")
		return
	}
	user := db.User{
		Email: email,
	}
	u, err := user.GetUser()
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, u)
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
func CreateUser(c *gin.Context) {
	user := db.User{
		FirstName: c.Query("firtname"),
		LastName:  c.Query("lastname"),
		Email:     c.Query("email"),
		UserName:  c.Query("username"),
		Created:   time.Now(),
	}

	if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.UserName == "" {
		c.String(http.StatusBadRequest, "Invalid data")
		return
	}

	err := user.CreateUser()
	if err != nil {
		c.String(http.StatusConflict, err.Error())
		return
	}
	c.String(http.StatusOK, "")
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
func UpdateUser(c *gin.Context) {
	user := db.User{
		FirstName: c.Query("firtname"),
		LastName:  c.Query("lastname"),
		Email:     c.Query("email"),
		UserName:  c.Query("username"),
	}
	if user.UserName == "" || user.FirstName == "" && user.LastName == "" && user.Email == "" {
		c.String(http.StatusBadRequest, "Invalid data")
		return
	}

	err := user.UpdateUser()
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, "")
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
func Feedback(c *gin.Context) {
	var course db.Courses
	typ := c.Query("type")
	cou := c.Query("course")
	vote := c.Query("vote")
	name := c.Query("name")
	if typ == "" || cou == "" || vote == "" {
		c.String(http.StatusBadRequest, "Invalid data")
		return
	}
	rate, err := strconv.ParseFloat(vote, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "rate must be a string")
		return
	}
	course.Rate = rate
	course.Name = name
	err = course.Feedback(typ, cou)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, "")
}

/*
title: indicate course
path: /course/indicate
method: POST
response:
	200: indicate ok
	400: invalid data
	409: course already exists
*/
func IndicateCourse(c *gin.Context) {
	indication := db.IndicateCourse{
		Type:   c.Query("type"),
		Course: c.Query("course"),
		Name:   c.Query("name"),
		Url:    c.Query("url"),
	}
	if indication.Type == "" || indication.Course == "" ||
		indication.Name == "" || indication.Url == "" {
		c.String(http.StatusBadRequest, "Invalid data")
		return
	}
	err := indication.IndicateCourse()
	if err != nil {
		c.String(http.StatusConflict, err.Error())
		return
	}
	c.String(http.StatusOK, "")
}

/*
title: save user preferences
path: /users/profile
method: POST
response:
	200: prefereces saved
	400: invalid data
	500: internal error
*/
func SaveUserPreferences(c *gin.Context) {
	preferences := db.UserPreferences{
		UserName: c.Query("username"),
		Based:    c.Query("based"),
		Dynamic:  c.Query("dynamic"),
		Platform: c.Query("platform"),
		Extra:    c.Query("extra"),
		Price:    c.Query("price"),
	}
	if preferences.UserName == "" {
		c.String(http.StatusBadRequest, "Invalid data")
		return
	}
	err := preferences.SaveUserPreferences()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, "")
}

func GetUserPreferences(c *gin.Context) {
	user := db.UserPreferences{
		UserName: c.Query("username"),
	}
	if user.UserName == "" {
		c.String(http.StatusBadRequest, "Invalid data")
		return
	}
	u, err := user.GetUserPreferences()
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, u)
}
