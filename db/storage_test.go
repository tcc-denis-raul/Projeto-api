package db

import (
	"testing"

	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/tcc-denis-raul/Projeto-api/conf"
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
	c.Check(conf.URL, Equals, "192.168.99.100")
	c.Check(conf.Name, Equals, "paloma_test")
	s.session, err = mgo.Dial(conf.URL)
	c.Check(err, IsNil)
	s.dbName = conf.Name
}

func (s *StorageTest) TestGetSession(c *C) {
	db, err := GetSession("data_test")
	c.Check(err, IsNil)
	c.Check(db.DBName, Equals, "paloma_test")
	c.Check(db.session, NotNil)
}

func (s *StorageTest) TestGetSessionWrongPath(c *C) {
	db, err := GetSession("data")
	c.Check(err, NotNil)
	c.Check(db.DBName, Equals, "")
	c.Check(db.session, IsNil)
}

func (s *StorageTest) TestGetSessionWrongURL(c *C) {
	db, err := GetSession("wrong_test")
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "no reachable servers")
	c.Check(db.DBName, Equals, "")
	c.Check(db.session, IsNil)
}

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

func (s *StorageTest) TestGetQuestionsEmptyList(c *C) {
	data, err := GetQuestions("language", "ingles", "data_test")
	c.Check(err, IsNil)
	c.Check(len(data), Equals, 0)
}

func (s *StorageTest) TestGetQuestionsReturnList(c *C) {
	questions := []Questions{
		{
			Based: []map[string]string{
				{"texto": "Textos"},
			},
			Price: []map[string]string{
				{"gratis": "Grátis"},
			},
			Dynamic: []map[string]string{
				{"curso_livre": "Curso Livre"},
			},
			Platform: []map[string]string{
				{"android_offline": "Android - Offline"},
			},
			Extra: []map[string]string{
				{"selecao_nivel": "Seleção de Nível de conhecimento"},
			},
		},
	}
	for _, question := range questions {
		err := s.session.DB(s.dbName).C("questions_language").Insert(&question)
		c.Check(err, IsNil)
		defer s.session.DB(s.dbName).C("questions_language").Remove(nil)
	}
	data, err := GetQuestions("language", "ingles", "data_test")
	c.Check(err, IsNil)
	c.Check(len(data), Equals, 1)
	c.Check(data[0].Based, DeepEquals, questions[0].Based)
	c.Check(data[0].Price, DeepEquals, questions[0].Price)
	c.Check(data[0].Dynamic, DeepEquals, questions[0].Dynamic)
	c.Check(data[0].Platform, DeepEquals, questions[0].Platform)
	c.Check(data[0].Extra, DeepEquals, questions[0].Extra)
}

func (s *StorageTest) TestGetQuestionsWrongPath(c *C) {
	data, err := GetQuestions("language", "ingles", "data")
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "open data/paloma.json: no such file or directory")
	c.Check(len(data), Equals, 0)
}

func (s *StorageTest) TestGetQuestionsWrongURLDDB(c *C) {
	data, err := GetQuestions("language", "ingles", "wrong_test")
	c.Check(err, NotNil)
	c.Check(err.Error(), Equals, "no reachable servers")
	c.Check(len(data), Equals, 0)
}
