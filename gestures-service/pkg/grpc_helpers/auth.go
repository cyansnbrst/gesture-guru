package grpchelpers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RequireAdmin(ctx context.Context) error {
	isAdmin, ok := ctx.Value("is_admin").(bool)
	if !ok || !isAdmin {
		return status.Error(codes.PermissionDenied, "admin access required")
	}
	return nil
}
