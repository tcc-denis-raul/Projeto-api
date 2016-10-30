package db

import (
	"os"

	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (s *StorageTest) TestGetSession(c *C) {
	os.Setenv("MONGOLAB_URL", "localhost:27017")
	os.Setenv("MONGO_DBNAME", "paloma_test")
	db, err := GetSession()
	c.Check(err, IsNil)
	c.Check(db.DBName, Equals, "paloma_test")
	c.Check(db.session, NotNil)
}

func (s *StorageTest) TestGetSessionWrongURL(c *C) {
	os.Setenv("MONGOLAB_URL", "wrong")
	db, err := GetSession()
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "no reachable servers")
	c.Check(db.DBName, Equals, "")
	c.Check(db.session, IsNil)
}

func (s *StorageTest) TestFeedBackOk(c *C) {
	os.Setenv("MONGO_DBNAME", "paloma_test")
	courses := []Courses{
		{
			Name:        "name",
			Based:       []string{"base1", "based2"},
			PriceReal:   []float64{2.0, 3.0},
			PriceDolar:  []float64{4.0, 5.0},
			Dynamic:     []string{"dyn 1", "dyn 2"},
			Platform:    []string{"desktop", "android"},
			Url:         "url_course",
			Extra:       []string{"ext 1", "ext 2"},
			Description: "descr",
			Rate:        1,
		},
		{
			Name:        "name2",
			Based:       []string{"base1", "based2"},
			PriceReal:   []float64{2.0, 3.0},
			PriceDolar:  []float64{4.0, 5.0},
			Dynamic:     []string{"dyn 1", "dyn 2"},
			Platform:    []string{"desktop", "android"},
			Url:         "url_course",
			Extra:       []string{"ext 1", "ext 2"},
			Description: "descr",
			Rate:        2,
		},
	}
	for _, course := range courses {
		err := s.session.DB(s.dbName).C("language_ingles").Insert(&course)
		c.Check(err, IsNil)
		defer s.session.DB(s.dbName).C("language_ingles").Remove(bson.M{"name": course.Name})
	}
	course := Courses{
		Name: "name",
		Rate: 1,
	}
	err := course.Feedback("language", "ingles")
	c.Check(err, IsNil)
	var result []Courses
	err = s.session.DB(s.dbName).C("language_ingles").Find(bson.M{"name": "name"}).All(&result)
	c.Check(err, IsNil)
	c.Check(len(result), Equals, 1)
	c.Check(result[0].Name, Equals, courses[0].Name)
	c.Check(result[0].Rate, Equals, courses[0].Rate+1)
}

func (s *StorageTest) TestFeedBackCourseNotFound(c *C) {
	os.Setenv("MONGO_DBNAME", "paloma_test")
	course := Courses{
		Name: "name",
		Rate: 1,
	}
	err := course.Feedback("language", "ingles")
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "not found")
}
