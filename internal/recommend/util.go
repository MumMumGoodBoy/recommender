package recommend

import "github.com/mummumgoodboy/recommender/internal/proto"

func convertEventTypes(eventType proto.EventType) string {
	switch eventType {
	case proto.EventType_VIEW:
		return "view"
	case proto.EventType_FAVORITE:
		return "favorite"
	case proto.EventType_RATING:
		return "rating"
	default:
		return ""
	}
}
