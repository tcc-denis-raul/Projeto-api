package db

import (
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (s *StorageTest) TestCreateUser(c *C) {
	user := User{
		FirstName: "first",
		LastName:  "last",
		Email:     "email",
		UserName:  "username",
	}
	err := user.CreateUser("data_test")
	c.Check(err, IsNil)
	defer s.session.DB(s.dbName).C("users").RemoveAll(bson.M{"username": "username"})
	var u []User
	err = s.session.DB(s.dbName).C("users").Find(bson.M{"username": "username"}).All(&u)
	c.Check(err, IsNil)
	c.Check(len(u), Equals, 1)
	c.Check(u[0].FirstName, Equals, user.FirstName)
	c.Check(u[0].LastName, Equals, user.LastName)
	c.Check(u[0].Email, Equals, user.Email)
	c.Check(u[0].UserName, Equals, user.UserName)
}

func (s *StorageTest) TestCreateUserWrongPath(c *C) {
	var user User
	err := user.CreateUser("data")
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "open data/paloma.json: no such file or directory")
}

func (s *StorageTest) TestCreateUserWrongURLDDB(c *C) {
	var user User
	err := user.CreateUser("wrong_test")
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "no reachable servers")
}

func (s *StorageTest) TestCreateUserAlreadyExists(c *C) {
	user := User{
		FirstName: "first",
		LastName:  "last",
		Email:     "email",
		UserName:  "username",
	}
	err := user.CreateUser("data_test")
	c.Check(err, IsNil)
	defer s.session.DB(s.dbName).C("users").RemoveAll(bson.M{"username": "username"})
	var u []User
	err = s.session.DB(s.dbName).C("users").Find(bson.M{"username": "username"}).All(&u)
	c.Check(err, IsNil)
	c.Check(len(u), Equals, 1)
	c.Check(u[0].FirstName, Equals, user.FirstName)
	c.Check(u[0].LastName, Equals, user.LastName)
	c.Check(u[0].Email, Equals, user.Email)
	c.Check(u[0].UserName, Equals, user.UserName)
	err = user.CreateUser("data_test")
	c.Check(err, NotNil)
}

func (s *StorageTest) TestUpdateUser(c *C) {
	user := User{
		FirstName: "first",
		LastName:  "last",
		Email:     "email",
		UserName:  "username",
	}
	err := user.CreateUser("data_test")
	c.Check(err, IsNil)
	defer s.session.DB(s.dbName).C("users").RemoveAll(bson.M{"username": "username"})
	var u []User
	err = s.session.DB(s.dbName).C("users").Find(bson.M{"username": "username"}).All(&u)
	c.Check(err, IsNil)
	c.Check(len(u), Equals, 1)
	c.Check(u[0].FirstName, Equals, user.FirstName)
	c.Check(u[0].LastName, Equals, user.LastName)
	c.Check(u[0].Email, Equals, user.Email)
	c.Check(u[0].UserName, Equals, user.UserName)
	updateData := User{
		FirstName: "name_mod",
		UserName:  "username",
	}
	err = updateData.UpdateUser("data_test")
	c.Check(err, IsNil)
	err = s.session.DB(s.dbName).C("users").Find(bson.M{"username": "username"}).All(&u)
	c.Check(err, IsNil)
	c.Check(len(u), Equals, 1)
	c.Check(u[0].FirstName, Equals, updateData.FirstName)
	c.Check(u[0].LastName, Equals, user.LastName)
	c.Check(u[0].Email, Equals, user.Email)
	c.Check(u[0].UserName, Equals, user.UserName)
}

func (s *StorageTest) TestUpdateUserNotExists(c *C) {
	user := User{
		FirstName: "first",
		LastName:  "last",
		Email:     "email",
		UserName:  "username",
	}
	err := user.UpdateUser("data_test")
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "not found")
}

func (s *StorageTest) TestSaveUserPreferencesUserDoesNotExists(c *C) {
	preferences := UserPreferences{
		UserName: "username",
		Based:    "based",
		Dynamic:  "dynamic",
		Platform: "platform",
		Extra:    "extra",
		Price:    "price",
	}
	err := preferences.SaveUserPreferences("data_test")
	c.Check(err, IsNil)
	var prefers []UserPreferences
	err = s.session.DB(s.dbName).C("user_profile_courses").Find(bson.M{"username": "username"}).All(&prefers)
	defer s.session.DB(s.dbName).C("user_profile_courses").RemoveAll(bson.M{"username": "username"})
	c.Check(err, IsNil)
	c.Check(len(prefers), Equals, 1)
	c.Check(prefers[0].UserName, Equals, preferences.UserName)
	c.Check(prefers[0].Based, Equals, preferences.Based)
	c.Check(prefers[0].Dynamic, Equals, preferences.Dynamic)
	c.Check(prefers[0].Platform, Equals, preferences.Platform)
	c.Check(prefers[0].Extra, Equals, preferences.Extra)
	c.Check(prefers[0].Price, Equals, preferences.Price)
}

func (s *StorageTest) TestSaveUserPreferencesUserExists(c *C) {
	preferences := UserPreferences{
		UserName: "username",
		Based:    "based",
		Dynamic:  "dynamic",
		Platform: "platform",
		Extra:    "extra",
		Price:    "price",
	}
	err := preferences.SaveUserPreferences("data_test")
	c.Check(err, IsNil)
	var prefers []UserPreferences
	err = s.session.DB(s.dbName).C("user_profile_courses").Find(bson.M{"username": "username"}).All(&prefers)
	defer s.session.DB(s.dbName).C("user_profile_courses").RemoveAll(bson.M{"username": "username"})
	c.Check(err, IsNil)
	c.Check(len(prefers), Equals, 1)
	c.Check(prefers[0].UserName, Equals, preferences.UserName)
	c.Check(prefers[0].Based, Equals, preferences.Based)
	c.Check(prefers[0].Dynamic, Equals, preferences.Dynamic)
	c.Check(prefers[0].Platform, Equals, preferences.Platform)
	c.Check(prefers[0].Extra, Equals, preferences.Extra)
	c.Check(prefers[0].Price, Equals, preferences.Price)
	preferences2 := UserPreferences{
		UserName: "username",
		Based:    "based2",
		Dynamic:  "dynamic",
		Platform: "platform",
		Extra:    "extra2",
		Price:    "price",
	}
	err = preferences2.SaveUserPreferences("data_test")
	c.Check(err, IsNil)
	err = s.session.DB(s.dbName).C("user_profile_courses").Find(bson.M{"username": "username"}).All(&prefers)
	c.Check(err, IsNil)
	c.Check(len(prefers), Equals, 1)
	c.Check(prefers[0].UserName, Equals, preferences2.UserName)
	c.Check(prefers[0].Based, Equals, preferences2.Based)
	c.Check(prefers[0].Dynamic, Equals, preferences2.Dynamic)
	c.Check(prefers[0].Platform, Equals, preferences2.Platform)
	c.Check(prefers[0].Extra, Equals, preferences2.Extra)
	c.Check(prefers[0].Price, Equals, preferences2.Price)
}
