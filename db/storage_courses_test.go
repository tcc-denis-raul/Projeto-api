package db

import (
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (s *StorageTest) TestGetCoursesEmptyList(c *C) {
	data, err := GetCourses("language", "ingles", "data_test")
	c.Check(err, IsNil)
	c.Check(len(data), Equals, 0)
}

func (s *StorageTest) TestGetCoursesReturnList(c *C) {
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
		},
	}
	for _, course := range courses {
		err := s.session.DB(s.dbName).C("language_ingles").Insert(&course)
		c.Check(err, IsNil)
		defer s.session.DB(s.dbName).C("language_ingles").Remove(bson.M{"name": course.Name})
	}
	data, err := GetCourses("language", "ingles", "data_test")
	c.Check(err, IsNil)
	c.Check(len(data), Equals, 2)
	c.Check(data[0].Name, Equals, courses[0].Name)
	c.Check(data[0].Based, DeepEquals, courses[0].Based)
	c.Check(data[0].PriceDolar, DeepEquals, courses[0].PriceDolar)
	c.Check(data[0].PriceReal, DeepEquals, courses[0].PriceReal)
	c.Check(data[0].Dynamic, DeepEquals, courses[0].Dynamic)
	c.Check(data[0].Platform, DeepEquals, courses[0].Platform)
	c.Check(data[0].Url, DeepEquals, courses[0].Url)
	c.Check(data[0].Extra, DeepEquals, courses[0].Extra)
	c.Check(data[0].Description, DeepEquals, courses[0].Description)
	c.Check(data[1].Name, Equals, courses[1].Name)
	c.Check(data[1].Based, DeepEquals, courses[1].Based)
	c.Check(data[1].PriceDolar, DeepEquals, courses[1].PriceDolar)
	c.Check(data[1].PriceReal, DeepEquals, courses[1].PriceReal)
	c.Check(data[1].Dynamic, DeepEquals, courses[1].Dynamic)
	c.Check(data[1].Platform, DeepEquals, courses[1].Platform)
	c.Check(data[1].Url, DeepEquals, courses[1].Url)
	c.Check(data[1].Extra, DeepEquals, courses[1].Extra)
	c.Check(data[1].Description, DeepEquals, courses[1].Description)
}

func (s *StorageTest) TestGetCoursesWrongPath(c *C) {
	data, err := GetCourses("language", "ingles", "data")
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "open data/paloma.json: no such file or directory")
	c.Check(len(data), Equals, 0)
}

func (s *StorageTest) TestGetCoursesWrongURLDDB(c *C) {
	data, err := GetCourses("language", "ingles", "wrong_test")
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "no reachable servers")
	c.Check(len(data), Equals, 0)
}
