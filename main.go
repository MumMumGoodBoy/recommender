package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/mummumgoodboy/recommender/internal/event"
	"github.com/mummumgoodboy/recommender/internal/recommend"
	"github.com/mummumgoodboy/recommender/proto"
	"github.com/wagslane/go-rabbitmq"
	"github.com/zhenghaoz/gorse/client"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file", err)
	}

	gorseServer := os.Getenv("GORSE_SERVER")
	gorseAPIKey := os.Getenv("GORSE_API_KEY")
	if gorseServer == "" || gorseAPIKey == "" {
		log.Panic("Please set GORSE_SERVER and GORSE_API_KEY in .env file")
	}
	gorse := client.NewGorseClient(gorseServer, gorseAPIKey)

	rabbitmqURL := os.Getenv("RABBITMQ_URL")
	rabbitmqConn, err := rabbitmq.NewConn(rabbitmqURL)
	if err != nil {
		log.Panic("failed to connect to rabbitmq", err)
	}

	recommendService := recommend.NewRecommendService(gorse)

	eventService := event.EventService{
		Conn:             rabbitmqConn,
		RecommendService: recommendService,
	}

	close, err := eventService.ListenToEvents()
	if err != nil {
		log.Panic("failed to listen to events", err)
	}
	defer func() {
		err := close()
		if err != nil {
			slog.Warn("failed to close event service",
				"error", err,
			)
		}
	}()

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Panic("PORT must be an integer", err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	proto.RegisterRecommendServiceServer(server, recommendService)

	log.Printf("Recommend service is running on port %d", port)
	log.Fatal(server.Serve(lis))
}
