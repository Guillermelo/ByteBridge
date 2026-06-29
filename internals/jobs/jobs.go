// Package jobs is for dividing the work dpending on the type
package jobs

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net"
)

func FlushConnJobs(Queue chan Job, Conn net.Conn) {
	for job := range Queue {
		err := job.Execute()
		if err != nil {
			fmt.Println("error in go routine executtion FlushConnJobs ", err)
		}
	}
}

func FillConnJobs(ctx context.Context, Queue chan Job, conn net.Conn) {
	select {
	case <-ctx.Done():
		return
	default:
		reader := bufio.NewReader(conn)
		writter := bufio.NewWriter(conn)
		readerWriter := bufio.NewReadWriter(reader, writter)

		for {
			header, err := reader.ReadBytes('\n')
			if err != nil {
				fmt.Println("connection closed or read error: ", err)
				return
			}
			var CurrentPacket Packet
			fmt.Println("printing the header: ?")
			fmt.Println(header)
			err = json.Unmarshal(header, &CurrentPacket)
			if err != nil {
				fmt.Println(err)
				return
			}
			currentJob := NewJob(&CurrentPacket, readerWriter) // not the reader this is wrong
			if currentJob != nil {
				Queue <- currentJob
				FlushConnJobs(Queue, conn)
			}
		}
	}
}
