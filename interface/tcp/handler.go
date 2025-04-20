package tcp

import (
	"context"
	"net"
)

/*
*
代表redis的业务引擎
*/
type Handler interface {
	Handle(ctx context.Context, conn net.Conn)
	Close() error
}
