package main

import (
	"fmt"

	"ByteBridge/internals/client"
)

func main() {
	fmt.Println("Starting Byte-Bridge Client")

	Client := &client.Client{
		Addr: ":4000",
	}
	Client.Connect()
	info := string("aint working here")
	Client.TestPacket()
	Client.Send(info)
	select {}
}
