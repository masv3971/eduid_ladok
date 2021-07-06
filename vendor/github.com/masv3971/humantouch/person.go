package humantouch

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	age "github.com/bearbin/go-age"
	//rand2 "github.com/milosgajdos83/go-estimate/rand"
)

const (
	// GenderFemale for female gender
	GenderFemale = "female"
	// GenderMale for male gender
	GenderMale = "male"
	// GenderUnspecified for unspecified gender
	GenderUnspecified = "unspecified"
	yearMin           = 1900
	yearMax           = 2010
)

// PersonConfig configuration for person
type PersonConfig struct {
}

// AgeData has data about each ageDistrubution
type AgeData struct {
	Weight float64
	id     int
}

type personClient struct {
	config *PersonConfig
}

func newPersonClient() (*personClient, error) {
	c := &personClient{}

	return c, nil
}

var months = map[int]int{
	1:  31,
	2:  28,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

// Gender object
type Gender struct {
	General string
	Ladok   int
}

// SocialSecurityNumber holds
type SocialSecurityNumber struct {
	Swedish12   *SwedishNIN
	Swedish10   *SwedishNIN
	SwedishTemp *SwedishNIN
}

// BirthYear data
type BirthYear struct {
	S            string
	I            int
	SLong        string
	ILong        int
	CenturyLongS string
	CenturyLongI int
	CenturyS     string
	CenturyI     int
}

// BirthMonth data
type BirthMonth struct {
	S string
	I int
}

// BirthDay day data
type BirthDay struct {
	S string
	I int
}

// Person object
type Person struct {
	Firstname            string
	Lastname             string
	BirthYear            BirthYear
	BirthMonth           BirthMonth
	BirthDay             BirthDay
	SocialSecurityNumber *SocialSecurityNumber
	Gender               Gender
	Age                  int
}

func (c *personClient) new(gender string) *Person {
	p := &Person{}
	year := c.createYear()

	p.setYear(year)
	p.setMonth()
	p.setDay()
	p.setAge()
	p.setGender(gender)
	p.setName(gender)

	return p
}

func (p *Person) setGender(gender string) {
	gender = randomGender(gender)

	if gender == GenderFemale {
		p.Gender.General = GenderFemale
		p.Gender.Ladok = 2
	} else if gender == GenderMale {
		p.Gender.General = GenderMale
		p.Gender.Ladok = 1
	}
}

func (p *Person) setAge() {
	p.Age = age.Age(time.Date(p.BirthYear.ILong, time.Month(p.BirthMonth.I), p.BirthDay.I, 0, 0, 0, 0, time.UTC))
}

func (c *personClient) createYear() int {
	return rand.Intn(yearMax-yearMin+1) + yearMin
}

func (p *Person) setYear(year int) {
	var err error

	p.BirthYear.ILong = year

	p.BirthYear.SLong = strconv.Itoa(p.BirthYear.ILong)

	p.BirthYear.S = p.BirthYear.SLong[2:]
	p.BirthYear.I, err = strconv.Atoi(p.BirthYear.S)
	if err != nil {
		panic(err)
	}

	p.BirthYear.CenturyS = p.BirthYear.SLong[:2]
	p.BirthYear.CenturyI, err = strconv.Atoi(p.BirthYear.CenturyS)
	if err != nil {
		panic(err)
	}

	p.BirthYear.CenturyLongI = p.BirthYear.CenturyI * 100
	p.BirthYear.CenturyLongS = strconv.Itoa(p.BirthYear.CenturyLongI)
}

func (p *Person) setMonth() {
	p.BirthMonth.I = rand.Intn(month(time.Now().Month())-1) + 1
	if p.BirthMonth.I < 10 {
		p.BirthMonth.S = fmt.Sprintf("0%d", p.BirthMonth.I)
	} else {
		p.BirthMonth.S = strconv.Itoa(p.BirthMonth.I)
	}
}

func (p *Person) setDay() {
	p.BirthDay.I = rand.Intn(months[p.BirthMonth.I]-1) + 1
	if p.BirthDay.I < 10 {
		p.BirthDay.S = fmt.Sprintf("0%d", p.BirthDay.I)
	} else {
		p.BirthDay.S = strconv.Itoa(p.BirthDay.I)
	}
}

func (p *Person) setName(gender string) {
	if gender == GenderFemale {
		p.Firstname = FirstnamesFemale[rand.Intn(len(FirstnamesFemale))]
	} else if gender == GenderMale {
		p.Firstname = FirstnamesMale[rand.Intn(len(FirstnamesMale))]
	}
	p.Lastname = Lastnames[rand.Intn(len(Lastnames))]
}
