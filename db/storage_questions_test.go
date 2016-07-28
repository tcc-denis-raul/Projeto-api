package db

import (
	. "gopkg.in/check.v1"
)

func (s *StorageTest) TestGetQuestionsEmptyList(c *C) {
	data, err := GetQuestions("language", "data_test")
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
	data, err := GetQuestions("language", "data_test")
	c.Check(err, IsNil)
	c.Check(len(data), Equals, 1)
	c.Check(data[0].Based, DeepEquals, questions[0].Based)
	c.Check(data[0].Price, DeepEquals, questions[0].Price)
	c.Check(data[0].Dynamic, DeepEquals, questions[0].Dynamic)
	c.Check(data[0].Platform, DeepEquals, questions[0].Platform)
	c.Check(data[0].Extra, DeepEquals, questions[0].Extra)
}
