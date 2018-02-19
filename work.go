package main

import (
"github.com/iyaoyao/logagent/tailf"
"github.com/iyaoyao/logagent/kafka"
"time"
logs "github.com/Sirupsen/logrus"
)


func Run() (err error){
	for {
		msg := tailf.GetOneLine()
		err = sendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed, err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

func sendToKafka(msg *tailf.TextMsg)(err error) {
	//fmt.Printf("read msg:%s, topic:%s\n", msg.Msg, msg.Topic)
	err = kafka.SendToKafka(msg.Msg, msg.Topic)
	return
}