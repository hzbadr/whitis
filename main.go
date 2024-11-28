package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Listening in port :6379")

	// Create a new server
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Listen to connections
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		resp := NewResp(conn)

		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		if value.typ != "array" {
			fmt.Println("invalid request, expected array got ", value.typ)
			return
		}

		if len(value.array) == 0 {
			fmt.Println("invalid request, expected array length > 0")
			return
		}

		command := strings.ToUpper(value.array[0].bulk)

		args := value.array[1:]

		writer := NewWriter(conn)

		handler, ok := Handlers[command]

		if !ok {
			fmt.Println("invalid command: ", value.array)
			writer.Write(Value{typ: "string", str: ""})
			continue
		}

		result := handler(args)
		writer.Write(result)
	}
}
