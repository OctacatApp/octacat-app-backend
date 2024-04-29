package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/irdaislakhuafa/go-sdk/appcontext"
	"github.com/irdaislakhuafa/go-sdk/language"
	"github.com/irdaislakhuafa/octacat-app-backend/src/business/usecase"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/key"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/operator"
)

func GraphQLMiddleware(cfg *config.AppConfig, uc *usecase.Usecase) func(next http.Handler) http.Handler {
	handler := func(next http.Handler) http.Handler {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lang := r.Header.Get(string(appcontext.AcceptLanguage))
			r = r.WithContext(context.WithValue(r.Context(), key.Key("headers"), r.Header.Clone()))
			r = r.WithContext(context.WithValue(r.Context(), key.Key("cfg"), cfg))
			r = r.WithContext(context.WithValue(r.Context(), key.Key("uc"), uc))
			r = r.WithContext(appcontext.SetAcceptLanguage(r.Context(), operator.Ternary(lang == "", language.English, language.Language(lang))))
			r = r.WithContext(appcontext.SetRequestStartTime(r.Context(), time.Now().UTC()))
			r = r.WithContext(appcontext.SetServiceVersion(r.Context(), cfg.App.Version))
			r = r.WithContext(appcontext.SetRequestID(r.Context(), r.Header.Get(string(appcontext.RequestID))))
			r = r.WithContext(appcontext.SetUserAgent(r.Context(), r.Header.Get(string(appcontext.UserAgent))))
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(handler)
	}
	return handler
}
