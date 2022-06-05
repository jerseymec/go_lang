package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var wc WriterCloser = NewBufferedWriterCloser()
	_, err := wc.Write([]byte("Hello MyDear Fellow Americans, this is an important message"))
	if err != nil {
		log.Fatalln("Error while writer to std", err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatalln("Error closing: ", err)
	}

}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(msg []byte) (int, error) {
	n, err := bwc.buffer.Write(msg)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil

}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil

}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
