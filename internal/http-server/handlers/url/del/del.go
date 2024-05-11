package del

import (
	"errors"
	"log/slog"
	"net/http"
	resp "url-shortener/internal/lib/api/response"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Request struct {
	Alias string `json:"alias" validate:"required"`
}

type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

type AliasDeleter interface {
	DeleteAlias(aliasToDel string) error
}

func New(log *slog.Logger, aliasDeleter AliasDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.del.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		// var req Request
		// err := render.DecodeJSON(r.Body, &req)
		// if err != nil {
		// 	log.Error("failed to decode request body", sl.Err(err))
		// 	render.JSON(w, r, resp.Error("failed to decod request"))
		// 	return
		// }
		// log.Info("request body decoded", slog.Any("request", req))

		// if err := validator.New().Struct(req); err != nil {
		// 	validateErr := err.(validator.ValidationErrors)

		// 	log.Error("invalid request", sl.Err(err))

		// 	render.JSON(w, r, resp.ValidationError(validateErr))
		// 	return
		// }

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")

			render.JSON(w, r, resp.Error("invalid request"))

			return
		}

		err := aliasDeleter.DeleteAlias(alias)
		if errors.Is(err, storage.ErrAliasNotFound) {
			log.Info("alias not exists", slog.String("alias", alias))
			render.JSON(w, r, resp.Error("alias not exists"))
			return
		}
		if err != nil {
			log.Error("failed to delete alias", sl.Err(err))
			render.JSON(w, r, resp.Error("failed to delete alias"))
			return
		}
		log.Info("alias deleted", slog.String("alias", alias))
		responseOK(w, r, alias)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request, alias string) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
		Alias:    alias,
	})
}
