package db

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Courses struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string
	Based       []string
	PriceReal   []float64 `bson:"price_real"`
	PriceDolar  []float64 `bson:"price_dolar"`
	Dynamic     []string
	Platform    []string
	Url         string
	Extra       []string
	Description string
	Rate        int
}

type CourseScore struct {
	Course Courses
	Score  int
}

type Filter struct {
	Type     string
	Course   string
	Based    string
	Dynamic  string
	Platform string
	Extra    string
	Price    string
	Length   int
}

type TypeCourses struct {
	Language []string
}

func hasStr(value string, list []string) bool {
	for i := range list {
		if list[i] == value {
			return true
		}
	}
	return false
}

func SortScore(cs []CourseScore) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(cs)-1; i++ {
			if cs[i+1].Score > cs[i].Score {
				cs[i], cs[i+1] = cs[i+1], cs[i]
				swapped = true
			}
		}
	}
}

func (f *Filter) filterCourse(data []Courses) []CourseScore {
	scored := make([]CourseScore, 0)
	for index := range data {
		score := 0
		if f.Based != "" && hasStr(f.Based, data[index].Based) {
			score++
		}
		if f.Dynamic != "" && hasStr(f.Dynamic, data[index].Dynamic) {
			score++
		}
		if f.Platform != "" && hasStr(f.Platform, data[index].Platform) {
			score++
		}
		if f.Extra != "" && hasStr(f.Extra, data[index].Extra) {
			score++
		}
		scored = append(scored, CourseScore{
			Course: data[index],
			Score:  score,
		})
	}
	SortScore(scored)
	return scored
}

func (f *Filter) limitCourse(courses []CourseScore) []Courses {
	if f.Length > len(courses) {
		f.Length = len(courses)
	}
	var result []Courses
	for i := 0; i < f.Length; i++ {
		result = append(result, courses[i].Course)
	}
	return result
}

func (f *Filter) GetCourses(path_json ...string) ([]Courses, error) {
	db, err := GetSession(path_json...)
	if err != nil {
		return nil, err
	}
	defer db.session.Close()
	var data []Courses
	err = db.session.DB(db.DBName).C(fmt.Sprintf("%s_%s", f.Type, f.Course)).Find(nil).All(&data)
	if err != nil {
		return nil, err
	}
	scoredCourses := f.filterCourse(data)
	if f.Length != 0 {
		data = f.limitCourse(scoredCourses)
	}
	return data, nil
}

func GetTypeCourses(path_json ...string) ([]TypeCourses, error) {
	db, err := GetSession(path_json...)
	if err != nil {
		return nil, err
	}
	defer db.session.Close()
	var data []TypeCourses
	err = db.session.DB(db.DBName).C("courses").Find(nil).All(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
