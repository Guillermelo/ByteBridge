package jobs

type Job struct {
	ID      string
	Payload []byte
	Type    string
}

type Packet struct {
	Type     string `json:"type"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}

type JobsContract interface {
	Execute() error
}
