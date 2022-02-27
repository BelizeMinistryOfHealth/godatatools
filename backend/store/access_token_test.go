package store

import (
	"context"
	"testing"
)

func TestStore_FindUserIDForAccessToken(t *testing.T) {
	ctx := context.Background()
	store := setupDb(t, ctx)
	defer store.Disconnect(ctx) //nolint:govet,errcheck
	token := "95unAdVOyQkwDBeNjohXyV2WXVReUAD4PF3mYDzPRSZYGTjYhjbxHmuYFwnl8KuP"

	userID, err := store.FindUserIDForAccessToken(ctx, token)
	if err != nil {
		t.Fatalf("FindUserIDForAccessToken failed: %v", err)
	}

	t.Logf("userID: %s", userID)
}
