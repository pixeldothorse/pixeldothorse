package horsevilleserver

import (
	"context"

	"github.com/horseville/horseville/rpc/horseville"
)

// Ping implements the ping service
type Ping struct{}

// Message replies to a ping with a pong.
func (p Ping) Message(ctx context.Context, n *horseville.Nil) (*horseville.Nil, error) {
	return n, nil
}
