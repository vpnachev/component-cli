// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors.
//
// SPDX-License-Identifier: Apache-2.0

package componentarchive

import (
	"context"
	"fmt"
	"os"

	"github.com/mandelsoft/vfs/pkg/osfs"
	"github.com/mandelsoft/vfs/pkg/vfs"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/gardener/component-cli/pkg/componentarchive"
)

// CreateOptions defines all options for the create command.
type CreateOptions struct {
	componentarchive.BuilderOptions
}

// NewCreateCommand creates a new component descriptor
func NewCreateCommand(ctx context.Context) *cobra.Command {
	opts := &ExportOptions{}
	cmd := &cobra.Command{
		Use:   "create [component-archive-path]",
		Args:  cobra.ExactArgs(1),
		Short: "Creates a component archive with a component descriptor",
		Long: `
Create command creates a new component archive directory with a "component-descriptor.yaml" file.
`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := opts.Complete(args); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			if err := opts.Run(ctx, osfs.New()); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fmt.Printf("Successfully created component archive at %s\n", opts.OutputPath)
		},
	}
	opts.AddFlags(cmd.Flags())
	return cmd
}

// Run runs the export for a component archive.
func (o *CreateOptions) Run(ctx context.Context, fs vfs.FileSystem) error {
	_, err := o.BuilderOptions.Build(fs)
	return err
}

// Complete parses the given command arguments and applies default options.
func (o *CreateOptions) Complete(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected exactly one argument that contains the path to the component archive")
	}
	o.ComponentArchivePath = args[0]

	return o.validate()
}

func (o *CreateOptions) validate() error {
	return o.BuilderOptions.Validate()
}

func (o *CreateOptions) AddFlags(fs *pflag.FlagSet) {
	o.BuilderOptions.AddFlags(fs)
}
