package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cyansnbrst/gesture-guru/gestures-service/internal/gestures"
	"github.com/cyansnbrst/gesture-guru/gestures-service/internal/models"
	"github.com/cyansnbrst/gesture-guru/gestures-service/pkg/db"
	grpchelpers "github.com/cyansnbrst/gesture-guru/gestures-service/pkg/grpc_helpers"
	gesturesv1 "github.com/cyansnbrst/gesture-guru/protos/gen/go/gestures"
)

type gesturesServer struct {
	gesturesv1.UnimplementedGesturesServiceServer
	gesturesUC gestures.UseCase
}

func NewGesturesServer(gRPCServer *grpc.Server, gesturesUC gestures.UseCase) {
	gesturesv1.RegisterGesturesServiceServer(gRPCServer, &gesturesServer{gesturesUC: gesturesUC})
}

// GetByID is the method to get a gesture by ID
func (s *gesturesServer) GetByID(ctx context.Context, in *gesturesv1.GetGestureByIDRequest) (*gesturesv1.GetGestureByIDResponse, error) {
	if in.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	gesture, err := s.gesturesUC.GetGesture(ctx, in.GetId())
	if err != nil {
		if errors.Is(err, db.ErrGestureNotFound) {
			return nil, status.Error(codes.InvalidArgument, "invalid gesture id")
		}
		return nil, status.Error(codes.Internal, "failed to get a gesture")
	}

	return &gesturesv1.GetGestureByIDResponse{Gesture: grpchelpers.ToProtoGesture(*gesture)}, nil
}

// GetAll returns all gestures
func (s *gesturesServer) GetAll(ctx context.Context, req *gesturesv1.GetAllGesturesRequest) (*gesturesv1.GetAllGesturesResponse, error) {
	gesturesList, err := s.gesturesUC.ListGestures(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to fetch gestures")
	}

	var protoGestures []*gesturesv1.Gesture
	for _, g := range gesturesList {
		protoGestures = append(protoGestures, grpchelpers.ToProtoGesture(g))
	}

	return &gesturesv1.GetAllGesturesResponse{Gestures: protoGestures}, nil
}

// Create creates a new gesture
func (s *gesturesServer) Create(ctx context.Context, req *gesturesv1.CreateGestureRequest) (*gesturesv1.CreateGestureResponse, error) {
	if err := grpchelpers.RequireAdmin(ctx); err != nil {
		return nil, err
	}

	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	if req.Description == "" {
		return nil, status.Error(codes.InvalidArgument, "description is required")
	}
	if req.VideoUrl == "" {
		return nil, status.Error(codes.InvalidArgument, "video URL is required")
	}

	id, err := s.gesturesUC.CreateGesture(ctx, &models.Gesture{
		Name:        req.Name,
		Description: req.Description,
		VideoURL:    req.VideoUrl,
		CategoryID:  req.CategoryId,
	})
	if err != nil {
		if errors.Is(err, db.ErrInvalidCategory) {
			return nil, status.Error(codes.InvalidArgument, "invalid category id")
		}
		return nil, status.Error(codes.Internal, "failed to create gesture")
	}

	return &gesturesv1.CreateGestureResponse{Id: id}, nil
}

// Update updates an existing gesture
func (s *gesturesServer) Update(ctx context.Context, req *gesturesv1.UpdateGestureRequest) (*gesturesv1.UpdateGestureResponse, error) {
	if err := grpchelpers.RequireAdmin(ctx); err != nil {
		return nil, err
	}

	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	gesture, err := s.gesturesUC.GetGesture(ctx, req.Id)
	if err != nil {
		if errors.Is(err, db.ErrGestureNotFound) {
			return nil, status.Error(codes.InvalidArgument, "gesture not found")
		}
		return nil, status.Error(codes.Internal, "failed to fetch gesture for update")
	}

	if req.Name != "" {
		gesture.Name = req.Name
	}
	if req.Description != "" {
		gesture.Description = req.Description
	}
	if req.VideoUrl != "" {
		gesture.VideoURL = req.VideoUrl
	}
	if req.CategoryId != 0 {
		gesture.CategoryID = req.CategoryId
	}

	err = s.gesturesUC.UpdateGesture(ctx, gesture)
	if err != nil {
		if errors.Is(err, db.ErrInvalidCategory) {
			return nil, status.Error(codes.InvalidArgument, "invalid category id")
		}
		return nil, status.Error(codes.Internal, "failed to update gesture")
	}

	return &gesturesv1.UpdateGestureResponse{}, nil
}

// Delete deletes a gesture by ID
func (s *gesturesServer) Delete(ctx context.Context, req *gesturesv1.DeleteGestureRequest) (*gesturesv1.DeleteGestureResponse, error) {
	if err := grpchelpers.RequireAdmin(ctx); err != nil {
		return nil, err
	}

	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	err := s.gesturesUC.DeleteGesture(ctx, req.Id)
	if err != nil {
		if errors.Is(err, db.ErrGestureNotFound) {
			return nil, status.Error(codes.InvalidArgument, "gesture not found")
		}
		return nil, status.Error(codes.Internal, "failed to delete gesture")
	}

	return &gesturesv1.DeleteGestureResponse{}, nil
}
