package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	BROKER    = "tcp://66.42.63.94:1883" // 服务器地址和端口
	CLIENT_ID = "LINE2-000010"           // 客户端名称
	TOPIC     = "demo/1"                 // 订阅主题
)

// 消息日志
func mqttLog() {
	mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	mqtt.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
	mqtt.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	mqtt.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)
}

// 消息回调
func mqttMessageHandler(c mqtt.Client, msg mqtt.Message) {
	log.Printf("TOPIC: %s MESSAGE: %s\n", msg.Topic(), msg.Payload())
}

// 订阅主题
func mqttSubscribe(c mqtt.Client) {
	// 订阅主题
	if token := c.Subscribe("demo/1", 2, nil); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}
}

// 发布消息
func mqttPublish(c mqtt.Client) {
	count := 0
	for {
		text := fmt.Sprintf("This is the %d message", count)
		// 发布消息
		if token := c.Publish("demo/1", 2, false, text); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
		}
		count += 1
		time.Sleep(1 * time.Second)
	}
}

// 程序入口
func main() {
	// 日志消息
	mqttLog()
	// 定义客户端选项
	opts := mqtt.NewClientOptions()
	// 设置服务器
	opts.AddBroker(BROKER)
	// 设置客户端名称
	opts.SetClientID(CLIENT_ID)
	// 设置心跳时间
	opts.SetKeepAlive(60 * time.Second)
	// 是否清除会话和离线消息
	opts.SetCleanSession(false)
	// 设置消息回调
	opts.SetDefaultPublishHandler(mqttMessageHandler)
	// 创建客户端
	c := mqtt.NewClient(opts)
	// 连接服务器
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}
	// 订阅主题
	go mqttSubscribe(c)
	// 发布信息
	go mqttPublish(c)
	// 阻塞程序
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
