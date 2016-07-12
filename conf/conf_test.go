package conf

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type ConfTest struct{}

var _ = Suite(&ConfTest{})

func (s *ConfTest) TestConf(c *C) {
	conf, err := Conf("data_test")
	c.Check(err, IsNil)
	c.Check(conf.URL, Equals, "url")
	c.Check(conf.Name, Equals, "dbname")
}

func (s *ConfTest) TestConfWrongPath(c *C) {
	conf, err := Conf("data")
	c.Check(err, NotNil)
	c.Check(conf.URL, Equals, "")
	c.Check(conf.Name, Equals, "")
}
