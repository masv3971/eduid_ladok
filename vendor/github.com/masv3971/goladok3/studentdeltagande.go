package goladok3

// studentdeltagandeService xx
type studentdeltagandeService struct {
	client  *Client
	service string
}

func (s *studentdeltagandeService) acceptHeader() string {
	return ladokAcceptHeader[s.service][s.client.format]
}
