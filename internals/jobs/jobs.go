// Package jobs is for dividing the work dpending on the type
package jobs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

func FlushConnJobs(Queue <-chan Job, Conn net.Conn) {
	for job := range Queue {
		job.Execute()
	}
}

func FillConnJobs(Queue chan<- Job, conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		header, _ := reader.ReadBytes('\n')
		var CurrentPacket Packet
		err := json.Unmarshal(header, &CurrentPacket)
		if err != nil {
			fmt.Println(err)
			err = nil
		}
		currentJob := NewJob(&CurrentPacket, reader)
		Queue <- currentJob

	}
}
