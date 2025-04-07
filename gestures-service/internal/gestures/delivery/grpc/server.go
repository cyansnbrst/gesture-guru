package grpc

import (
	"context"

	"github.com/cyansnbrst/gesture-guru/gestures-service/internal/gestures"
	"github.com/cyansnbrst/gesture-guru/gestures-service/internal/models"
)

type gesturesServer struct {
	gesturesv1
	repo gestures.Repository
}

func NewGesturesServiceServer(repo gestures.Repository) *GesturesServiceServer {
	return &GesturesServiceServer{repo: repo}
}

func (s *GesturesServiceServer) GetByID(ctx context.Context, req *pb.GetGestureByIDRequest) (*pb.GetGestureByIDResponse, error) {
	gesture, err := s.repo.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err // ты можешь оборачивать ошибки через кастомные gRPC codes
	}

	return &pb.GetGestureByIDResponse{
		Gesture: toProtoGesture(gesture),
	}, nil
}

func (s *GesturesServiceServer) GetAll(ctx context.Context, req *pb.GetAllGesturesRequest) (*pb.GetAllGesturesResponse, error) {
	gesturesList, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var protoGestures []*pb.Gesture
	for _, g := range gesturesList {
		protoGestures = append(protoGestures, toProtoGesture(&g))
	}

	return &pb.GetAllGesturesResponse{
		Gestures: protoGestures,
	}, nil
}

func (s *GesturesServiceServer) Create(ctx context.Context, req *pb.CreateGestureRequest) (*pb.CreateGestureResponse, error) {
	id, err := s.repo.Create(ctx, &models.Gesture{
		Name:        req.Name,
		Description: req.Description,
		VideoURL:    req.VideoUrl,
		CategoryID:  req.CategoryId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateGestureResponse{Id: id}, nil
}

func (s *GesturesServiceServer) Update(ctx context.Context, req *pb.UpdateGestureRequest) (*pb.UpdateGestureResponse, error) {
	err := s.repo.Update(ctx, &models.Gesture{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		VideoURL:    req.VideoUrl,
		CategoryID:  req.CategoryId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateGestureResponse{}, nil
}

func (s *GesturesServiceServer) Delete(ctx context.Context, req *pb.DeleteGestureRequest) (*pb.DeleteGestureResponse, error) {
	err := s.repo.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteGestureResponse{}, nil
}
