package main

import (
	"context"
	"fmt"
	"image/color"
	"log/slog"
	"net/http"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/spf13/cobra"
	"github.com/way-platform/vin-go"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	if err := fang.Execute(
		context.Background(),
		newRootCommand(),
		fang.WithColorSchemeFunc(func(c lipgloss.LightDarkFunc) fang.ColorScheme {
			base := c(lipgloss.Black, lipgloss.White)
			baseInverted := c(lipgloss.White, lipgloss.Black)
			return fang.ColorScheme{
				Base:         base,
				Title:        base,
				Description:  base,
				Comment:      base,
				Flag:         base,
				FlagDefault:  base,
				Command:      base,
				QuotedString: base,
				Argument:     base,
				Help:         base,
				Dash:         base,
				ErrorHeader:  [2]color.Color{baseInverted, base},
				ErrorDetails: base,
			}
		}),
	); err != nil {
		os.Exit(1)
	}
}

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vin [command]",
		Short: "A VIN decoder CLI tool",
		Long:  `vin is a command-line tool for decoding Vehicle Identification Numbers (VINs).`,
	}
	cmd.AddGroup(&cobra.Group{ID: "decode", Title: "Decode Commands"})
	cmd.AddCommand(newDecodeCmd())
	cmd.AddGroup(&cobra.Group{ID: "serve", Title: "Serve Commands"})
	cmd.AddCommand(newServeCommand())
	cmd.AddGroup(&cobra.Group{ID: "util", Title: "Utility Commands"})
	cmd.SetCompletionCommandGroupID("util")
	cmd.SetHelpCommandGroupID("util")
	return cmd
}

func newDecodeCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "decode [vin...]",
		Short:   "Decode one or more VINs",
		GroupID: "decode",
		Args:    cobra.MinimumNArgs(1),
		RunE: (func(cmd *cobra.Command, args []string) error {
			for _, v := range args {
				decoded, err := vin.Decode(v)
				if err != nil {
					fmt.Printf("%s: %v\n", v, err)
					continue
				}
				fmt.Println(protojson.Format(decoded))
			}
			return nil
		}),
	}
}

func newServeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "serve",
		Short:   "Serve a VIN decoder HTTP API",
		GroupID: "serve",
	}
	port := cmd.Flags().Int("port", 8080, "The port to serve the API on")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		mux := http.NewServeMux()
		mux.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
			vinArg := r.URL.Query().Get("vin")
			decoded, err := vin.Decode(vinArg)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			data, err := protojson.Marshal(decoded)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if _, err := w.Write(data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
		slog.Info("serving VIN decoder HTTP API", slog.Int("port", *port))
		return http.ListenAndServe(fmt.Sprintf(":%d", *port), mux)
	}
	return cmd
}
