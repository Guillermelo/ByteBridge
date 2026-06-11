package jobs

type Job struct {
	ID     string
	Type   string
	Packet *Packet
}

type Packet struct {
	Type     string `json:"type"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}

type JobsContract interface {
	Execute() error
}
