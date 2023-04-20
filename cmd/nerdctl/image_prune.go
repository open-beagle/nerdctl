/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"fmt"
	"strings"

	"github.com/containerd/nerdctl/pkg/api/types"
	"github.com/containerd/nerdctl/pkg/clientutil"
	"github.com/containerd/nerdctl/pkg/cmd/image"
	"github.com/spf13/cobra"
)

func newImagePruneCommand() *cobra.Command {
	imagePruneCommand := &cobra.Command{
		Use:           "prune [flags]",
		Short:         "Remove unused images",
		Args:          cobra.NoArgs,
		RunE:          imagePruneAction,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	imagePruneCommand.Flags().BoolP("all", "a", false, "Remove all unused images, not just dangling ones")
	imagePruneCommand.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")
	return imagePruneCommand
}

func processImagePruneOptions(cmd *cobra.Command) (types.ImagePruneOptions, error) {
	globalOptions, err := processRootCmdFlags(cmd)
	if err != nil {
		return types.ImagePruneOptions{}, err
	}
	all, err := cmd.Flags().GetBool("all")
	if err != nil {
		return types.ImagePruneOptions{}, err
	}

	force, err := cmd.Flags().GetBool("force")
	if err != nil {
		return types.ImagePruneOptions{}, err
	}

	return types.ImagePruneOptions{
		Stdout:   cmd.OutOrStdout(),
		GOptions: globalOptions,
		All:      all,
		Force:    force,
	}, err
}

func imagePruneAction(cmd *cobra.Command, _ []string) error {
	options, err := processImagePruneOptions(cmd)
	if err != nil {
		return err
	}

	if !options.Force {
		var (
			confirm string
			msg     string
		)
		if !options.All {
			msg = "This will remove all dangling images."
		} else {
			msg = "This will remove all images without at least one container associated to them."
		}
		msg += "\nAre you sure you want to continue? [y/N] "

		fmt.Fprintf(cmd.OutOrStdout(), "WARNING! %s", msg)
		fmt.Fscanf(cmd.InOrStdin(), "%s", &confirm)

		if strings.ToLower(confirm) != "y" {
			return nil
		}
	}

	client, ctx, cancel, err := clientutil.NewClient(cmd.Context(), options.GOptions.Namespace, options.GOptions.Address)
	if err != nil {
		return err
	}
	defer cancel()

	return image.Prune(ctx, client, options)
}
