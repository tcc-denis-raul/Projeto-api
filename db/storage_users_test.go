package db

import (
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (s *StorageTest) TestCreateUser(c *C) {
	user := User{
		Name:     "name",
		Email:    "email",
		Password: "pass",
	}
	err := user.CreateUser("data_test")
	c.Check(err, IsNil)
	defer s.session.DB(s.dbName).C("users").RemoveAll(bson.M{"email": "email"})
	var u []User
	err = s.session.DB(s.dbName).C("users").Find(bson.M{"email": "email"}).All(&u)
	c.Check(err, IsNil)
	c.Check(len(u), Equals, 1)
	c.Check(u[0].Name, Equals, user.Name)
	c.Check(u[0].Email, Equals, user.Email)
	c.Check(u[0].Password, Equals, user.Password)
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
		Name:     "name",
		Email:    "email",
		Password: "pass",
	}
	err := user.CreateUser("data_test")
	c.Check(err, IsNil)
	defer s.session.DB(s.dbName).C("users").RemoveAll(bson.M{"email": "email"})
	var u []User
	err = s.session.DB(s.dbName).C("users").Find(bson.M{"email": "email"}).All(&u)
	c.Check(err, IsNil)
	c.Check(len(u), Equals, 1)
	c.Check(u[0].Name, Equals, user.Name)
	c.Check(u[0].Email, Equals, user.Email)
	c.Check(u[0].Password, Equals, user.Password)
	err = user.CreateUser("data_test")
	c.Check(err, NotNil)
}

func (s *StorageTest) TestUpdateUser(c *C) {
	user := User{
		Name:     "name",
		Email:    "email",
		Password: "pass",
	}
	err := user.CreateUser("data_test")
	c.Check(err, IsNil)
	defer s.session.DB(s.dbName).C("users").RemoveAll(bson.M{"email": "email"})
	var u []User
	err = s.session.DB(s.dbName).C("users").Find(bson.M{"email": "email"}).All(&u)
	c.Check(err, IsNil)
	c.Check(len(u), Equals, 1)
	c.Check(u[0].Name, Equals, user.Name)
	c.Check(u[0].Email, Equals, user.Email)
	c.Check(u[0].Password, Equals, user.Password)
	updateData := User{
		Name:     "name_mod",
		Email:    "email",
		Password: "pass_mod",
	}
	err = updateData.UpdateUser("data_test")
	c.Check(err, IsNil)
	err = s.session.DB(s.dbName).C("users").Find(bson.M{"email": "email"}).All(&u)
	c.Check(err, IsNil)
	c.Check(len(u), Equals, 1)
	c.Check(u[0].Name, Equals, updateData.Name)
	c.Check(u[0].Email, Equals, updateData.Email)
	c.Check(u[0].Password, Equals, updateData.Password)
}

func (s *StorageTest) TestUpdateUserNotExists(c *C) {
	user := User{
		Name:     "name",
		Email:    "email",
		Password: "pass",
	}
	err := user.UpdateUser("data_test")
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "not found")
}
