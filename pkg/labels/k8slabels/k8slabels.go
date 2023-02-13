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

// Package k8slabels defines Kubernetes container labels
package k8slabels

const (
	PodNamespace  = "io.kubernetes.pod.namespace"
	PodName       = "io.kubernetes.pod.name"
	ContainerName = "io.kubernetes.container.name"

	ContainerMetadataExtension = "io.cri-containerd.container.metadata"
	ContainerType              = "io.cri-containerd.kind"
)
