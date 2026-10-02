package api

import (
	_ "embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

// letsgobold additions live in this one file so upstream merges stay trivial.
// The panel's Apps card and the browser tab pick the favicon up from here.

//go:embed lgb_favicon.svg
var lgbFavicon []byte

func (s *Server) registerLGBRoutes() {
	favicon := func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=86400")
		c.Data(http.StatusOK, "image/svg+xml", lgbFavicon)
	}
	s.engine.GET("/favicon.svg", favicon)
	s.engine.GET("/favicon.ico", favicon)
}
