package awsmiscclient

import (
	"github.com/STollenaar/aws-misc-client/lib/profiles"
)

// Client -
type Client struct {
	profiler *profiles.ProfileClient
}

// NewClient -
func NewClient() (*Client, error) {
	c := Client{
		profiler: profiles.NewClient(),
	}

	return &c, nil
}
