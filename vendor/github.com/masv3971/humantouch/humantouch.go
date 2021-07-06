package humantouch

import (
	"errors"
)

var (
	// ErrAgeDistrubutionNotConfigured error about not configured age distrubution
	ErrAgeDistrubutionNotConfigured = errors.New("ERR_AGE_DISTRUBUTION_NOT_CONFIGURED")
	// ErrKeyCollide error for a key that collide
	ErrKeyCollide = errors.New("ERR_KEY_COLLIDE")
	// ErrNoKey error for key that does not exists
	ErrNoKey = errors.New("ERR_NO_KEY")
)

// Client holds humantouch object
type Client struct {
	nin          *ninClient
	person       *personClient
	Distribution *Distribution
}

// Config holds configuration for humantouch
type Config struct {
	//Person *PersonConfig
	DistrubutionCFG *DistributionCfg
}

// New creates a new instance of humantouch
func New(config *Config) (*Client, error) {
	var err error
	c := &Client{}

	c.nin, err = newNINClient()
	if err != nil {
		return nil, err
	}

	if config == nil {
		c.Distribution, err = newDistributionClient(&Config{})
		if err != nil {
			return nil, err
		}
	} else {
		c.Distribution, err = newDistributionClient(config)
		if err != nil {
			return nil, err
		}
	}

	c.person, err = newPersonClient()
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Female return a female person, or error
func (c *Client) Female() (*Person, error) {
	person := c.person.new(GenderFemale)

	person.SocialSecurityNumber = c.nin.newSwedish(person)
	return person, nil
}

// Male return a male person, or error
func (c *Client) Male() (*Person, error) {
	person := c.person.new(GenderMale)

	person.SocialSecurityNumber = c.nin.newSwedish(person)
	return person, nil
}

// RandomHuman return all kinds of humans, full span of age and gender
func (c *Client) RandomHuman() (*Person, error) {
	switch r := randomGender(""); r {
	case GenderFemale:
		human, err := c.Female()
		if err != nil {
			return nil, err
		}
		return human, nil
	case GenderMale:
		human, err := c.Male()
		if err != nil {
			return nil, err
		}
		return human, nil
	}

	return nil, nil
}

// Females return n slice of *Person
func (c *Client) Females(n int) ([]*Person, error) {
	persons := []*Person{}
	for i := 1; i <= n; i++ {
		human, err := c.Female()
		if err != nil {
			return nil, err
		}
		persons = append(persons, human)
	}
	return persons, nil
}

// Males return n slice of *Person
func (c *Client) Males(n int) ([]*Person, error) {
	persons := []*Person{}
	for i := 1; i <= n; i++ {
		human, err := c.Male()
		if err != nil {
			return nil, err
		}
		persons = append(persons, human)
	}
	return persons, nil
}

// RandomHumans return a slice of female humans
func (c *Client) RandomHumans(n int) ([]*Person, error) {
	persons := []*Person{}
	for i := 1; i <= n; i++ {
		human, err := c.RandomHuman()
		if err != nil {
			return nil, err
		}
		persons = append(persons, human)
	}
	return persons, nil
}
