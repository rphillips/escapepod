package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"

	"github.com/rphillips/escapepod/pkg/routes"
)

var routesCmd = &cobra.Command{
	Use:   "routes",
	Short: "Dump available HTTP routes",
	Run: func(cmd *cobra.Command, args []string) {
		r := echo.New()
		fs, _ := loadFilesystem(os.Args[0])
		routes.RegisterRoutes(r, fs)
		data, _ := json.MarshalIndent(r.Routes(), "", "  ")
		fmt.Println(string(data))
	},
}

func init() {
	rootCmd.AddCommand(routesCmd)
}
