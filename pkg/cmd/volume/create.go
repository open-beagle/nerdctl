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

package volume

import (
	"fmt"

	"github.com/containerd/containerd/identifiers"
	"github.com/containerd/nerdctl/pkg/api/types"
	"github.com/containerd/nerdctl/pkg/inspecttypes/native"
	"github.com/containerd/nerdctl/pkg/strutil"
)

func Create(name string, options types.VolumeCreateOptions) (*native.Volume, error) {
	if err := identifiers.Validate(name); err != nil {
		return nil, fmt.Errorf("malformed name %s: %w", name, err)
	}
	volStore, err := Store(options.GOptions.Namespace, options.GOptions.DataRoot, options.GOptions.Address)
	if err != nil {
		return nil, err
	}
	labels := strutil.DedupeStrSlice(options.Labels)
	vol, err := volStore.Create(name, labels)
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(options.Stdout, "%s\n", name)
	return vol, nil
}
