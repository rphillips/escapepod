package app

import (
	"github.com/jinzhu/gorm"
	"github.com/knadh/stuffbin"
	"github.com/labstack/echo/v4"

	"github.com/rphillips/escapepod/pkg/crawler"
)

type Constants struct {
	Address string // listen address
}

type App struct {
	R       *echo.Echo
	DB      *gorm.DB
	FS      stuffbin.FileSystem
	Crawler *crawler.Crawler
	Constants
}

func New(r *echo.Echo, db *gorm.DB, fs stuffbin.FileSystem, constants Constants) *App {
	return &App{
		R:         r,
		DB:        db,
		FS:        fs,
		Crawler:   crawler.NewCrawler(db),
		Constants: constants,
	}
}
