package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/tarm/serial"
)

// 写入程序
func WriteSerial(s *serial.Port) {
	// 数据计数
	dataCount := 0
	for {
		// 数据内容
		data := fmt.Sprintf("This is the %v data", dataCount)
		buf := []byte(data)
		// 写入数据
		n, err := s.Write([]byte(buf))
		if err != nil {
			log.Fatal(err)
		}
		// 打印写入数据内容
		log.Printf("Write: %v , Size: %v", data, n)
		// 数据计数+1
		dataCount += 1
		// 暂停1秒
		time.Sleep(time.Second * 1)
	}
}

// 程序入口
func main() {
	// 指定写入的串口和波特率
	writeConfig := &serial.Config{
		Name: "/dev/pts/2",
		Baud: 9600,
		// ReadTimeout: time.Second * 5,
	}
	// 使用指定的串口和波特率打开写入串口
	writeSerialPort, err := serial.OpenPort(writeConfig)
	if err != nil {
		log.Fatal(err)
	}
	// 使用协程执行写入程序
	go WriteSerial(writeSerialPort)
	// 阻塞程序
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
