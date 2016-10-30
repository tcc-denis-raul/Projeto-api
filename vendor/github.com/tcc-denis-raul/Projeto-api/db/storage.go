package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/tcc-denis-raul/Projeto-api/conf"
)

const (
	DefaultDataBaseURL  = "localhost:27017"
	DefaultDataBaseName = "paloma"
)

type DB struct {
	session *mgo.Session
	DBName  string
}

func GetSession(path_json ...string) (DB, error) {
	host := DefaultDataBaseURL
	name := DefaultDataBaseName
	conf, err := conf.Conf(path_json...)
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

func (c *Courses) Feedback(typ, course string, path_json ...string) error {
	db, err := GetSession(path_json...)
	if err != nil {
		return err
	}
	defer db.session.Close()
	return db.session.DB(db.DBName).C(fmt.Sprintf("%s_%s", typ, course)).Update(bson.M{"name": c.Name}, bson.M{"$inc": bson.M{"rate": c.Rate}})
}
