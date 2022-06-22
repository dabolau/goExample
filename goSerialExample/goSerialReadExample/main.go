package main

import (
	"log"
	"sync"

	"github.com/tarm/serial"
)

// 读取程序
func ReadSerial(s *serial.Port) {
	for {
		// 指定数据大小
		buf := make([]byte, 1024)
		// 读取数据
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		// 打印读取数据内容
		log.Printf("Read: %v", string(buf[:n]))
	}
}

// 程序入口
func main() {
	// 指定读取串口和波特率
	readConfig := &serial.Config{
		Name: "/dev/pts/3",
		Baud: 9600,
		// ReadTimeout: time.Second * 5,
	}
	// 使用指定的串口和波特率打开读取串口
	readSerialPort, err := serial.OpenPort(readConfig)
	if err != nil {
		log.Fatal(err)
	}
	// 使用协程执行读取程序
	go ReadSerial(readSerialPort)
	// 阻塞程序
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
