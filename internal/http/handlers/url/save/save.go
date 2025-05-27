package save

import (
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/urusofam/urlShortener/internal/http/api/response"
	"github.com/urusofam/urlShortener/internal/log/sl"
	"github.com/urusofam/urlShortener/internal/random"
	"github.com/urusofam/urlShortener/internal/storage"
	"log/slog"
	"net/http"
)

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	response.Response
	Alias string `json:"alias,omitempty"`
}

type UrlSaver interface {
	SaveURL(urlToSave, alias string) error
}

func New(log *slog.Logger, urlSaver UrlSaver, aliasLength int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.saveURL"

		log = log.With(slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to parse request body", sl.Err(err))

			render.JSON(w, r, response.Error("failed to parse request"))

			return
		}

		log.Info("request body parsed", slog.Any("req", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("failed to validate request", sl.Err(err))

			render.JSON(w, r, response.ValidationError(validateErr))

			return
		}

		alias := req.Alias
		if alias == "" {
			alias = random.NewRandomString(aliasLength)
		}

		err = urlSaver.SaveURL(req.URL, alias)
		if errors.Is(err, storage.ErrURLExists) {
			log.Info("url already exists", sl.Err(err))

			render.JSON(w, r, response.Error("failed to add url"))

			return
		}
		if err != nil {
			log.Error("failed to save url", sl.Err(err))

			render.JSON(w, r, response.Error("failed to save url"))

			return
		}

		log.Info("url saved")

		render.JSON(w, r, Response{
			Response: response.OK(),
			Alias:    alias,
		})
	}
}
