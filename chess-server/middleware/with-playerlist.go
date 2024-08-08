package middleware

import (
	"context"
	"kzinthant-d3v/go-chess-server/record"
	"net/http"
)

const PlayerListKey contextKey = "playerList"

func WithPlayerList(pr *record.PlayerGameList) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), PlayerListKey, pr)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
