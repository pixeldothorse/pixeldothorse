package pixeldothorseserver

import (
	"context"

	"github.com/pixeldothorse/pixeldothorse/rpc/pixeldothorse"
)

// Ping implements the ping service
type Ping struct{}

// Message replies to a ping with a pong.
func (p Ping) Message(ctx context.Context, n *pixeldothorse.Nil) (*pixeldothorse.Nil, error) {
	return n, nil
}
