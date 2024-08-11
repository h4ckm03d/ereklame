package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/h4ckm03d/ereklame/internal/database/sqlc"
	"github.com/h4ckm03d/ereklame/internal/resource"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/oauth"
	"github.com/go-chi/render"
)

func (s *Server) RegisterRoutes() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.CleanPath)
	r.Use(middleware.Recoverer)
	r.Use(JSONLogger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	queries := sqlc.New(s.db.GetConnection())

	// API version 1.

	bearerServer := oauth.NewBearerServer(
		os.Getenv("OAUTH_SECRET"),
		time.Second*120,
		&TestUserVerifier{},
		nil)
	r.Post("/token", bearerServer.UserCredentials)
	r.Post("/auth", bearerServer.ClientCredentials)

	r.Route("/v1", func(r chi.Router) {
		r.Use(apiVersionCtx("v1"))
		r.Get("/", s.HelloWorldHandler)
		r.Get("/health", s.healthHandler)
		r.Mount("/users", resource.NewUsers(queries).Routes())
	})

	return r
}

func apiVersionCtx(version string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			type contextKey string

			const apiVersionKey contextKey = "api.version"

			r = r.WithContext(context.WithValue(r.Context(), apiVersionKey, version))
			next.ServeHTTP(w, r)
		})
	}
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]any)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
