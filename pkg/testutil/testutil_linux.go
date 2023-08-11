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

package testutil

import (
	"fmt"
	"io"
	"sync"
	"time"
)

func mirrorOf(s string) string {
	// plain mirror, NOT stargz-converted images
	return fmt.Sprintf("ghcr.io/stargz-containers/%s-org", s)
}

var (
	BusyboxImage                = "ghcr.io/containerd/busybox:1.28"
	AlpineImage                 = mirrorOf("alpine:3.13")
	NginxAlpineImage            = mirrorOf("nginx:1.19-alpine")
	NginxAlpineIndexHTMLSnippet = "<title>Welcome to nginx!</title>"
	RegistryImage               = mirrorOf("registry:2")
	WordpressImage              = mirrorOf("wordpress:5.7")
	WordpressIndexHTMLSnippet   = "<title>WordPress &rsaquo; Installation</title>"
	MariaDBImage                = mirrorOf("mariadb:10.5")
	DockerAuthImage             = mirrorOf("cesanta/docker_auth:1.7")
	FluentdImage                = mirrorOf("fluent/fluentd:v1.14-1")
	KuboImage                   = mirrorOf("ipfs/kubo:v0.16.0")

	// Source: https://gist.github.com/cpuguy83/fcf3041e5d8fb1bb5c340915aabeebe0
	NonDistBlobImage = "ghcr.io/cpuguy83/non-dist-blob:latest"
	// Foreign layer digest
	NonDistBlobDigest = "sha256:be691b1535726014cdf3b715ff39361b19e121ca34498a9ceea61ad776b9c215"

	CommonImage = AlpineImage

	// This error string is expected when attempting to connect to a TCP socket
	// for a service which actively refuses the connection.
	// (e.g. attempting to connect using http to an https endpoint).
	// It should be "connection refused" as per the TCP RFC.
	// https://www.rfc-editor.org/rfc/rfc793
	ExpectedConnectionRefusedError = "connection refused"
)

const (
	FedoraESGZImage = "ghcr.io/stargz-containers/fedora:30-esgz"            // eStargz
	FfmpegSociImage = "public.ecr.aws/soci-workshop-examples/ffmpeg:latest" // SOCI
)

type delayOnceReader struct {
	once    sync.Once
	wrapped io.Reader
}

// NewDelayOnceReader returns a wrapper around io.Reader that delays the first Read() by one second.
// It is used to test detaching from a container, and the reason why we need this is described below:
//
// Since detachableStdin.closer cancels the corresponding container's IO,
// it has to be invoked after the corresponding task is started,
// or the container could be resulted in an invalid state.
//
// However, in taskutil.go, the goroutines that copy the container's IO start
// right after container.NewTask(ctx, ioCreator) is invoked and before the function returns,
// which means that detachableStdin.closer could be invoked before the task is started,
// and that's indeed the case for e2e test as the detach keys are "entered immediately".
//
// Since detaching from a container is only applicable when there is a TTY,
// which usually means that there's a human in front of the computer waiting for a prompt to start typing,
// it's reasonable to assume that the user will not type the detach keys before the task is started.
func NewDelayOnceReader(wrapped io.Reader) io.Reader {
	return &delayOnceReader{wrapped: wrapped}
}

func (r *delayOnceReader) Read(p []byte) (int, error) {
	r.once.Do(func() { time.Sleep(time.Second) })
	return r.wrapped.Read(p)
}
