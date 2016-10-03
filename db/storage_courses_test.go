package db

import (
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (s *StorageTest) TestGetCoursesEmptyList(c *C) {
	f := Filter{
		Type:   "language",
		Course: "ingles",
	}
	data, err := f.GetCourses("data_test")
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
	f := Filter{
		Type:   "language",
		Course: "ingles",
	}
	data, err := f.GetCourses("data_test")
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
	f := Filter{
		Type:   "language",
		Course: "ingles",
	}
	data, err := f.GetCourses("data")
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "open data/paloma.json: no such file or directory")
	c.Check(len(data), Equals, 0)
}

func (s *StorageTest) TestGetCoursesWrongURLDDB(c *C) {
	f := Filter{
		Type:   "language",
		Course: "ingles",
	}
	data, err := f.GetCourses("wrong_test")
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "no reachable servers")
	c.Check(len(data), Equals, 0)
}

func (s *StorageTest) TestGetTypeCoursesEmptyList(c *C) {
	data, err := GetTypeCourses("data_test")
	c.Check(err, IsNil)
	c.Check(len(data), Equals, 0)
}

func (s *StorageTest) TestGetTypeCoursesLanguage(c *C) {
	typeCourse := TypeCourses{
		Language: []string{"ingles"},
	}
	err := s.session.DB(s.dbName).C("courses").Insert(&typeCourse)
	c.Check(err, IsNil)
	defer s.session.DB(s.dbName).C("courses").Remove(nil)
	data, err := GetTypeCourses("data_test")
	c.Check(err, IsNil)
	c.Check(len(data), Equals, 1)
	c.Check(data[0].Language, DeepEquals, []string{"ingles"})
}

func (s *StorageTest) TestFilterCourse(c *C) {
	courses := []Courses{
		{
			Name:        "name2",
			Based:       []string{"base3", "based4"},
			PriceReal:   []float64{2.0, 3.0},
			PriceDolar:  []float64{4.0, 5.0},
			Dynamic:     []string{"dyn 3", "dyn 4"},
			Platform:    []string{"desktop", "android"},
			Url:         "url_course",
			Extra:       []string{"ext 3", "ext 4"},
			Description: "descr",
		},
		{
			Name:        "name",
			Based:       []string{"base1", "based2"},
			PriceReal:   []float64{2.0, 3.0},
			PriceDolar:  []float64{4.0, 5.0},
			Dynamic:     []string{"dyn 1", "dyn 2"},
			Platform:    []string{"desktop", "android", "dif"},
			Url:         "url_course",
			Extra:       []string{"ext 1", "ext 2"},
			Description: "descr",
		},
	}
	f := Filter{
		Type:     "language",
		Course:   "ingles",
		Based:    "base1",
		Dynamic:  "dyn 1",
		Platform: "desktop",
		Extra:    "ext 1",
	}
	data := f.filterCourse(courses)
	c.Check(len(data), Equals, 2)
	c.Check(data[0].Course.Name, Equals, courses[1].Name)
	c.Check(data[0].Score, Equals, 4)
	c.Check(data[1].Course.Name, Equals, courses[0].Name)
	c.Check(data[1].Score, Equals, 1)
}

func (s *StorageTest) TestLimitCourse(c *C) {
	courses := []CourseScore{
		CourseScore{
			Course: Courses{
				Name:        "name",
				Based:       []string{"base1", "based2"},
				PriceReal:   []float64{2.0, 3.0},
				PriceDolar:  []float64{4.0, 5.0},
				Dynamic:     []string{"dyn 1", "dyn 2"},
				Platform:    []string{"desktop", "android", "dif"},
				Url:         "url_course",
				Extra:       []string{"ext 1", "ext 2"},
				Description: "descr",
			},
			Score: 10,
		},
		CourseScore{
			Course: Courses{
				Name:        "name2",
				Based:       []string{"base3", "based4"},
				PriceReal:   []float64{2.0, 3.0},
				PriceDolar:  []float64{4.0, 5.0},
				Dynamic:     []string{"dyn 3", "dyn 4"},
				Platform:    []string{"desktop", "android"},
				Url:         "url_course",
				Extra:       []string{"ext 3", "ext 4"},
				Description: "descr",
			},
			Score: 20,
		},
	}
	f := Filter{
		Type:   "language",
		Course: "ingles",
		Length: 1,
	}
	data := f.limitCourse(courses)
	c.Check(len(data), Equals, 1)
	c.Check(data[0].Name, Equals, courses[0].Course.Name)
}

func (s *StorageTest) TestSort(c *C) {
	courses := []CourseScore{
		CourseScore{
			Course: Courses{
				Name:        "name",
				Based:       []string{"base1", "based2"},
				PriceReal:   []float64{2.0, 3.0},
				PriceDolar:  []float64{4.0, 5.0},
				Dynamic:     []string{"dyn 1", "dyn 2"},
				Platform:    []string{"desktop", "android", "dif"},
				Url:         "url_course",
				Extra:       []string{"ext 1", "ext 2"},
				Description: "descr",
			},
			Score: 1,
		},
		CourseScore{
			Course: Courses{
				Name:        "name2",
				Based:       []string{"base3", "based4"},
				PriceReal:   []float64{2.0, 3.0},
				PriceDolar:  []float64{4.0, 5.0},
				Dynamic:     []string{"dyn 3", "dyn 4"},
				Platform:    []string{"desktop", "android"},
				Url:         "url_course",
				Extra:       []string{"ext 3", "ext 4"},
				Description: "descr",
			},
			Score: 55,
		},
		CourseScore{
			Course: Courses{
				Name:        "name3",
				Based:       []string{"base3", "based4"},
				PriceReal:   []float64{2.0, 3.0},
				PriceDolar:  []float64{4.0, 5.0},
				Dynamic:     []string{"dyn 3", "dyn 4"},
				Platform:    []string{"desktop", "android"},
				Url:         "url_course",
				Extra:       []string{"ext 3", "ext 4"},
				Description: "descr",
			},
			Score: 10,
		},
	}
	SortScore(courses)
	c.Check(len(courses), Equals, 3)
	c.Check(courses[0].Course.Name, Equals, "name2")
	c.Check(courses[0].Score, Equals, 55)
	c.Check(courses[1].Course.Name, Equals, "name3")
	c.Check(courses[1].Score, Equals, 10)
	c.Check(courses[2].Course.Name, Equals, "name")
	c.Check(courses[2].Score, Equals, 1)

}

