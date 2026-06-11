package main

import (
	"flag"
	"fmt"

	"ByteBridge/internals/client"
)

func main() {
	fmt.Println("Starting Byte-Bridge Client")
	file := flag.String("file", "./test.png", "file to send")
	flag.Parse()
	Client := &client.Client{
		Addr: ":4000",
	}
	Client.Connect()
	// info := string("aint working here")
	err := Client.TestPacket(*file)
	if err != nil {
		fmt.Println("something happened ahh: ", err)
	}
	// Client.Send(info)
	select {}
}
