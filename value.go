package main

import (
	"fmt"
	"io"
	"strconv"
)

type Value struct {
	typ ValueType
	bulk string
	str string
	array []Value
	err string
}


func (v * Value)readArray(reader io.Reader){
	buf := make([]byte, 4)
	reader.Read(buf)
	arrLen, err := strconv.Atoi(string(buf[1]))
	if err!= nil {
		fmt.Println(err)
		return
	}

	for range arrLen {
		bulk := v.readBulk(reader)
		v.array = append(v.array, bulk)
	}
}


func (v * Value) readBulk(reader io.Reader) Value{
	buf := make([]byte, 4)
	reader.Read(buf)

	n, err := strconv.Atoi(string(buf[1]))
	if err != nil {
		fmt.Println(err)
		return Value{}
	}

	bulkBuf := make([]byte, n+2)
	reader.Read(bulkBuf)

	theWord := string(bulkBuf[:n])
	return Value{typ: BULK, bulk: theWord}
}

