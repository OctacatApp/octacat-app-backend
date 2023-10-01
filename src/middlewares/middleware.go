package middlewares

import (
	"context"
	"net/http"

	"github.com/irdaislakhuafa/octacat-app-backend/src/business/usecase"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
)

func GraphQLMiddleware(cfg *config.AppConfig, uc *usecase.Usecase) func(next http.Handler) http.Handler {
	handler := func(next http.Handler) http.Handler {
		handler := func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), "headers", r.Header))
			r = r.WithContext(context.WithValue(r.Context(), "cfg", cfg))
			r = r.WithContext(context.WithValue(r.Context(), "uc", uc))
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(handler)
	}
	return handler
}
