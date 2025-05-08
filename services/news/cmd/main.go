package main

import (
	"context"
	"fmt"
	"github.com/crypto-pulse/news/internal/integration/crypto_panic"
	"github.com/crypto-pulse/news/internal/integration/kafka/producer"
	"github.com/crypto-pulse/news/internal/integration/redis"
	"github.com/crypto-pulse/news/internal/route"
	"github.com/crypto-pulse/sdk"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

// TODO need to check all required env variables
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, gCtx := errgroup.WithContext(ctx)

	cryptoPanicApi := crypto_panic.NewClient()

	rdb, err := redis.NewClient()
	if err != nil {
		panic(err)
	}

	brokers := strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	topic := os.Getenv("KAFKA_TOPIC")
	maxRetries, err := strconv.Atoi(os.Getenv("KAFKA_MAX_RETRIES"))
	if err != nil {
		panic(err)
	}

	newProducer, err := producer.NewProducer(brokers, topic, maxRetries)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	route.RegisterRoutes(router, cryptoPanicApi, rdb, newProducer)

	srv := sdk.NewServer(ctx, "8082", router)

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
		<-exit
		cancel()
	}()

	g.Go(func() error {
		return srv.Run()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return srv.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
}
