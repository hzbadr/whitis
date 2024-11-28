package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	// starts with $
// 	// 5 is the length
// 	// \r\n is a separator
// 	// Ahmed is the value to store
// 	input := "$5\r\nAhemd\r\n"
// 	reader := bufio.NewReader(strings.NewReader(input))

// 	b, _ := reader.ReadByte()

// 	if b != '$' {
// 		fmt.Println("invalid type, expecting bulk strings only")
// 		os.Exit(1)
// 	}

// 	size, _ := reader.ReadByte()

// 	strSize, _ := strconv.ParseInt(string(size), 10, 64)

// 	reader.ReadByte()
// 	reader.ReadByte()

// 	name := make([]byte, strSize)
// 	reader.Read(name)

// 	fmt.Println(string(name))
// }
