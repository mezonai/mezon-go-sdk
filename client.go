package client

import (
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/johannesridho/protobuf-over-tcp/message"
)

func Authenticate() {

}

func sendMessage(conn net.Conn) {
	log.Println("client connected")

	defer conn.Close()

	messageProto := message.Message{Text: "Hello World", Timestamp: time.Now().Unix()}
	data, err := proto.Marshal(&messageProto)
	checkError(err)

	length, err := conn.Write(data)
	checkError(err)

	log.Printf("Hello world sent, length %d bytes", length)
}

func startClient() {
	log.Println("starting tcp client")

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	checkError(err)

	defer conn.Close()

	data := make([]byte, 4096)
	length, err := conn.Read(data)
	checkError(err)

	messagePb := message.Message{}
	err = proto.Unmarshal(data[:length], &messagePb)
	checkError(err)

	log.Printf("received message: %s, timestamp: %v", messagePb.Text, messagePb.Timestamp)
}

func checkError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

//go:generate protoc --go_out=. --go-grpc_out=. -I=./message ./message/message.proto
