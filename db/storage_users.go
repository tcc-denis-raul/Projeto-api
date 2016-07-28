package db

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Name      string
	Email     string
	Password  string
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
	err = db.session.DB(db.DBName).C("users").Find(bson.M{"email": u.Email}).One(&user)
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
