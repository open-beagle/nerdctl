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

package types

import "io"

// ApparmorListOptions specifies options for `nerdctl apparmor ls`.
type ApparmorListOptions struct {
	Stdout io.Writer
	// Only display profile names
	Quiet bool
	// Format the output using the given go template
	Format string
}

// ApparmorInspectOptions specifies options for `nerdctl apparmor inspect`
type ApparmorInspectOptions struct {
	Stdout io.Writer
}
