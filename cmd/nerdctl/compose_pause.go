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
	"sync"

	"github.com/containerd/containerd"
	"github.com/containerd/nerdctl/pkg/clientutil"
	"github.com/containerd/nerdctl/pkg/cmd/compose"
	"github.com/containerd/nerdctl/pkg/containerutil"
	"github.com/containerd/nerdctl/pkg/labels"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func newComposePauseCommand() *cobra.Command {
	var composePauseCommand = &cobra.Command{
		Use:                   "pause [SERVICE...]",
		Short:                 "Pause all processes within containers of service(s). They can be unpaused with nerdctl compose unpause",
		RunE:                  composePauseAction,
		SilenceUsage:          true,
		SilenceErrors:         true,
		DisableFlagsInUseLine: true,
	}
	return composePauseCommand
}

func composePauseAction(cmd *cobra.Command, args []string) error {
	globalOptions, err := processRootCmdFlags(cmd)
	if err != nil {
		return err
	}

	client, ctx, cancel, err := clientutil.NewClient(cmd.Context(), globalOptions.Namespace, globalOptions.Address)
	if err != nil {
		return err
	}
	defer cancel()
	options, err := getComposeOptions(cmd, globalOptions.DebugFull, globalOptions.Experimental)
	if err != nil {
		return err
	}
	c, err := compose.New(client, globalOptions, options, cmd.OutOrStdout(), cmd.ErrOrStderr())
	if err != nil {
		return err
	}

	serviceNames, err := c.ServiceNames(args...)
	if err != nil {
		return err
	}
	containers, err := c.Containers(ctx, serviceNames...)
	if err != nil {
		return err
	}

	stdout := cmd.OutOrStdout()
	var mu sync.Mutex

	eg, ctx := errgroup.WithContext(ctx)
	for _, c := range containers {
		c := c
		eg.Go(func() error {
			if err := containerutil.Pause(ctx, client, c.ID()); err != nil {
				return err
			}
			info, err := c.Info(ctx, containerd.WithoutRefreshedMetadata)
			if err != nil {
				return err
			}

			mu.Lock()
			defer mu.Unlock()
			_, err = fmt.Fprintf(stdout, "%s\n", info.Labels[labels.Name])

			return err
		})
	}

	return eg.Wait()
}

func newComposeUnpauseCommand() *cobra.Command {
	var composeUnpauseCommand = &cobra.Command{
		Use:                   "unpause [SERVICE...]",
		Short:                 "Unpause all processes within containers of service(s).",
		RunE:                  composeUnpauseAction,
		SilenceUsage:          true,
		SilenceErrors:         true,
		DisableFlagsInUseLine: true,
	}
	return composeUnpauseCommand
}

func composeUnpauseAction(cmd *cobra.Command, args []string) error {
	globalOptions, err := processRootCmdFlags(cmd)
	if err != nil {
		return err
	}
	client, ctx, cancel, err := clientutil.NewClient(cmd.Context(), globalOptions.Namespace, globalOptions.Address)
	if err != nil {
		return err
	}
	defer cancel()

	options, err := getComposeOptions(cmd, globalOptions.DebugFull, globalOptions.Experimental)
	if err != nil {
		return err
	}
	c, err := compose.New(client, globalOptions, options, cmd.OutOrStdout(), cmd.ErrOrStderr())
	if err != nil {
		return err
	}
	serviceNames, err := c.ServiceNames(args...)
	if err != nil {
		return err
	}
	containers, err := c.Containers(ctx, serviceNames...)
	if err != nil {
		return err
	}

	stdout := cmd.OutOrStdout()
	var mu sync.Mutex

	eg, ctx := errgroup.WithContext(ctx)
	for _, c := range containers {
		c := c
		eg.Go(func() error {
			if err := unpauseContainer(ctx, client, c.ID()); err != nil {
				return err
			}
			info, err := c.Info(ctx, containerd.WithoutRefreshedMetadata)
			if err != nil {
				return err
			}

			mu.Lock()
			defer mu.Unlock()
			_, err = fmt.Fprintf(stdout, "%s\n", info.Labels[labels.Name])

			return err
		})
	}

	return eg.Wait()
}
