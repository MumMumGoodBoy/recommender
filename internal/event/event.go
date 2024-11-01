package event

import (
	"context"
	"encoding/json"
	"log"
	"log/slog"

	"github.com/mummumgoodboy/recommender/internal/model"
	"github.com/mummumgoodboy/recommender/internal/recommend"
	"github.com/mummumgoodboy/recommender/pkg/closer"
	"github.com/mummumgoodboy/recommender/proto"
	"github.com/wagslane/go-rabbitmq"
)

type EventService struct {
	Conn             *rabbitmq.Conn
	RecommendService *recommend.RecommendService
}

func (e *EventService) ListenToEvents() (close closer.CloseFunc, err error) {
	closes := closer.MultipleCloser{}
	{ // Review Event
		reviewConsumer, err := rabbitmq.NewConsumer(e.Conn,
			"review_recommend_queue",
			rabbitmq.WithConsumerOptionsExchangeName("review_topic"),
			rabbitmq.WithConsumerOptionsBinding(rabbitmq.Binding{
				RoutingKey: "review.*",
				BindingOptions: rabbitmq.BindingOptions{
					Declare: true,
				},
			}),
			rabbitmq.WithConsumerOptionsConcurrency(5),
		)
		if err != nil {
			slog.Error("failed to create review consumer", "error", err)
			return closes.Close, err
		}

		closes.AppendNoErr(reviewConsumer.Close)

		go func() {
			err := reviewConsumer.Run(e.handleReviewEvent)
			if err != nil {
				log.Panic("failed to run review consumer", err)
			}
		}()
	}

	{ // Favorite Event
		favoriteConsumer, err := rabbitmq.NewConsumer(e.Conn,
			"favorite_recommend_queue",
			rabbitmq.WithConsumerOptionsExchangeName("favorite_topic"),
			rabbitmq.WithConsumerOptionsBinding(rabbitmq.Binding{
				RoutingKey: "favorite.*",
				BindingOptions: rabbitmq.BindingOptions{
					Declare: true,
				},
			}),
			rabbitmq.WithConsumerOptionsConcurrency(5),
		)
		if err != nil {
			slog.Error("failed to create review consumer", "error", err)
			return closes.Close, err
		}

		closes.AppendNoErr(favoriteConsumer.Close)

		go func() {
			err := favoriteConsumer.Run(e.handleFavoriteEvent)
			if err != nil {
				log.Panic("failed to run review consumer", err)
			}
		}()
	}

	return closes.Close, nil
}

func (e *EventService) handleReviewEvent(d rabbitmq.Delivery) rabbitmq.Action {
	ctx := context.Background()

	var data model.ReviewEventDTO
	err := json.Unmarshal(d.Body, &data)
	if err != nil {
		slog.Error("failed to unmarshal review event", "error", err)
		return rabbitmq.NackDiscard
	}

	switch data.EventType {
	case "review.update":
		fallthrough
	case "review.create":
		if data.Rating < 3.0 {
			if data.EventType == "review.create" {
				return rabbitmq.Ack
			}
			_, err := e.RecommendService.RemoveEvent(ctx, &proto.RemoveEventReq{
				EventType: proto.EventType_RATING,
				UserId:    int64(data.ReviewerId),
				ItemId:    data.FoodId,
			})
			if err != nil {
				slog.Error("failed to remove review event", "error", err)
				return rabbitmq.NackRequeue
			}
			return rabbitmq.Ack
		}
		_, err := e.RecommendService.AddEvent(ctx, &proto.AddEventReq{
			EventType: proto.EventType_RATING,
			UserId:    int64(data.ReviewerId),
			ItemId:    data.FoodId,
		})
		if err != nil {
			slog.Error("failed to add review event", "error", err)
			return rabbitmq.NackRequeue
		}

		return rabbitmq.Ack
	case "review.delete":
		_, err := e.RecommendService.RemoveEvent(ctx, &proto.RemoveEventReq{
			EventType: proto.EventType_RATING,
			UserId:    int64(data.ReviewerId),
			ItemId:    data.FoodId,
		})
		if err != nil {
			slog.Error("failed to remove review event", "error", err)
			return rabbitmq.NackRequeue
		}
		return rabbitmq.Ack
	default:
		slog.Warn("unknown event type",
			"event_type", data.EventType,
			"reviewer_id", data.ReviewerId,
		)
		return rabbitmq.Ack
	}
}

func (e *EventService) handleFavoriteEvent(d rabbitmq.Delivery) rabbitmq.Action {
	ctx := context.Background()

	var data model.FavoriteEventDTO
	err := json.Unmarshal(d.Body, &data)
	if err != nil {
		slog.Error("failed to unmarshal review event", "error", err)
		return rabbitmq.NackDiscard
	}

	switch data.EventType {
	case "favorite.add":
		_, err := e.RecommendService.AddEvent(ctx, &proto.AddEventReq{
			EventType: proto.EventType_FAVORITE,
			UserId:    int64(data.UserId),
			ItemId:    data.FoodId,
		})
		if err != nil {
			slog.Error("failed to add favorite event", "error", err)
			return rabbitmq.NackRequeue
		}

		return rabbitmq.Ack
	case "favorite.remove":
		_, err := e.RecommendService.RemoveEvent(ctx, &proto.RemoveEventReq{
			EventType: proto.EventType_FAVORITE,
			UserId:    int64(data.UserId),
			ItemId:    data.FoodId,
		})
		if err != nil {
			slog.Error("failed to remove favorite event", "error", err)
			return rabbitmq.NackRequeue
		}
		return rabbitmq.Ack
	default:
		slog.Warn("unknown event type",
			"event_type", data.EventType,
			"reviewer_id", data.UserId,
		)
		return rabbitmq.Ack
	}
}
