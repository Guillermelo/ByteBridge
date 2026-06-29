package jobs

import "bufio"

type Packet struct {
	User_data string `json:"user_data"`
	Type      string `json:"type"`
	Filename  string `json:"filename"`
	Size      int64  `json:"size"`
}

type Job interface {
	Execute() error
}

func NewJob(p *Packet, rw *bufio.ReadWriter) Job {
	switch p.Type {
	case "SendFileJob":
		return &SendFileJob{
			Writer:   rw.Writer,
			Filepath: "/random/path/for/now",
		}
	case "ReceiveFileJob":

		return &ReceiveFileJob{
			Size:     p.Size,
			Filename: p.Filename,
			Reader:   rw.Reader,
		}
	case "SendMessage":
		return &SendMessageJob{
			Userdata: p.User_data,
			Reader:   rw.Writer,
		}
	default:
		return nil
	}
}
