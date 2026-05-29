package serverconn

import (
	"fmt"
	"net"

	"ByteBridge/internals/jobs"
)

const MaxJobPerConn int = 100

type ServerConn struct {
	CID   string
	Queue chan jobs.Job
	Conn  net.Conn
}

// func (c *ServerConn) FlushConnJobs() {
// 	for j := range c.Queue {
// 		switch j.Type {
// 		case "Message":
//
// 		case "File":
// 		}
// 	}
// }

func FillConnPool(Addr string, ConnPool chan<- *ServerConn) {
	AssignCID := 1
	listener, err := net.Listen("tcp", Addr)
	if err != nil {
		fmt.Println("Listener Err")
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Conn Err", err)
			continue
		}
		fmt.Println("new Client Connected")
		if conn != nil {
			AssignCID++
			CID := fmt.Sprintf("%d", AssignCID)
			Connection := &ServerConn{
				CID:   CID,
				Queue: make(chan jobs.Job, MaxJobPerConn),
				Conn:  conn,
			}
			ConnPool <- Connection
		}
		// aca tendrain que guardar en el chan de conn? socketlist <- net.conn?
	}
}
