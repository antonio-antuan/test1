package posts

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/qdm12/go-template/internal/server/contenttype"
	"github.com/qdm12/go-template/internal/server/httperr"
)

func (h *handler) getPosts(w http.ResponseWriter, r *http.Request) {
	_, responseContentType, err := contenttype.APICheck(r.Header)
	w.Header().Set("Content-Type", responseContentType)
	errResponder := httperr.NewResponder(responseContentType, h.logger)

	if err != nil {
		errResponder.Respond(w, http.StatusNotAcceptable, err.Error())
		return
	}

	posts, err := h.proc.GetPosts(r.Context())
	if err != nil {
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			errResponder.Respond(w, http.StatusRequestTimeout, "")
		default:
			h.logger.Error(err.Error())
			errResponder.Respond(w, http.StatusInternalServerError, "")
		}
		return
	}

	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		h.logger.Error(err.Error())
		errResponder.Respond(w, http.StatusInternalServerError, "")
		return
	}
}

var (
	errUserIDMissingURLPath = errors.New("user ID must be provided in the URL path")
	errUserIDMalformed      = errors.New("user ID is malformed")
)

func extractUserID(r *http.Request) (id uint64, err error) {
	s := chi.URLParam(r, "id")
	if s == "" {
		return 0, errUserIDMissingURLPath
	}
	id, err = strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: %q", errUserIDMalformed, s)
	}
	return id, nil
}
