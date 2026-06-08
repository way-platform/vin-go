package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/way-platform/vin-go"
	"google.golang.org/protobuf/encoding/protojson"
)

// NewCommand builds the full CLI command tree for the VIN decoder.
func NewCommand(opts ...Option) *cobra.Command {
	cfg := config{}
	for _, opt := range opts {
		opt(&cfg)
	}
	cmd := &cobra.Command{
		Use:   "vin [command]",
		Short: "A VIN decoder CLI tool",
		Long:  `vin is a command-line tool for decoding Vehicle Identification Numbers (VINs).`,
	}
	cmd.AddGroup(&cobra.Group{ID: "decode", Title: "Decode Commands"})
	cmd.AddCommand(newDecodeCmd())
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

