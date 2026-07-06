# nerdctl

<https://github.com/containerd/nerdctl>

```bash
git -C ansible-docker-nerdctl remote add upstream git@github.com:containerd/nerdctl.git

git -C ansible-docker-nerdctl fetch upstream

git -C ansible-docker-nerdctl merge v2.1.6
```

## git

```bash
# ./.github/workflows/build-2.1.yml
git -C ansible-docker-nerdctl checkout release-v2.1 && \
git -C ansible-docker-nerdctl merge main && \
git -C ansible-docker-nerdctl push origin release-v2.1 && \
git -C ansible-docker-nerdctl checkout main
```

## debug

```bash
# cache
docker run -it --rm \
  -v $HOME/go/pkg:/go/pkg \
  -v $PWD/:/go/src/github.com/containerd/ \
  -v $PWD/ansible-docker-nerdctl:/go/src/github.com/containerd/nerdctl \
  -w /go/src/github.com/containerd/nerdctl \
  registry.cn-qingdao.aliyuncs.com/wod/golang:1.24-bookworm \
  rm -rf vendor && go mod vendor

# build cross
docker run -it --rm \
  -v $HOME/go/pkg:/go/pkg \
  -v $PWD/:/go/src/github.com/containerd/ \
  -v $PWD/ansible-docker-nerdctl:/go/src/github.com/containerd/nerdctl \
  -w /go/src/github.com/containerd/nerdctl \
  -e BUILD_VERSION=v2.1.6 \
  registry.cn-qingdao.aliyuncs.com/wod/golang:1.24-bookworm \
  bash .beagle/build.sh
```
