package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bxcodec/faker"
)

type record struct {
	Raw string
}

func writeRandomLog(fname string) {
	var randRecord record
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		panic(err)
	}

	for {
		faker.FakeData(&randRecord)
		loc, _ := time.LoadLocation("Asia/Tashkent")
		t := time.Now().In(loc)
		_, err := f.WriteString(fmt.Sprintf("%v | %v\n", t.Format("Jan 2, 2006 at 3:04:05pm (UTC)"), randRecord.Raw))
		if err != nil {
			fmt.Println(err.Error())
		}
		time.Sleep(time.Millisecond * 5)
	}
}

func main() {
	if len(os.Args) <= 1 {
		panic("Required files list")
	}
	files := os.Args[1:]
	for _, v := range files {
		if _, err := os.Stat(v); os.IsNotExist(err) {
			os.Create(v)
		}
		go writeRandomLog(v)
	}

	time.Sleep(time.Second * 10)
}
