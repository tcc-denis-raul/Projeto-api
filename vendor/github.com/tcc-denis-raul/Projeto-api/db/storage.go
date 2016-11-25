package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

const (
	DefaultDataBaseURL  = "localhost:27017"
	DefaultDataBaseName = "paloma"
)

type DB struct {
	session *mgo.Session
	DBName  string
}

func GetSession() (DB, error) {
	host := os.Getenv("MONGOLAB_URL")
	name := os.Getenv("MONGO_DBNAME")
	if host == "" {
		host = DefaultDataBaseURL
	}
	if name == "" {
		name = DefaultDataBaseName
	}
	session, err := mgo.Dial(host)
	if err != nil {
		return DB{}, err
	}
	session.SetMode(mgo.Monotonic, true)

	return DB{session, name}, nil
}

func (c *Courses) Feedback(typ, course string) error {
	db, err := GetSession()
	if err != nil {
		return err
	}
	defer db.session.Close()
	var cs Courses
	err = db.session.DB(db.DBName).C(fmt.Sprintf("%s_%s", typ, course)).Find(bson.M{"name": c.Name}).One(&cs)
	if err != nil {
		return err
	}
	count := cs.Count + 1
	rate := (c.Rate + cs.Rate) / float64(count)
	return db.session.DB(db.DBName).C(fmt.Sprintf("%s_%s", typ, course)).Update(bson.M{"name": c.Name}, bson.M{"$set": bson.M{"rate": rate, "count": count}})
}
