package main

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"

	"github.com/rphillips/escapepod/pkg/db"
	"github.com/rphillips/escapepod/pkg/models"
)

var (
	clean bool
)

var dbCmd = &cobra.Command{
	Use: "db",
	Run: func(cmd *cobra.Command, args []string) {
		dbctx, err := db.OpenDB()
		if err != nil {
			log.Fatal(err)
		}
		if clean {
			if err := dbctx.DropTableIfExists(&models.Podcast{}).Error; err != nil {
				log.Fatal(err)
			}
			log.Info("Cleaned database")
		}
		if err := models.Migrate(dbctx).Error; err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)

	dbCmd.PersistentFlags().BoolVarP(&clean, "clean", "c", false, ``)
}
