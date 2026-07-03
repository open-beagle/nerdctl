# nerdctl

<https://github.com/containerd/nerdctl>

```bash
git remote add upstream git@github.com:containerd/nerdctl.git

git fetch upstream

git merge v2.1.6
```

## debug

```bash
# git remote
git -C ansible-docker-nerdctl merge v2.1.6

# git local
git -C ansible-docker-nerdctl checkout release-v2.1 && \
  git -C ansible-docker-nerdctl merge main --ff-only && \
  git -C ansible-docker-nerdctl push origin release-v2.1 && \
  git -C ansible-docker-nerdctl checkout main

# cache
docker run -it \
  --rm \
  -v $PWD/ansible-docker-nerdctl:/go/src/github.com/containerd/nerdctl \
  -w /go/src/github.com/containerd/nerdctl \
  registry.cn-qingdao.aliyuncs.com/wod/golang:1.24-bookworm \
  rm -rf vendor && go mod vendor

# build cross
docker run -it \
  --rm \
  -v $PWD/ansible-docker-nerdctl:/go/src/github.com/containerd/nerdctl \
  -w /go/src/github.com/containerd/nerdctl \
  -e BUILD_VERSION=v2.1.6 \
  registry.cn-qingdao.aliyuncs.com/wod/golang:1.24-bookworm \
  bash .beagle/build.sh
```
