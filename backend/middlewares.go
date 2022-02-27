package godatatools

import (
	"context"
	"net/http"
	"strings"
)

type AuthToken struct {
	UserID string
}

// VerifyToken is a middleware that retrieves the Authorization header,
// and checks that the token is valid by querying GoData.
func (s Server) VerifyToken() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodOptions {
				f(w, r)
				return
			}
			token := strings.Split(r.Header.Get("Authorization"), " ")
			if len(token) != 2 {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			userID, err := s.DbRepository.FindUserIDForAccessToken(r.Context(), token[1])

			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), "authToken", AuthToken{UserID: userID}) //nolint
			f(w, r.WithContext(ctx))
		}
	}
}
