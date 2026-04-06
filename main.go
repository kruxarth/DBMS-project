package main

import (

	"fmt"
	
	"log"
	"net"
	"os"
	
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
	conf := readConf("./redis.conf")

	state := NewAppState(conf)

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

		handle(conn, &v, state)

		fmt.Println(v.array)
		

		

	}

}



type AppState struct{
	conf *Config
	aof *Aof
}

func NewAppState(conf *Config)*AppState{
	state:= AppState{
		conf: conf,
	}

	if conf.aofEnabled{
		state.aof = NewAof(conf)
	}

	return &state
}














