package awsmisc

// Client -
type Client struct{}

// NewClient -
func NewClient() (*Client, error) {
	c := Client{}

	return &c, nil
}
