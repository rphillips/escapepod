package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/knadh/stuffbin"
	_ "github.com/lib/pq"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/rphillips/escapepod/pkg/app"
	"github.com/rphillips/escapepod/pkg/db"
	"github.com/rphillips/escapepod/pkg/routes"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		r := echo.New()
		r.HidePort = true
		r.HideBanner = true

		// DB Setup
		db, err := db.OpenDB()
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
		defer db.Close()

		log.Info("Connected to database...")
		if debug {
			r.Debug = true
			r.Logger.SetLevel(log.DEBUG)
			db.LogMode(true)
		}

		// Load filesystem
		fs, err := loadFilesystem(os.Args[0])
		if err != nil {
			log.Fatal(err)
		}

		// Start crawler
		app := app.New(r, db, fs, app.Constants{
			Address: viper.GetString("app.address"),
		})
		stopCh := make(chan struct{})
		go app.Crawler.Run(5, stopCh)

		// Web
		r.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				c.Set("app", app)
				return next(c)
			}
		})
		r.Use(middleware.Logger())
		r.Use(middleware.Recover())
		r.Use(middleware.RequestID())
		r.Use(middleware.CORS())
		routes.RegisterRoutes(r, fs)

		go func() {
			address := viper.GetString("app.address")
			log.Infof("Starting server on: %v", address)
			if err := r.Start(address); err != nil && err != http.ErrServerClosed {
				r.Logger.Error(err)
			}
		}()

		// gracefull shutdown
		sigChan := make(chan os.Signal)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()

		if err := r.Shutdown(ctx); err != nil {
			r.Logger.Errorf("error on shutdown: %v", err)
		}

		r.Logger.Info("Shutdown")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func loadFilesystem(path string) (stuffbin.FileSystem, error) {
	fs, err := stuffbin.UnStuff(path)
	if err == nil {
		return fs, nil
	}
	return stuffbin.NewLocalFS("/", []string{
		"frontend-vue/dist:/",
	}...)
}
