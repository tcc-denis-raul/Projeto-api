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

type UserPreferences struct {
	UserName string
	Based    string
	Dynamic  string
	Platform string
	Extra    string
	Price    string
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
	updateData, err := u.GetUser(path_json...)
	if err != nil {
		return err
	}
	if u.FirstName != "" {
		updateData.FirstName = u.FirstName
	}
	if u.LastName != "" {
		updateData.LastName = u.LastName
	}
	if u.Email != "" {
		updateData.Email = u.Email
	}
	if !u.LastAcess.IsZero() {
		updateData.LastAcess = u.LastAcess
	}
	return db.session.DB(db.DBName).C("users").Update(bson.M{"username": u.UserName}, bson.M{"$set": updateData})
}

func (up *UserPreferences) SaveUserPreferences(path_json ...string) error {
	db, err := GetSession(path_json...)
	if err != nil {
		return err
	}
	defer db.session.Close()
	_, err = db.session.DB(db.DBName).C("user_profile_courses").Upsert(bson.M{"username": up.UserName}, bson.M{"$set": up})
	return err
}
