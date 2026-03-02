package reqctx

import (
	"context"
	"fmt"
)

type contextKey int

const userIDKey contextKey = iota

func WithUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func UserIDFromContext(ctx context.Context) (int, error) {
	userID, ok := ctx.Value(userIDKey).(int)
	if !ok {
		return 0, fmt.Errorf("no user ID in context")
	}
	return userID, nil
}
