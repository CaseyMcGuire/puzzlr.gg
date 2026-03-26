package reqctx

import (
	"context"
	"errors"
)

type contextKey int

const userIDKey contextKey = iota

var ErrNoUserIDInContext = errors.New("no user ID in context")

func WithUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func UserIDFromContext(ctx context.Context) (int, error) {
	userID, ok := ctx.Value(userIDKey).(int)
	if !ok {
		return 0, ErrNoUserIDInContext
	}
	return userID, nil
}
