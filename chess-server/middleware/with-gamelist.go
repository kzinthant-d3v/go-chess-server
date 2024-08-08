package middleware

import (
	"context"
	"kzinthant-d3v/go-chess-server/record"
	"net/http"
)

type contextKey string

const GameListKey contextKey = "gameList"

func WithGameList(gl *record.GameList) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), GameListKey, gl)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
