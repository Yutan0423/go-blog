package middlewares

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/Yutan0423/go-medium-level/apperrors"
	"google.golang.org/api/idtoken"
)

var (
	googleClientID = os.Getenv("GOOGLE_CLIENT_ID")
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorization := req.Header.Get("Authorization")

		authHeaders := strings.Split(authorization, " ")
		if len(authHeaders) != 2 {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		bearer, idToken := authHeaders[0], authHeaders[1]
		if bearer != "Bearer" || idToken == "" {
			err := apperrors.Unauthorized.Wrap(errors.New("incvalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		tokenValidator, err := idtoken.NewValidator(context.Background())
		if bearer != "Bearer" || idToken == "" {
			err := apperrors.CannotMakeValidator.Wrap(errors.New("incvalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

	})
}