func (s *StorageTest) TestIndicateCourse(c *C) {
	indication := IndicateCourse{
		Type:   "language",
		Course: "ingles",
		Name:   "indication",
		Url:    "url",
	}
	err := indication.IndicateCourse("data_test")
	defer s.session.DB(s.dbName).C("indicate_courses").Remove(nil)
	c.Check(err, IsNil)
	var result []IndicateCourse
	err = s.session.DB(s.dbName).C("indicate_courses").Find(bson.M{"url": indication.Url}).All(&result)
	c.Check(err, IsNil)
	c.Check(len(result), Equals, 1)
	c.Check(result[0].Name, Equals, indication.Name)
	c.Check(result[0].Type, Equals, indication.Type)
	c.Check(result[0].Course, Equals, indication.Course)
	c.Check(result[0].Url, Equals, indication.Url)
}

func (s *StorageTest) TestIndicateCourseAlreadyIndicate(c *C) {
	indication := IndicateCourse{
		Type:   "language",
		Course: "ingles",
		Name:   "indication",
		Url:    "url",
	}
	err := indication.IndicateCourse("data_test")
	defer s.session.DB(s.dbName).C("indicate_courses").Remove(nil)
	c.Check(err, IsNil)
	var result []IndicateCourse
	err = s.session.DB(s.dbName).C("indicate_courses").Find(bson.M{"url": indication.Url}).All(&result)
	c.Check(err, IsNil)
	c.Check(len(result), Equals, 1)
	c.Check(result[0].Name, Equals, indication.Name)
	c.Check(result[0].Type, Equals, indication.Type)
	c.Check(result[0].Course, Equals, indication.Course)
	c.Check(result[0].Url, Equals, indication.Url)
	err = indication.IndicateCourse("data_test")
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "Course already indicate")
}

func (s *StorageTest) TestIndicateCourseAlreadyInDataBase(c *C) {
	Course := Courses{
		Name:        "name3",
		Based:       []string{"base3", "based4"},
		PriceReal:   []float64{2.0, 3.0},
		PriceDolar:  []float64{4.0, 5.0},
		Dynamic:     []string{"dyn 3", "dyn 4"},
		Platform:    []string{"desktop", "android"},
		Url:         "url_course",
		Extra:       []string{"ext 3", "ext 4"},
		Description: "descr",
	}
	err := s.session.DB(s.dbName).C("language_ingles").Insert(&Course)
	c.Check(err, IsNil)
	defer s.session.DB(s.dbName).C("language_ingles").Remove(bson.M{"name": Course.Name})
	f := Filter{
		Type:   "language",
		Course: "ingles",
	}
	data, err := f.GetCourses("data_test")
	c.Check(err, IsNil)
	c.Check(len(data), Equals, 1)
	c.Check(data[0].Name, Equals, Course.Name)
	indication := IndicateCourse{
		Type:   "language",
		Course: "ingles",
		Name:   "name3",
		Url:    "url_course",
	}
	err = indication.IndicateCourse("data_test")
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "Course already exists on list of courses")
}

func (s *StorageTest) TestCourseDetail(c *C) {
	Course := Courses{
		Name:        "name3",
		Based:       []string{"base3", "based4"},
		PriceReal:   []float64{2.0, 3.0},
		PriceDolar:  []float64{4.0, 5.0},
		Dynamic:     []string{"dyn 3", "dyn 4"},
		Platform:    []string{"desktop", "android"},
		Url:         "url_course",
		Extra:       []string{"ext 3", "ext 4"},
		Description: "descr",
	}
	err := s.session.DB(s.dbName).C("language_ingles").Insert(&Course)
	c.Check(err, IsNil)
	defer s.session.DB(s.dbName).C("language_ingles").Remove(bson.M{"name": Course.Name})
	f := CourseDetail{
		Type:   "language",
		Course: "ingles",
		Name:   "name3",
	}
	data, err := f.GetDetailCourse("data_test")
	c.Check(err, IsNil)
	c.Check(data.Name, DeepEquals, Course.Name)
	c.Check(data.Url, DeepEquals, Course.Url)
	c.Check(data.Platform, DeepEquals, Course.Platform)
	c.Check(data.Based, DeepEquals, Course.Based)
	c.Check(data.PriceReal, DeepEquals, Course.PriceReal)
	c.Check(data.PriceDolar, DeepEquals, Course.PriceDolar)
	c.Check(data.Dynamic, DeepEquals, Course.Dynamic)
	c.Check(data.Extra, DeepEquals, Course.Extra)
}

func (s *StorageTest) TestCourseDetailNotExistCourse(c *C) {
	f := CourseDetail{
		Type:   "language",
		Course: "ingles",
		Name:   "aba",
	}
	data, err := f.GetDetailCourse("data_test")
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "not found")
	c.Check(data, DeepEquals, Courses{})
}
