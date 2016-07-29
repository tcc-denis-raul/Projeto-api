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

type TypeCourses struct {
	Language []string
}

func GetCourses(typ, course string, path_json ...string) ([]Courses, error) {
	db, err := GetSession(path_json...)
	if err != nil {
		return nil, err
	}
	defer db.session.Close()
	var data []Courses
	err = db.session.DB(db.DBName).C(fmt.Sprintf("%s_%s", typ, course)).Find(nil).All(&data)
	if err != nil {
		return nil, err
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
