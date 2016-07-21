package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"

	"Projeto-api/conf"
)

const (
	DefaultDataBaseURL  = "127.0.0.1:27017"
	DefaultDataBaseName = "paloma"
)

type DB struct {
	session *mgo.Session
	DBName  string
}

type User struct {
	Name      string
	Email     string
	Password  string
	Created   time.Time
	LastAcess time.Time
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
	Rate        int
}

type Questions struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Based    []map[string]string
	Price    []map[string]string
	Dynamic  []map[string]string
	Platform []map[string]string
	Extra    []map[string]string
}

func GetSession(path_json string) (DB, error) {
	host := DefaultDataBaseURL
	name := DefaultDataBaseName
	conf, err := conf.Conf(path_json)
	if err != nil {
		return DB{}, err
	}
	if conf.URL != "" {
		host = conf.URL
	}
	if conf.Name != "" {
		name = conf.Name
	}
	session, err := mgo.Dial(host)
	if err != nil {
		return DB{}, err
	}
	session.SetMode(mgo.Monotonic, true)

	return DB{session, name}, nil
}

func GetCourses(typ, course, path_json string) ([]Courses, error) {
	db, err := GetSession(path_json)
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

func GetQuestions(typ, path_json string) ([]Questions, error) {
	db, err := GetSession(path_json)
	if err != nil {
		return nil, err
	}
	defer db.session.Close()
	var data []Questions
	err = db.session.DB(db.DBName).C(fmt.Sprintf("questions_%s", typ)).Find(nil).All(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *User) CreateUser(path_json string) error {
	db, err := GetSession(path_json)
	if err != nil {
		return err
	}
	defer db.session.Close()
	return db.session.DB(db.DBName).C("users").Insert(u)
}

func (u *User) UpdateUser(path_json string) error {
	db, err := GetSession(path_json)
	if err != nil {
		return err
	}
	defer db.session.Close()
	var updateData User
	if u.Name != "" {
		updateData.Name = u.Name
	}
	if u.Email != "" {
		updateData.Email = u.Email
	}
	if u.Password != "" {
		updateData.Password = u.Password
	}
	if !u.LastAcess.IsZero() {
		updateData.LastAcess = u.LastAcess
	}
	return db.session.DB(db.DBName).C("users").Update(bson.M{"email": u.Email}, bson.M{"$set": updateData})
}

func (c *Courses) Feedback(typ, course, path_json string) error {
	db, err := GetSession(path_json)
	if err != nil {
		return err
	}
	defer db.session.Close()
	return db.session.DB(db.DBName).C(fmt.Sprintf("%s_%s", typ, course)).Update(bson.M{"name": c.Name}, bson.M{"$inc": bson.M{"rate": c.Rate}})
}
