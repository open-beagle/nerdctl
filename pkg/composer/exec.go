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

package composer

import (
	"context"
	"fmt"
	"os"

	"github.com/containerd/containerd"
	"github.com/sirupsen/logrus"
)

// ExecOptions stores options passed from users as flags and args.
type ExecOptions struct {
	ServiceName string
	Index       int
	// params to be passed to `nerdctl exec`
	Detach      bool
	Interactive bool
	Tty         bool
	Privileged  bool
	User        string
	WorkDir     string
	Env         []string
	Args        []string
}

// Exec executes a given command on a running container specified by
// `ServiceName` (and `Index` if it has multiple instances).
func (c *Composer) Exec(ctx context.Context, eo ExecOptions) error {
	containers, err := c.Containers(ctx, eo.ServiceName)
	if err != nil {
		return fmt.Errorf("fail to get containers for service %s: %s", eo.ServiceName, err)
	}
	if len(containers) == 0 {
		return fmt.Errorf("no running containers from service %s", eo.ServiceName)
	}
	if eo.Index > len(containers) {
		return fmt.Errorf("index (%d) out of range: only %d running instances from service %s",
			eo.Index, len(containers), eo.ServiceName)
	}
	container := containers[eo.Index-1]

	return c.exec(ctx, container, eo)
}

// exec constructs/executes the `nerdctl exec` command to be executed on the given container.
func (c *Composer) exec(ctx context.Context, container containerd.Container, eo ExecOptions) error {
	args := []string{
		"exec",
		fmt.Sprintf("--detach=%t", eo.Detach),
		fmt.Sprintf("--interactive=%t", eo.Interactive),
		fmt.Sprintf("--tty=%t", eo.Tty),
		fmt.Sprintf("--privileged=%t", eo.Privileged),
	}
	if eo.User != "" {
		args = append(args, "--user", eo.User)
	}
	if eo.WorkDir != "" {
		args = append(args, "--workdir", eo.WorkDir)
	}
	for _, e := range eo.Env {
		args = append(args, "--env", e)
	}
	args = append(args, container.ID())
	args = append(args, eo.Args...)
	cmd := c.createNerdctlCmd(ctx, args...)

	if eo.Interactive {
		cmd.Stdin = os.Stdin
	}
	if !eo.Detach {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if c.DebugPrintFull {
		logrus.Debugf("Executing %v", cmd.Args)
	}
	return cmd.Run()
}
