package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rsdoiel/opml"

	"github.com/rphillips/escapepod/pkg/app"
	"github.com/rphillips/escapepod/pkg/db"
	"github.com/rphillips/escapepod/pkg/models"
)

func PodcastsList(c echo.Context) error {
	app := c.Get("app").(*app.App)
	podcasts := []models.Podcast{}
	if err := app.DB.Find(&podcasts).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, podcasts)
}

func PodcastsGet(c echo.Context) error {
	app := c.Get("app").(*app.App)
	id := c.Param("id")
	podcast := models.Podcast{}
	if err := app.DB.Where(id).First(&podcast).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, podcast)
}

func PodcastsGetEpisodes(c echo.Context) error {
	app := c.Get("app").(*app.App)
	id := c.Param("id")
	podcast := models.Podcast{}
	if err := app.DB.Where(id).First(&podcast).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := app.DB.Model(&podcast).Association("Episodes").Find(&podcast.Episodes).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, podcast.Episodes)
}

func PodcastsCreate(c echo.Context) error {
	app := c.Get("app").(*app.App)
	podcast := &models.Podcast{}
	if err := c.Bind(podcast); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := podcast.Crawl(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := app.DB.Create(podcast).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, podcast)
}

func PodcastsImport(c echo.Context) error {
	app := c.Get("app").(*app.App)
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files, ok := form.File["file"]
	if !ok {
		return c.JSON(http.StatusBadRequest, "File not found in upload")
	}
	src, err := files[0].Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	defer src.Close()
	data, err := ioutil.ReadAll(src)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	parsed, err := opml.Parse(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	go importFromOPML(app, parsed)
	return c.JSON(http.StatusOK, "OK")
}

func importFromOPML(app *app.App, o *opml.OPML) error {
	if len(o.Body.Outline) <= 0 {
		return fmt.Errorf("invalid opml")
	}
	if o.Body.Outline[0].Text != "feeds" {
		return fmt.Errorf("opml does not contain a feeds attribute")
	}
	for _, feed := range o.Body.Outline[0].Outline {
		go func(feed opml.Outline) {
			app.R.Logger.Infof("Importing feed: %v", feed.Title)
			podcast := &models.Podcast{
				FeedURL: db.NewString(feed.XMLURL),
			}
			if err := podcast.Crawl(); err != nil {
				app.R.Logger.Errorf("error crawling feed: %v", err)
				return
			}
			if err := app.DB.Create(podcast).Error; err != nil {
				app.R.Logger.Errorf("error create podcast: %v", err)
				return
			}
		}(feed)
	}
	return nil
}

func PodcastsImageGet(c echo.Context) error {
	app := c.Get("app").(*app.App)
	id := c.Param("id")
	podcast := models.Podcast{}
	if err := app.DB.Where(id).First(&podcast).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.Blob(http.StatusOK, "image/jpeg", podcast.ImageResized200)
}
