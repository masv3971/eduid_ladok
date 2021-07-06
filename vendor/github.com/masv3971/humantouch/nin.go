package humantouch

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type ninClient struct {
	store *storeClient
}

func newNINClient() (*ninClient, error) {
	var err error
	c := &ninClient{}

	c.store, err = newStoreClient()
	if err != nil {
		return nil, err
	}

	return c, nil
}

// BirthNumber data
type BirthNumber struct {
	N1s      string
	N2s      string
	N3s      string
	N1i      int
	N2i      int
	N3i      int
	Complete string
}

func (c *ninClient) birthNumber(gender string) BirthNumber {
	bn := BirthNumber{}
	bn.N1i = rand.Intn(9)
	bn.N2i = rand.Intn(9)
	bn.N1s = strconv.Itoa(bn.N1i)
	bn.N2s = strconv.Itoa(bn.N2i)

	// this can be stuck, theoretical at least.
	if gender == GenderFemale {
		bn.N3i = []int{0,2,4,6,8}[rand.Intn(5)]
	} else {
		bn.N3i = []int{1,3,5,7,9}[rand.Intn(5)]
	}

	bn.N3s = strconv.Itoa(bn.N3i)
	bn.Complete = fmt.Sprintf("%d%d%d", bn.N1i, bn.N2i, bn.N3i)

	return bn
}

// LuhnNumber data
type LuhnNumber struct {
	S string
	I int
}

func (s *SwedishNIN) luhn(p *Person) {
	pn := []string{
		p.BirthYear.S,
		p.BirthMonth.S,
		p.BirthDay.S,
		s.BirthNumber.Complete,
	}

	c := strings.Join(pn, "")

	sum := 0
	for i, n := range strings.Split(c, "") {
		m, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		if i%2 == 0 {
			for _, l := range strings.Split(strconv.Itoa(m*2), "") {
				a, err := strconv.Atoi(l)
				if err != nil {
					panic(err)
				}
				sum = sum + a
			}
		} else {
			sum = sum + m
		}
	}

	s.LuhnNumber.I = (10 - (sum % 10)) % 10
}

// SwedishNIN is the object to store swedish socialnumber
type SwedishNIN struct {
	BirthNumber BirthNumber
	LuhnNumber  LuhnNumber
	Complete    string
	Delimiter   string
}

func (s *SwedishNIN) delimiter(age int) {
	if age > 99 {
		s.Delimiter = "+"
	} else {
		s.Delimiter = "-"
	}
}

func (s *SwedishNIN) setComplet(p *Person) {
	s.Complete = fmt.Sprintf("%s%s%s%s%s%d", p.BirthYear.S, p.BirthMonth.S, p.BirthDay.S, s.Delimiter, s.BirthNumber.Complete, s.LuhnNumber.I)
}

func (s *SwedishNIN) setComplete12(p *Person) {
	s.Complete = fmt.Sprintf("%s%s%s%s%s%d", p.BirthYear.SLong, p.BirthMonth.S, p.BirthDay.S, s.Delimiter, s.BirthNumber.Complete, s.LuhnNumber.I)
}

//func (c *ninClient) femaleSwedish10(person *Person) *SwedishNIN {
func (c *ninClient) newSwedish(person *Person) *SocialSecurityNumber {
	ssn := &SocialSecurityNumber{}

	s10 := SwedishNIN{}
	makeSSN := func() {
		s10.BirthNumber = c.birthNumber(person.Gender.General)
		s10.luhn(person)
		s10.delimiter(person.Age)
		s10.setComplet(person)
	}
	for i := 1; i <= 5; i++ {
		if c.store.exists(s10.Complete) || s10.Complete == "" {
			makeSSN()
		} else {
			ssn.Swedish10 = &s10
			if err := c.store.add(s10.Complete); err != nil {
				panic(err)
			}
			break
		}
	}

	s12 := SwedishNIN{
		BirthNumber: s10.BirthNumber,
		LuhnNumber:  s10.LuhnNumber,
		Delimiter:   s10.Delimiter,
	}
	s12.setComplete12(person)
	ssn.Swedish12 = &s12

	return ssn
}
