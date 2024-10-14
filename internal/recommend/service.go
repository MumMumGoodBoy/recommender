package recommend

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/mummumgoodboy/recommender/internal/model"
	"github.com/mummumgoodboy/recommender/internal/proto"
	"github.com/zhenghaoz/gorse/client"
)

var _ proto.RecommendServiceServer = &RecommendService{}

func NewRecommendService(gorse *client.GorseClient) *RecommendService {
	return &RecommendService{gorse: gorse}
}

type RecommendService struct {
	proto.UnimplementedRecommendServiceServer

	gorse *client.GorseClient
}

// AddEvent implements proto.RecommendServiceServer.
func (r *RecommendService) AddEvent(ctx context.Context, req *proto.AddEventReq) (*proto.Empty, error) {
	eventType := convertEventTypes(req.EventType)
	if eventType == "" {
		slog.Warn(
			"invalid event type when adding event",
			"event_type", req.EventType,
			"user", req.UserId,
		)
		return nil, errors.New("invalid event type")
	}

	userID := strconv.Itoa(int(req.UserId))
	if req.UserId == 0 {
		userID = "default"
	}
	_, err := r.gorse.PutFeedback(ctx, []client.Feedback{
		{
			FeedbackType: eventType,
			UserId:       userID,
			ItemId:       req.ItemId,
			Timestamp:    time.Now().Format(time.RFC3339),
		},
	})
	if err != nil {
		slog.Warn(
			"failed to insert feedback",
			"error", err,
			"event_type", req.EventType,
			"user", req.UserId,
		)
		return nil, err
	}

	return &proto.Empty{}, nil
}

func (r *RecommendService) RemoveEvent(ctx context.Context, req *proto.RemoveEventReq) (*proto.Empty, error) {
	eventType := convertEventTypes(req.EventType)
	if eventType == "" {
		slog.Warn(
			"invalid event type when removing event",
			"event_type", req.EventType,
			"user", req.UserId,
		)
		return nil, errors.New("invalid event type")
	}

	_, err := r.gorse.DelFeedback(ctx,
		eventType,
		strconv.Itoa(int(req.UserId)),
		req.ItemId,
	)
	if err != nil {
		slog.Warn(
			"failed to delete feedback",
			"error", err,
			"event_type", req.EventType,
			"user", req.UserId,
		)
		return nil, err
	}

	return &proto.Empty{}, nil
}

// GetFoodRecommendations implements proto.RecommendServiceServer.
func (r *RecommendService) GetFoodRecommendations(ctx context.Context, req *proto.GetRecommendationsRequest) (*proto.GetRecommendationsResponse, error) {
	if req.Limit > 20 {
		req.Limit = 20
	}

	userId := strconv.Itoa(int(req.UserId))
	if req.UserId == 0 {
		userId = "default"
	}

	ids, err := r.gorse.GetItemRecommend(ctx, userId, nil,
		string(model.FeedbackRead),
		"0h5m0s",
		int(req.Limit),
		int(req.Offset))
	if err != nil {
		slog.Warn(
			"failed to get item recommendations",
			"error", err,
			"user", req.UserId,
		)
		return nil, fmt.Errorf("failed to get item recommendations: %w", err)
	}

	if len(ids) != 0 {
		return &proto.GetRecommendationsResponse{ItemIds: ids}, nil
	}

	// Opt-out to popular items if no recommendations found
	slog.Warn(
		"no recommendations found",
		"user", req.UserId,
	)

	scores, err := r.gorse.GetItemPopular(ctx, "", int(req.Limit), int(req.Offset))
	if err != nil {
		slog.Warn(
			"failed to get popular items",
			"error", err,
			"user", req.UserId,
		)
		return nil, fmt.Errorf("failed to get popular items: %w", err)
	}

	for _, score := range scores {
		ids = append(ids, score.Id)
	}

	return &proto.GetRecommendationsResponse{ItemIds: ids}, nil
}
