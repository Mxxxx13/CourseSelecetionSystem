// @Title : rabbitmq
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2022/5/8 22:38

package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"CourseSeletionSystem/dao"
	"CourseSeletionSystem/model"
	"github.com/streadway/amqp"
)

// MQURL 连接信息
const MQURL = "amqp://root:123456@127.0.0.1:5672/rabbitmq_test"

// RabbitMQ rabbitMQ结构体
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind Key 名称
	Key string
	//连接信息
	Mqurl string
}

// NewRabbitMQ 创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}

// Destory 断开channel 和 connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// NewRabbitMQWork 创建work模式下RabbitMQ实例
func NewRabbitMQWork(queueName string) *RabbitMQ {
	//创建RabbitMQ实例
	rabbitmq := NewRabbitMQ(queueName, "", "")
	var err error
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "failed to connect rabbitmq!")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel")
	return rabbitmq
}

// PublishWork work模式队列生产
func (r *RabbitMQ) PublishWork(message string) {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//调用channel 发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// ConsumeWork  work模式下消费者
func (r *RabbitMQ) ConsumeWork() {
	for i := 0; i < 5; i++ {
		go func() {
			//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
			q, err := r.channel.QueueDeclare(
				r.QueueName,
				//是否持久化
				false,
				//是否自动删除
				false,
				//是否具有排他性
				false,
				//是否阻塞处理
				false,
				//额外的属性
				nil,
			)
			if err != nil {
				fmt.Println(err)
			}

			//接收消息
			msgs, err := r.channel.Consume(
				q.Name, // queue
				//用来区分多个消费者
				"", // consumer
				//是否自动应答
				false, // auto-ack
				//是否独有
				false, // exclusive
				//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
				false, // no-local
				//列是否阻塞
				false, // no-wait
				nil,   // args
			)
			if err != nil {
				fmt.Println(err)
			}

			//启用协程处理消息
			go func() {
				for d := range msgs {
					//消息逻辑处理，可以自行设计逻辑
					//log.Printf("Received a message: %s", d.Body)
					selection := model.CourseSelection{}
					err := json.Unmarshal(d.Body, &selection)
					if err != nil {
						log.Printf("unmarshal failed")
						continue
					}
					err = dao.StudentSelectCourse(selection)
					if err != nil {
						log.Println(err)
						continue
					}
					err = dao.CourseStuNumAddOne(selection.CourseID)
					if err != nil {
						log.Println(err)
						continue
					}
					//如果为true表示确认所有未确认的消息，
					//为false表示确认当前消息
					_ = d.Ack(false)
				}
			}()
		}()
	}

	forever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
