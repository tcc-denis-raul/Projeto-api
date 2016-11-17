package db

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Courses struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string
	Based       []string
	PriceReal   []float64 `bson:"price_real"`
	PriceDolar  []float64 `bson:"price_dolar"`
	Dynamic     []string
	Platform    []string
	Url         string
	Extra       []string
	Description string
	Rate        int
}

type CourseScore struct {
	Course Courses
	Score  int
}

type Filter struct {
	Type     string
	Course   string
	Based    string
	Dynamic  string
	Platform string
	Extra    string
	Price    string
	Length   int
}

type TypeCourses struct {
	Language []string
}

type IndicateCourse struct {
	Type   string
	Course string
	Name   string
	Url    string
}

type CourseDetail struct {
	Type   string
	Course string
	Name   string
}

func hasStr(value string, list []string) bool {
	for i := range list {
		if list[i] == value {
			return true
		}
	}
	return false
}

func SortScore(cs []CourseScore) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(cs)-1; i++ {
			if cs[i+1].Score > cs[i].Score {
				cs[i], cs[i+1] = cs[i+1], cs[i]
				swapped = true
			}
		}
	}
}

func check(values, limit []float64, mult float64) bool {
	for _, value := range values {
		if value*mult >= limit[0] && value*mult <= limit[1] {
			return true
		}
	}
	return false
}

func (f *Filter) checkPrice(data Courses, mult float64, dolar bool) bool {
	price_range := map[string][]float64{
		"gratis":  []float64{0.0, 0.0},
		"ate30":   []float64{0.0, 30.0},
		"31a60":   []float64{31.0, 60.0},
		"61a100":  []float64{61.0, 100.0},
		"101a150": []float64{101.0, 150.0},
		"151mais": []float64{151.0, 10000.0},
	}
	limit := price_range[f.Price]
	if !dolar {
		return check(data.PriceReal, limit, mult)
	} else {
		return check(data.PriceDolar, limit, mult)
	}
}

func getDolar() float64 {
	resp, err := http.Get("http://dolarhoje.com/cotacao.txt")
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt := strings.Replace(string(body), ",", ".", 1)
	f, err := strconv.ParseFloat(fmt, 64)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	return f
}

func (f *Filter) filterCourse(data []Courses) []CourseScore {
	dolar := getDolar()
	scored := make([]CourseScore, 0)
	for index := range data {
		score := 0
		if f.Based != "" && hasStr(f.Based, data[index].Based) {
			score++
		}
		if f.Dynamic != "" && hasStr(f.Dynamic, data[index].Dynamic) {
			score++
		}
		if f.Platform != "" && hasStr(f.Platform, data[index].Platform) {
			score++
		}
		if f.Extra != "" && hasStr(f.Extra, data[index].Extra) {
			score++
		}
		if f.Price != "" {
			if f.checkPrice(data[index], 1.0, false) {
				score++
			} else if f.checkPrice(data[index], dolar, true) {
				score++
			}
		}
		scored = append(scored, CourseScore{
			Course: data[index],
			Score:  score,
		})
	}
	SortScore(scored)
	return scored
}

func (f *Filter) limitCourse(courses []CourseScore) []Courses {
	if f.Length > len(courses) {
		f.Length = len(courses)
	}
	var result []Courses
	for i := 0; i < f.Length; i++ {
		result = append(result, courses[i].Course)
	}
	return result
}

func (f *Filter) GetCourses() ([]Courses, error) {
	db, err := GetSession()
	if err != nil {
		return nil, err
	}
	defer db.session.Close()
	var data []Courses
	err = db.session.DB(db.DBName).C(fmt.Sprintf("%s_%s", f.Type, f.Course)).Find(nil).All(&data)
	if err != nil {
		return nil, err
	}
	scoredCourses := f.filterCourse(data)
	if f.Length != 0 {
		data = f.limitCourse(scoredCourses)
	}
	return data, nil
}

func (f *CourseDetail) GetDetailCourse() (Courses, error) {
	db, err := GetSession()
	if err != nil {
		return Courses{}, err
	}
	defer db.session.Close()
	var course Courses
	err = db.session.DB(db.DBName).C(fmt.Sprintf("%s_%s", f.Type, f.Course)).Find(bson.M{"name": f.Name}).One(&course)
	if err != nil {
		return Courses{}, err
	}
	return course, nil
}

func GetTypeCourses() ([]TypeCourses, error) {
	db, err := GetSession()
	if err != nil {
		return nil, err
	}
	defer db.session.Close()
	var data []TypeCourses
	err = db.session.DB(db.DBName).C("courses").Find(nil).All(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (i *IndicateCourse) IndicateCourse() error {
	db, err := GetSession()
	if err != nil {
		return err
	}
	defer db.session.Close()
	var indicate []IndicateCourse
	err = db.session.DB(db.DBName).C("indicate_courses").Find(bson.M{"url": i.Url}).All(&indicate)
	if err != nil {
		return err
	}
	if len(indicate) > 0 && indicate[0].Url == i.Url {
		return errors.New("Course already indicate")
	}
	var courses []Courses
	err = db.session.DB(db.DBName).C(fmt.Sprintf("%s_%s", i.Type, i.Course)).Find(bson.M{"url": i.Url}).All(&courses)
	if err != nil {
		return err
	}
	if len(courses) > 0 && courses[0].Url == i.Url {
		return errors.New("Course already exists on list of courses")
	}
	err = db.session.DB(db.DBName).C("indicate_courses").Insert(i)
	return err
}
