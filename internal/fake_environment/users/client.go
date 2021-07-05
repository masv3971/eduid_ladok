package users

import (
	"github.com/google/uuid"
	"github.com/masv3971/humantouch"
)

// Client holds the client object for users
type Client struct {
	Ladok map[string]*Student
	Eduid map[string]*Student
}

// Student is the student account object
type Student struct {
	Firstname  string
	Lastname   string
	NIN        string
	StudentUID string
}

// New creates a new instance of users
func New() (*Client, error) {
	var err error

	c := &Client{}

	c.Ladok, err = c.createUsers()
	c.Eduid, err = c.createUsers()

	return c, nil
}

func (c *Client) createUsers() (map[string]Student, error) {
	students := make(map[string]*humantouch.Person)

	human, err := humantouch.New(&humantouch.Config{
		DistrubutionCFG: &humantouch.DistributionCfg{
			Age0to10: humantouch.AgeData{
				Weight: 0,
			},
			Age10to20: humantouch.AgeData{
				Weight: 0,
			},
			Age20to30: humantouch.AgeData{
				Weight: 100,
			},
			Age30to40: humantouch.AgeData{
				Weight: 75,
			},
			Age40to50: humantouch.AgeData{
				Weight: 10,
			},
			Age50to60: humantouch.AgeData{
				Weight: 20,
			},
			Age60to70:   humantouch.AgeData{},
			Age70to80:   humantouch.AgeData{},
			Age80to90:   humantouch.AgeData{},
			Age90to100:  humantouch.AgeData{},
			Age100to110: humantouch.AgeData{},
		},
	})
	if err != nil {
		return nil, err
	}

	persons, err := human.Distribution.RandomHumans(10000)
	if err != nil {
		return nil, err
	}

	for _, person := range persons {
		uid := uuid.NewString()

		student := &Student{
			Firstname:  person.Firstname,
			Lastname:   person.Lastname,
			NIN:        person.SocialSecurityNumber.Swedish10.Complete,
			StudentUID: uid,
		}
		students[uid] = student
	}

	return m, nil
}
