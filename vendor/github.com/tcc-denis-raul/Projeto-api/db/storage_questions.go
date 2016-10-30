package db

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Questions struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Based    []map[string]string
	Price    []map[string]string
	Dynamic  []map[string]string
	Platform []map[string]string
	Extra    []map[string]string
}

func GetQuestions(typ string, path_json ...string) ([]Questions, error) {
	db, err := GetSession()
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
