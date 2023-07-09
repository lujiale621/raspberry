# 树莓派系统初始化

1. 使用 imager_1.7.5 烧录树莓派系统
2. 将 `cdline.txt`、`config.txt`、`ssh`、`userconf.txt`、`wpa_supplicant.conf` 复制到 `boot` 目录

## 内网穿透

1. 在 [https://www.natfrp.com/user/](https://www.natfrp.com/user/) 上设置映射
2. 编译 `raspberry/frp/cmd/frpc/main.go`
3. 配置 `frpc.ini`
4. 编写 `frpc.service` 实现开机自启动

## 启动代理

1. 运行命令：`sudo ./lovelace`
2. 扫描端口：`scan 200`
3. 启动 Clash：`clash`

## 代理设置

在终端中执行以下命令：

```shell
export http_proxy=http://127.0.0.1:7899
export https_proxy=http://127.0.0.1:7899
```

## 安装 Docker

```shell
sudo apt-get update
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg2 \
    lsb-release \
    software-properties-common

curl -fsSL https://mirrors.aliyun.com/docker-ce/linux/raspbian/gpg | sudo apt-key add -
sudo apt-get install docker-ce
```

## 启动 Docker

```shell
sudo systemctl enable docker
sudo systemctl start docker
```

## 安装 Docker Compose

```shell
sudo apt install docker-compose
sudo ln -s /usr/bin/docker-compose /usr/local/bin/docker-compose
docker-compose version
```

## 安装 Python 依赖

```shell
pip install fastapi
pip install uvicorn
## 编译 Go 代码
go mod init <module_name>
go clean -modcache
```

## 如果你想获取 github.com/fatedier/golib 的最新提交而不是特定的版本标签，可以使用以下命令：

```shell
go get github.com/fatedier/golib@master

go mod tidy
GOOS=linux GOARCH=amd64 go build -o frps
go list -m all
go env
```

## 启动 frps

```shell
./frps -c ./frps.ini
```
