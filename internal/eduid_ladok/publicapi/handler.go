package publicapi

// RequestPublic request
type RequestPublic struct{}

// ReplyPublic reply
type ReplyPublic struct{}

// Public handler
func (c *Client) Public(indata *RequestPublic) (*ReplyPublic, error) {
	return nil, nil
}
