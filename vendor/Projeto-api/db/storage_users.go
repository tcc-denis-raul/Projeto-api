package db

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	UserName  string
	Email     string
	Created   time.Time
	LastAcess time.Time
}

func (u *User) GetUser(path_json ...string) (User, error) {
	db, err := GetSession(path_json...)
	if err != nil {
		return User{}, err
	}
	defer db.session.Close()
	var user User
	err = db.session.DB(db.DBName).C("users").Find(bson.M{"username": u.UserName}).One(&user)
	return user, err
}

func (u *User) CreateUser(path_json ...string) error {
	db, err := GetSession(path_json...)
	if err != nil {
		return err
	}
	defer db.session.Close()
	return db.session.DB(db.DBName).C("users").Insert(u)
}

func (u *User) UpdateUser(path_json ...string) error {
	db, err := GetSession(path_json...)
	if err != nil {
		return err
	}
	defer db.session.Close()
	var updateData User
	if u.FirstName != "" {
		updateData.FirstName = u.FirstName
	}
	if u.LastName != "" {
		updateData.LastName = u.LastName
	}
	if u.Email != "" {
		updateData.Email = u.Email
	}
	if u.UserName != "" {
		updateData.UserName = u.UserName
	}
	if !u.LastAcess.IsZero() {
		updateData.LastAcess = u.LastAcess
	}
	return db.session.DB(db.DBName).C("users").Update(bson.M{"email": u.Email}, bson.M{"$set": updateData})
}
