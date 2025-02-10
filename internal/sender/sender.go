package sender

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
	"github.com/osamikoyo/hrm-vocation/internal/data/models"
	"github.com/osamikoyo/hrm-vocation/pkg/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Sender struct{
	AmqpChannel *amqp.Channel
	AmqpQue amqp.Queue
}

func Init(cfg *config.Config) (*Sender, error) {
	conn, err := amqp.Dial(cfg.RabbitMqURl)
	if err != nil{
		return nil, err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil{
		return nil, err
	}
	defer ch.Close()

	que, err := ch.QueueDeclare(
		"message",
		false, 
		false,
		false,
		false,
		nil,
	)

	return &Sender{
		AmqpQue: que,
		AmqpChannel: ch,
	}, err
}

func (s *Sender) Send(message models.Msg) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := sonic.Marshal(message)
	if err != nil{
		return err
	}

	err = s.AmqpChannel.PublishWithContext(ctx,
  		"",     // exchange
  		s.AmqpQue.Name, // routing key
  		false,  // mandatory
  		false,  // immediate
  		amqp.Publishing {
    		ContentType: "application/json",
    		Body:        []byte(body),
  	})

	if err != nil{
		return err
	}

	return nil
}