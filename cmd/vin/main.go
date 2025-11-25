package main

import (
	"context"
	"fmt"
	"image/color"
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
	cmd.AddCommand(newDecodeCmd())
	return cmd
}

func newDecodeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "decode [vin...]",
		Short: "Decode one or more VINs",
		Args:  cobra.MinimumNArgs(1),
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
