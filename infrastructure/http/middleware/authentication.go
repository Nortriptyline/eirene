package middleware

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

// responseRecorder is a wrapper around gin.ResponseWriter that captures the status code
type responseRecorder struct {
	gin.ResponseWriter
	statusCode int
}

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Scope string `json:"scope"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid issuer URL"})
			c.Abort()
			return
		}

		provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

		jwtValidator, err := validator.New(
			provider.KeyFunc,
			validator.RS256,
			issuerURL.String(),
			[]string{os.Getenv("AUTH0_AUDIENCE")},
			validator.WithCustomClaims(
				func() validator.CustomClaims {
					return &CustomClaims{}
				},
			),
			validator.WithAllowedClockSkew(time.Minute),
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create JWT validator"})
			c.Abort()
			return
		}

		errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
			log.Printf("Encountered error while validating JWT: %v", err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message":"Failed to validate JWT."}`))
			c.Abort()
		}

		m := jwtmiddleware.New(
			jwtValidator.ValidateToken,
			jwtmiddleware.WithErrorHandler(errorHandler),
		)

		// Adapt the middleware to work with Gin
		handler := func(w http.ResponseWriter, r *http.Request) {
			m.CheckJWT(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// This handler will be called if the JWT passes validation
				c.Next()
			})).ServeHTTP(w, r)
		}

		// Create a ResponseRecorder to capture the response
		recorder := &responseRecorder{ResponseWriter: c.Writer, statusCode: http.StatusOK}
		c.Writer = recorder

		// Call the handler
		handler(c.Writer, c.Request)

		// Check if the middleware aborted the request
		if recorder.statusCode == http.StatusUnauthorized {
			c.Abort()
			return
		}

		c.Next()
	}
}
