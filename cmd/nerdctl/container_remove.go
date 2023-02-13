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
	"github.com/containerd/nerdctl/pkg/api/types"
	"github.com/containerd/nerdctl/pkg/clientutil"
	"github.com/containerd/nerdctl/pkg/cmd/container"
	"github.com/spf13/cobra"
)

func newRmCommand() *cobra.Command {
	var rmCommand = &cobra.Command{
		Use:               "rm [flags] CONTAINER [CONTAINER, ...]",
		Args:              cobra.MinimumNArgs(1),
		Short:             "Remove one or more containers",
		RunE:              rmAction,
		ValidArgsFunction: rmShellComplete,
		SilenceUsage:      true,
		SilenceErrors:     true,
	}
	rmCommand.Flags().BoolP("force", "f", false, "Force the removal of a running|paused|unknown container (uses SIGKILL)")
	rmCommand.Flags().BoolP("volumes", "v", false, "Remove volumes associated with the container")
	return rmCommand
}

func rmAction(cmd *cobra.Command, args []string) error {
	globalOptions, err := processRootCmdFlags(cmd)
	if err != nil {
		return err
	}
	force, err := cmd.Flags().GetBool("force")
	if err != nil {
		return err
	}
	removeAnonVolumes, err := cmd.Flags().GetBool("volumes")
	if err != nil {
		return err
	}
	options := types.ContainerRemoveOptions{
		GOptions: globalOptions,
		Force:    force,
		Volumes:  removeAnonVolumes,
		Stdout:   cmd.OutOrStdout(),
	}

	client, ctx, cancel, err := clientutil.NewClient(cmd.Context(), options.GOptions.Namespace, options.GOptions.Address)
	if err != nil {
		return err
	}
	defer cancel()

	return container.Remove(ctx, client, args, options)
}

func rmShellComplete(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// show container names
	return shellCompleteContainerNames(cmd, nil)
}
