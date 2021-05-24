package ex6

import (
	"io/ioutil"
	"time"
)

type UUIDGenerator interface {
	New() string
}

type FTimeGenerator interface {
	Now() time.Time
}

type Writer struct {
	uuid  UUIDGenerator
	ftime FTimeGenerator
}

func New(uuidgen UUIDGenerator, ftimegen FTimeGenerator) *Writer {
	return &Writer{
		uuid:  uuidgen,
		ftime: ftimegen,
	}
}

func (writer *Writer) Write() error {
	u := writer.uuid.New()
	ft := writer.ftime.Now().Format(time.RFC3339)
	return ioutil.WriteFile(u+".txt", []byte(ft), 0644)
}
