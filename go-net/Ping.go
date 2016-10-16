package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

//错误检查的方法
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func checkSum(msg []byte) uint16 {
	sum := 0
	//assum even for now
	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}

	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break //跳出循环，处理下一次conn
			}
			return nil, err //读取io出现错误，返回nil
		}
	}
	return result.Bytes(), nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ", os.Args[0], "host")
		os.Exit(1)
	}
	//ip resolver
	addr, err := net.ResolveIPAddr("ip", os.Args[1])
	if err != nil {
		fmt.Println("resolution error", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialIP("ip4:icmp", addr, addr)
	checkError(err)

	var msg [512]byte
	msg[0] = 8  //echo
	msg[1] = 0  //code 0
	msg[2] = 0  //checksum,fix laster
	msg[3] = 0  //checksum
	msg[4] = 0  //identifier
	msg[5] = 13 //identifier[1]
	msg[6] = 0  //sequence[0]
	msg[7] = 37 //sequence[1]
	len := 8

	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	_, err = conn.Read(msg[0:])
	checkError(err)

	fmt.Println("got response")

	if msg[5] == 13 {
		fmt.Println("identifier matches")
	}

	if msg[7] == 37 {
		fmt.Println("sequence matches")
	}

	os.Exit(0)
}
