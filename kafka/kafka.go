package kafka

import (
	logs "github.com/Sirupsen/logrus"
	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer
)


func InitKafka(addr string) (err error){
	logs.Info("init kafka start")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logs.Error("init kafka producer failed, err:", err)
		return
	}

	logs.Debug("init kafka succ")
	return
}

func SendToKafka(data, topic string) (err error) {

	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	_, _, err = client.SendMessage(msg)
	if err != nil {
		logs.Error("send message failed, err:%v data:%v topic:%v", err, data, topic)
		return
	}

	//logs.Debug("send succ, pid:%v offset:%v, topic:%v\n", pid, offset, topic)
	return
}