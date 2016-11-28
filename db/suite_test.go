package db

import (
	"os"
	"testing"

	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type StorageTest struct {
	session *mgo.Session
	dbName  string
}

var _ = Suite(&StorageTest{})

func (s *StorageTest) TearDownSuite(c *C) {
	defer s.session.Close()
}

func (s *StorageTest) SetUpTest(c *C) {
	os.Setenv("MONGO_DBNAME", "paloma_test")
	session, err := mgo.Dial("localhost:27017")
	c.Check(err, IsNil)
	s.dbName = "paloma_test"
	s.session = session
	err = s.session.DB(s.dbName).C("users").EnsureIndex(mgo.Index{Key: []string{"username"}, Unique: true})
	c.Check(err, IsNil)
	err = s.session.DB(s.dbName).C("language_ingles").EnsureIndex(mgo.Index{Key: []string{"name"}, Unique: true})
	c.Check(err, IsNil)
	err = s.session.DB(s.dbName).C("user_profile_courses").EnsureIndex(mgo.Index{Key: []string{"username"}, Unique: true})
	c.Check(err, IsNil)
	err = s.session.DB(s.dbName).C("indicate_courses").EnsureIndex(mgo.Index{Key: []string{"url"}, Unique: true})
	c.Check(err, IsNil)

}
