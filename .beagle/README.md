# nerdctl

<https://github.com/containerd/nerdctl>

```bash
git remote add upstream git@github.com:containerd/nerdctl.git

git fetch upstream

git merge v2.0.5
```

## debug

```bash
# cache
docker run -it \
  --rm \
  -v $PWD/:/go/src/github.com/containerd/nerdctl \
  -w /go/src/github.com/containerd/nerdctl \
  registry.cn-qingdao.aliyuncs.com/wod/golang:1.23 \
  rm -rf vendor && go mod vendor

# build cross
docker run -it \
  --rm \
  -v $PWD/:/go/src/github.com/containerd/nerdctl \
  -w /go/src/github.com/containerd/nerdctl \
  -e BUILD_VERSION=v2.0.5 \
  registry.cn-qingdao.aliyuncs.com/wod/golang:1.23 \
  bash .beagle/build.sh

# build loong64
docker run -it \
  --rm \
  -v $PWD/:/go/src/github.com/containerd/nerdctl \
  -w /go/src/github.com/containerd/nerdctl \
  -e BUILD_VERSION=v2.0.5 \
  registry.cn-qingdao.aliyuncs.com/wod/golang:1.23-loongnix \
  bash .beagle/build-loong64.sh

# check
file _output/linux/loong64/nerdctl
```

## cache

```bash
# 构建缓存-->推送缓存至服务器
docker run --rm \
  -e PLUGIN_REBUILD=true \
  -e PLUGIN_ENDPOINT=${S3_ENDPOINT_ALIYUN} \
  -e PLUGIN_ACCESS_KEY=${S3_ACCESS_KEY_ALIYUN} \
  -e PLUGIN_SECRET_KEY=${S3_SECRET_KEY_ALIYUN} \
  -e DRONE_REPO_OWNER="open-beagle" \
  -e DRONE_REPO_NAME="nerdctl" \
  -e PLUGIN_MOUNT="./.git" \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  registry.cn-qingdao.aliyuncs.com/wod/devops-s3-cache:1.0

# 读取缓存-->将缓存从服务器拉取到本地
docker run --rm \
  -e PLUGIN_RESTORE=true \
  -e PLUGIN_ENDPOINT=${S3_ENDPOINT_ALIYUN} \
  -e PLUGIN_ACCESS_KEY=${S3_ACCESS_KEY_ALIYUN} \
  -e PLUGIN_SECRET_KEY=${S3_SECRET_KEY_ALIYUN} \
  -e DRONE_REPO_OWNER="open-beagle" \
  -e DRONE_REPO_NAME="nerdctl" \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  registry.cn-qingdao.aliyuncs.com/wod/devops-s3-cache:1.0
```
