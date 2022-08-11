package awsmiscclient

import (
	"github.com/STollenaar/aws-misc-client/lib/profiles"
)

// Client -
type Client struct {
	Profiler *profiles.ProfileClient
}

// NewClient -
func NewClient() (*Client, error) {
	c := Client{
		Profiler: profiles.NewClient(),
	}

	return &c, nil
}
