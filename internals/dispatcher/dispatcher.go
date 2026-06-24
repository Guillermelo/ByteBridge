// Package dispatcher should send filter jobs and complete them
package dispatcher

import (
	"context"
	"fmt"

	"ByteBridge/internals/jobs"
	"ByteBridge/internals/serverconn"
)

type Dispatcher struct {
	ConnPool chan *serverconn.ServerConn
}

func (d *Dispatcher) Dispatch(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		fmt.Println("in dispatcher resolving Connections")
		newctx, cancel := context.WithCancel(ctx)
		for serverconn := range d.ConnPool {
			go jobs.FillConnJobs(newctx, serverconn.Queue, serverconn.Conn)
			go jobs.FlushConnJobs(newctx, serverconn.Queue, serverconn.Conn)
		}
		defer cancel()
	}

	return nil
}
