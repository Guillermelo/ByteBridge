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

func NewJob(p *Packet, reader *bufio.Reader) Job {
	switch p.Type {
	case "SendFileJob":
		return &SendFileJob{
			Size:     p.Size,
			Filename: p.Filename,
			Reader:   reader,
		}
	case "ReceiveFileJob":

		return &ReceiveFileJob{
			Size:     p.Size,
			Filename: p.Filename,
			Reader:   reader,
		}
	case "SendMessage":
		return &SendMessageJob{
			Userdata: p.User_data,
			Reader:   reader,
		}
	default:
		return nil
	}
}
