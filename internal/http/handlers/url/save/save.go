package save

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/urusofam/urlShortener/internal/http/api/response"
	"github.com/urusofam/urlShortener/internal/log/sl"
	"log/slog"
	"net/http"
)

type Request struct {
	URL   string `json:"url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	response.Response
	Alias string `json:"alias,omitempty"`
}

type UrlSaver interface {
	SaveURL(urlToSave, alias string) error
}

func New(log *slog.Logger, urlSaver UrlSaver) http.Handler {
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
	}
}
