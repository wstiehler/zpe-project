package worker

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/wstiehler/zpecreateuser-api/internal/environment"
	"go.uber.org/zap"
)

func ListenToNats(input Input, consumer Consumer) {

	logger := input.Logger

	env := environment.GetInstance()

	nc, err := nats.Connect(env.NATS_URL)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}

	defer nc.Close()

	subject := consumer.QueueSubject()

	_, err = nc.Subscribe(subject, func(msg *nats.Msg) {
		var message UserEntity

		if err := json.Unmarshal(msg.Data, &message); err != nil {
			logger.Error("Error decoding message: %v", zap.String("Error", err.Error()))
			return
		}
		logger.Info("Received message: %s", zap.ByteString("Message", msg.Data))
		consumer.Handler(input, message)
	})

	if err != nil {
		logger.Fatal("Error subscribing to subject %s: %v", zap.String("Message", subject), zap.String("Error", err.Error()))
	}

	logger.Info("Worker subscribed to subject: ", zap.String("Message", subject))

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	logger.Debug("Worker shutting down...")
	time.Sleep(time.Second)
	logger.Debug("Worker shutdown complete")

}
