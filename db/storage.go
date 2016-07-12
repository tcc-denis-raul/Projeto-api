package main

import (
	// "encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	// "io/ioutil"
	// "os"
	// "time"

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

func main() {
	a, b := GetSession()
	fmt.Println(a, b)

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
