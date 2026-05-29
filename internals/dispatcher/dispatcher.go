package dispatcher

import (
	"fmt"

	"ByteBridge/internals/jobs"
	"ByteBridge/internals/serverconn"
)

type Dispatcher struct {
	ConnPool chan *serverconn.ServerConn
}

func (d *Dispatcher) Dispatch() {
	fmt.Println("in dispatcher resolving Connections")
	for serverconn := range d.ConnPool {
		go jobs.FillConnJobs(serverconn.Queue, serverconn.Conn)
		go jobs.FlushConnJobs(serverconn.Queue, serverconn.Conn)
	}
}
