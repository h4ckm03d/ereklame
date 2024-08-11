package resource

import (
	"net/http"
	"os"

	"github.com/h4ckm03d/ereklame/internal/database/sqlc"

	"github.com/rs/zerolog/log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/oauth"
	"github.com/go-chi/render"
)

type userDto struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type userResource struct {
	query *sqlc.Queries
}

func NewUsers(query *sqlc.Queries) *userResource {
	return &userResource{query: query}
}

func (rs userResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(oauth.Authorize(os.Getenv("OAUTH_SECRET"), nil))
	r.Get("/", rs.ListUsers)

	return r
}

func (rs *userResource) ListUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := rs.query.GetUsers(ctx)
	if err != nil {
		log.Err(err).Msg("failed to fetch users")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userData := make([]userDto, len(users))
	for i, user := range users {
		userData[i] = userDto{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
	}

	render.JSON(w, r, userData)
}
