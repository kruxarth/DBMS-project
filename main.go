package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
)


type ValueType string

const (
	ARRAY	 ValueType = "*"
	BULK 	 ValueType = "$"
	STRING 	 ValueType = "+"
	ERROR	 ValueType = "-"
	NULL 	 ValueType = ""
)




func main(){
	log.Println("reading config files")

	l, err := net.Listen("tcp", ":6379")
	if err != nil{
		log.Fatal("cannot listen on :6379")
	}

	defer l.Close()
	log.Println("listening on :6379")


	conn, err := l.Accept()
	if err!= nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer conn.Close()
	log.Println("connection accepted")

	for{

		v := Value{typ: ARRAY}
		v.readArray(conn)

		handle(conn, &v)

		fmt.Println(v.array)
		

		

	}

}

















