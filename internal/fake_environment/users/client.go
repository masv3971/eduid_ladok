package users

import (
	"eduid_ladok/pkg/logger"
)

// Client holds the client object for users
type Client struct {
	logger *logger.Logger
	Ladok  map[string]*Student
	Eduid  map[string]*Student
}

// Student is the student account object
type Student struct {
	Firstname  string
	Lastname   string
	NIN        string
	StudentUID string
}

// New creates a new instance of users
func New(logger *logger.Logger) (*Client, error) {
	var err error

	c := &Client{
		logger: logger,
	}

	c.Ladok, err = c.createUsers()
	if err != nil {
		return nil, err
	}

	c.Eduid, err = c.createUsers()
	if err != nil {
		return nil, err
	}

	return c, nil
}
