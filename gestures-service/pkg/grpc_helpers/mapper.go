package grpchelpers

import (
	"github.com/cyansnbrst/gesture-guru/gestures-service/internal/models"
	gesturesv1 "github.com/cyansnbrst/gesture-guru/protos/gen/go/gestures"
)

func ToProtoGesture(g models.Gesture) *gesturesv1.Gesture {
	return &gesturesv1.Gesture{
		Id:          g.ID,
		Name:        g.Name,
		Description: g.Description,
		VideoUrl:    g.VideoURL,
		CategoryId:  g.CategoryID,
		CreatedAt:   g.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
