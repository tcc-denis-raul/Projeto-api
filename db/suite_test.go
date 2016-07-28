package db

import (
	"testing"

	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2"

	"Projeto-api/conf"
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
	conf, err := conf.Conf("data_test")
	c.Check(err, IsNil)
	c.Check(conf.URL, Equals, "localhost")
	c.Check(conf.Name, Equals, "paloma_test")
	s.session, err = mgo.Dial(conf.URL)
	c.Check(err, IsNil)
	s.dbName = conf.Name
	err = s.session.DB(s.dbName).C("users").EnsureIndex(mgo.Index{Key: []string{"$text:email"}, Unique: true})
	c.Check(err, IsNil)
	err = s.session.DB(s.dbName).C("language_ingles").EnsureIndex(mgo.Index{Key: []string{"$text:name"}, Unique: true})
	c.Check(err, IsNil)
}
