package routes

import (
	"net/http"

	"github.com/knadh/stuffbin"
	"github.com/labstack/echo/v4"

	"github.com/rphillips/escapepod/pkg/app"
	"github.com/rphillips/escapepod/pkg/handlers"
)

func RegisterRoutes(r *echo.Echo, fs stuffbin.FileSystem) {
	r.GET("/", handleIndexPage)
	r.GET("/*", echo.WrapHandler(fs.FileServer()))

	api := r.Group("/api/v1")
	{
		api.POST("/podcasts", handlers.PodcastsCreate)
		api.POST("/podcasts/import", handlers.PodcastsImport)
		api.GET("/podcasts", handlers.PodcastsList)
		api.GET("/podcasts/:id", handlers.PodcastsGet)
		api.GET("/podcasts/:id/episodes", handlers.PodcastsGetEpisodes)
		api.GET("/podcasts/:id/image", handlers.PodcastsImageGet)
	}
}

func handleIndexPage(c echo.Context) error {
	app := c.Get("app").(*app.App)
	b, err := app.FS.Read("/index.html")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set("Content-Type", "text/html")
	return c.String(http.StatusOK, string(b))
}
