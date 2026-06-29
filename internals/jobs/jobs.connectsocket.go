package jobs

type ConnectToSocket struct {
	UserName string
	Adress   string
}

func (s ConnectToSocket) Execute() error {
	//
	return nil
}
