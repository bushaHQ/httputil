package render

import (
	"github.com/bushaHQ/httputil/errors"
	"net/http"

	"github.com/go-chi/render"
)

// Render renders res for display
func Render(w http.ResponseWriter, r *http.Request, res interface{}) {
	switch res.(type) {
	case render.Renderer:
		render.Render(w, r, res.(render.Renderer))
	case error:
		render.Render(w, r, errors.New(res.(error).Error(), http.StatusBadRequest))
	default:
		render.JSON(w, r, res)
	}
}
