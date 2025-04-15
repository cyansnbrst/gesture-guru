package main

import (
	"context"
	"log"
	"net/http"

	gesturesv1 "github.com/cyansnbrst/gesture-guru/protos/gen/go/gestures"
	ssov1 "github.com/cyansnbrst/gesture-guru/protos/gen/go/sso"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := gesturesv1.RegisterGesturesHandlerFromEndpoint(ctx, mux, "gestures-service:3001", opts)
	if err != nil {
		log.Fatalf("failed to register Gestures: %v", err)
	}

	err = ssov1.RegisterAuthHandlerFromEndpoint(ctx, mux, "sso-service:3000", opts)
	if err != nil {
		log.Fatalf("failed to register Auth: %v", err)
	}

	log.Println("API Gateway started on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("gateway failed: %v", err)
	}
}
