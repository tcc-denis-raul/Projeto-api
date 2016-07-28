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
