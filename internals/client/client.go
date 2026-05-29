package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	Addr string
	Conn net.Conn
}

type Packet struct {
	Type     string `json:"type"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}

func (c *Client) Connect() {
	conn, err := net.Dial("tcp", c.Addr)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	c.Conn = conn
}

func (c *Client) Send(info string) {
	_, err := c.Conn.Write([]byte(info))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("sent a message")
	defer c.Close()
}

func (c *Client) Close() {
	c.Conn.Close()
}

func (c *Client) TestPacket() {
	filepath := "./test.png"
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	packet := Packet{
		Type:     "ReceiveFile",
		Filename: info.Name(),
		Size:     info.Size(),
	}
	jsonData, _ := json.Marshal(packet)
	jsonData = append(jsonData, '\n')

	c.Conn.Write(jsonData)
	fmt.Println("json sent", string(jsonData))

	io.Copy(c.Conn, file)

	fmt.Println("file sent", jsonData)
}
