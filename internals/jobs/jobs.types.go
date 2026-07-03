package jobs

import (
	"bufio"
	"encoding/json"
)

type Job interface {
	Execute() error
}

type Header struct {
	Type string `json:"type"`
}

func ReturnJob(data []byte, rw bufio.ReadWriter) Job {
	var h Header
	err := json.Unmarshal(data, &h)
	if err != nil {
		return nil
	}
	switch h.Type {
	case "ReceiveFileJob":
		var j *ReceiveFileJob
		err = json.Unmarshal(data, &j)
		j.Reader = rw.Reader
		return j
	case "SendFileJob":
		var j *SendFileJob
		j.Writer = rw.Writer
		return j
	default:
		return nil

	}
}
