package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/tcc-denis-raul/Projeto-api/conf"
)

const (
	DefaultDataBaseURL  = "127.0.0.1:27017"
	DefaultDataBaseName = "paloma"
)

type DB struct {
	session *mgo.Session
	DBName  string
}

func GetSession() (DB, error) {
	host := DefaultDataBaseURL
	name := DefaultDataBaseName
	conf, err := conf.Conf()
	if err != nil {
		return DB{}, err
	}
	if conf.Database.URL != "" {
		host = conf.Database.URL
	}
	if conf.Database.Name != "" {
		name = conf.Database.Name
	}
	session, err := mgo.Dial(host)
	if err != nil {
		return DB{}, err
	}
	session.SetMode(mgo.Monotonic, true)

	return DB{session, name}, nil
}

type Courses struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string
	Based       []string
	PriceReal   []float64
	PriceDolar  []float64
	Dynamic     []string
	Platform    []string
	Url         string
	Extra       []string
	Description string
}

func GetCourses(typ, course string) ([]Courses, error) {
	db, err := GetSession()
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
