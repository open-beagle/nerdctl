# nerdctl

<https://github.com/containerd/nerdctl>

```bash
git remote add upstream git@github.com:containerd/nerdctl.git

git fetch upstream

git merge v1.4.0
```

## debug

```bash
# cache
docker run -it \
--rm \
-v $PWD/:/go/src/github.com/containerd/nerdctl \
-w /go/src/github.com/containerd/nerdctl \
registry.cn-qingdao.aliyuncs.com/wod/golang:1.20-loongnix \
rm -rf vendor && go mod vendor

# patch vendor
git apply .beagle/0001-support-loong64.patch

# build
docker run -it \
--rm \
-v $PWD/:/go/src/github.com/containerd/nerdctl \
-w /go/src/github.com/containerd/nerdctl \
registry.cn-qingdao.aliyuncs.com/wod/golang:1.20-loongnix \
bash .beagle/build.sh

# check
file _output/linux/loong64/nerdctl
```

## cache

```bash
# 构建缓存-->推送缓存至服务器
docker run --rm \
  -e PLUGIN_REBUILD=true \
  -e PLUGIN_ENDPOINT=$PLUGIN_ENDPOINT \
  -e PLUGIN_ACCESS_KEY=$PLUGIN_ACCESS_KEY \
  -e PLUGIN_SECRET_KEY=$PLUGIN_SECRET_KEY \
  -e DRONE_REPO_OWNER="open-beagle" \
  -e DRONE_REPO_NAME="nerdctl" \
  -e PLUGIN_MOUNT="./.git,./vendor" \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  registry.cn-qingdao.aliyuncs.com/wod/devops-s3-cache:1.0

# 读取缓存-->将缓存从服务器拉取到本地
docker run --rm \
  -e PLUGIN_RESTORE=true \
  -e PLUGIN_ENDPOINT=$PLUGIN_ENDPOINT \
  -e PLUGIN_ACCESS_KEY=$PLUGIN_ACCESS_KEY \
  -e PLUGIN_SECRET_KEY=$PLUGIN_SECRET_KEY \
  -e DRONE_REPO_OWNER="open-beagle" \
  -e DRONE_REPO_NAME="nerdctl" \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  registry.cn-qingdao.aliyuncs.com/wod/devops-s3-cache:1.0
```
