package middlewares

import (
	"context"
	"net/http"

	"github.com/irdaislakhuafa/octacat-app-backend/src/business/usecase"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/key"
)

func GraphQLMiddleware(cfg *config.AppConfig, uc *usecase.Usecase) func(next http.Handler) http.Handler {
	handler := func(next http.Handler) http.Handler {
		handler := func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), key.Key("headers"), r.Header.Clone()))
			r = r.WithContext(context.WithValue(r.Context(), key.Key("cfg"), cfg))
			r = r.WithContext(context.WithValue(r.Context(), key.Key("uc"), uc))
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(handler)
	}
	return handler
}
