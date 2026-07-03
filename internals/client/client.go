package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"ByteBridge/internals/jobs"
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

	defer func() {
		if err := c.Conn.Close(); err != nil {
			log.Println(err)
		}
	}()
}

func (c *Client) TestPacket(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return err
	}

	packet := jobs.ReceiveFileJob{
		Type:     "ReceiveFileJob",
		Size:     info.Size(),
		Filename: info.Name(),
		Userdata: "testJob",
	}

	jsonData, err := json.Marshal(&packet)
	if err != nil {
		return err
	}

	jsonData = append(jsonData, '\n')
	_, err = c.Conn.Write(jsonData)
	if err != nil {
		return err
	}

	fmt.Println("json sent", string(jsonData))

	if _, err := io.Copy(c.Conn, file); err != nil {
		return err
	}

	fmt.Println("file sent", jsonData)
	return nil
}
