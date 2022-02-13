package godatatools

import (
	"context"
	"net/http"
)

// VerifyToken is a middleware that retrieves the Authorization header,
// and checks that the token is valid by querying GoData.
func (s Server) VerifyToken() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodOptions {
				f(w, r)
				return
			}
			token := r.Header.Get("Authorization")
			userID, err := s.DbRepository.FindUserIDForAccessToken(r.Context(), token)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), "userID", userID)
			f(w, r.WithContext(ctx))
		}
	}
}
