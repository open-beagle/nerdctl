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

package platform

import (
	"fmt"
)

func DataHome() (string, error) {
	panic("not supported")
}

// The following are here solely for windows to compile. They are not used, as the corresponding tests are running only on linux.
func mirrorOf(s string) string {
	return fmt.Sprintf("ghcr.io/stargz-containers/%s-org", s)
}

var (
	RegistryImageStable = mirrorOf("registry:2")
	RegistryImageNext   = "ghcr.io/distribution/distribution:"
	KuboImage           = mirrorOf("ipfs/kubo:v0.16.0")
	DockerAuthImage     = mirrorOf("cesanta/docker_auth:1.7")
)
