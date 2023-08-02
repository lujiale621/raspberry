# 树莓派系统初始化

1. 使用 imager_1.7.5 烧录树莓派系统
2. 将 `cdline.txt`、`config.txt`、`ssh`、`userconf.txt`、`wpa_supplicant.conf` 复制到 `boot` 目录

## 内网穿透

1. 在 [https://www.natfrp.com/user/](https://www.natfrp.com/user/) 上设置映射
2. 编译 `raspberry/frp/cmd/frpc/main.go`
3. 配置 `frpc.ini`
4. 编写 `frpc.service` 实现开机自启动

## 启动代理

1. 运行命令：`sudo ./lovelace_linux_armv6l`
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
go mod init xiaomihu
go mod tidy
GOOS=linux GOARCH=amd64 go build -o frps
GOOS=linux GOARM=6 GOARCH=arm go build -o armv6-main
go list -m all
go env
```

## 启动 frps

```shell
./frps -c ./frps.ini
```

## 安装 postgresql 数据库

```shell
sudo apt install postgresql libpq-dev postgresql-client postgresql-client-common -y
当安装完成后，切换到 Postgres 用户去配置数据库：

sudo su postgres
现在，你可以创建一个数据库用户。如果你创建了一个与你的 Unix 用户帐户相同名字的用户，那个用户将被自动授权访问该数据库。因此在本教程中，为简单起见，我们将假设你使用了默认用户 pi 。运行 createuser 命令以继续：

createuser pi -P --interactive
现在，使用 Postgres shell 连接到 Postgres 去创建一个测试数据库：

$ psql
> create database serverdata;
```

## 解析日志的仪表盘

```shell
apt-get install goaccess
```

## 仪表盘

```shell
sudo dnf install htop
htop
```

## docker 篇

```shell
要进入 Docker 容器，可以使用以下命令：

首先，使用 docker ps 命令查看正在运行的容器列表。找到你想要进入的容器的名称或 ID。

使用以下命令进入容器的交互式终端：

docker exec -it <容器名称或 ID> /bin/bash

docker ps -a
请注意，此命令只会对容器的下一个重启生效。因此，在设置内存限制后，您需要重启容器才能使更改生效。您可以使用以下命令重启容器：

sudo docker restart <container_id>
这将输出容器的日志，可以帮助你确定问题所在。另外，你还可以使用以下命令查看容器的健康状态：

sudo docker inspect --format='{{.State.Health}}' uptime-kuma
运行以下命令以停止特定的 Docker 容器：

sudo docker stop c2b3664d68bc
要删除 Docker 容器和镜像，可以按照以下步骤进行操作：

首先，删除退出状态的容器。运行以下命令以删除特定容器：

sudo docker rm c2b3664d68bc
这将删除容器 ID 为 c2b3664d68bc 的容器。请确保替换成你要删除的容器的实际 ID。

接下来，删除没有被使用的镜像。运行以下命令以删除特定镜像：

sudo docker rmi louislam/uptime-kuma:1
这将删除名称为 louislam/uptime-kuma:1 的镜像。同样，请根据需要替换成你要删除的镜像的实际名称和标签。

请注意，删除容器和镜像是不可逆的操作，请确保你真的想要删除它们。如果你需要保留某些数据或配置，请在删除之前进行备份
要查看当前系统中的所有 Docker 镜像，你可以使用以下命令：

sudo docker images
要在 Docker 中以守护进程（后台）模式运行容器，您可以使用 -d 或 --detach 参数
sudo docker run cloudflare/cloudflared:latest tunnel --no-autoupdate run --token eyJhIjoiYzcxZWY0OTAyNjBjNTIzNDY5YTNhYzM4YWVmNDA5YTIiLCJ0IjoiNmM2NTBiNDYtMmUxYy00YzRhLTgwN2YtMDE4YjFkYTUwNGM4IiwicyI6Ill6TmlOVFE1TnpRdE5tSTBOUzAwTWpCaUxXRXpaVE10T1RSbE1tTmxPREkxTW1FMyJ9
sudo docker network ls
sudo docker container ls
查看容器的ip
sudo docker inspect 245b3b06ea54
```

## docker-mysql

```
mysql -u <用户名> -p
CREATE DATABASE mydb;
USE mydb;
CREATE TABLE spriteinfo (
    id INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(255),
    Url VARCHAR(255),
    Width FLOAT,
    Height FLOAT,
    X FLOAT,
    Y FLOAT,
    Visable BOOLEAN
);
SHOW DATABASES;
DROP TABLE sprites;
SELECT * FROM sprites;
DESCRIBE sprites;
```

## docker-compose 篇

```shell
要使用 docker-compose 来启动 Docker 容器，你需要准备一个 docker-compose.yml 文件并确保已安装 Docker Compose 工具。下面是启动 Docker 容器的步骤：

在项目目录中创建一个名为 docker-compose.yml 的文件。

在 docker-compose.yml 中定义你要启动的容器服务。格式如下：

yaml
version: '3'
services:
  service1:
    image: <image_name>
    ...
  service2:
    image: <image_name>
    ...
替换 <image_name> 为你想要使用的镜像名称，并根据需要添加其他配置选项。

保存并关闭 docker-compose.yml 文件。

打开终端或命令行窗口，并进入包含 docker-compose.yml 文件的项目目录。

运行以下命令来启动容器：

docker-compose up
这将会根据 docker-compose.yml 文件中的定义启动指定的容器服务。如果镜像不存在，会尝试自动拉取相应的镜像。

如果你想在后台运行容器，可以使用 -d 或 --detach 选项：

docker-compose up -d
Docker Compose 将会自动创建和启动所需的容器。你可以观察到容器的日志输出。

要停止容器并删除相关资源，可以运行以下命令：

docker-compose down
```

## ssh

```shell
要修改 SSH 的配置文件，你可以按照以下步骤进行操作：

使用超级用户或具有适当权限的帐户登录到目标主机。

打开 SSH 配置文件（在大多数 Linux 发行版中为 /etc/ssh/sshd_config）。

例如，使用文本编辑器（如 vi 或 nano）打开配置文件：
sudo vi /etc/ssh/sshd_config
在打开的文件中，找到要修改的配置选项，并按需进行更改。根据需要，你可以修改以下常见的配置选项：

端口号：通过修改 Port 配置选项来更改 SSH 服务器监听的端口号。默认情况下，SSH 使用 22 端口。
允许登录的用户：通过修改 AllowUsers 或 AllowGroups 配置选项，限制可以登录到 SSH 的用户或组。
登录密码认证：通过修改 PasswordAuthentication 配置选项，启用或禁用基于密码的身份验证。
公钥认证：通过修改 PubkeyAuthentication 配置选项，启用或禁用基于公钥的身份验证。
允许 root 登录：通过修改 PermitRootLogin 配置选项，启用或禁用允许 root 用户登录。
配置选项通常有注释提供了说明和默认值。确保你仔细阅读并理解了配置选项的含义，以便正确地进行修改。

修改完配置后，保存文件并退出编辑器。

重新加载 SSH 配置，使更改生效。可以使用以下命令重新启动 SSH 服务：

shell
sudo systemctl restart ssh

```

每秒刷新命令
watch -n 1 cat nohup.out

## 防火墙

```shell
查看防火墙状态：
sudo systemctl status firewalld
启动防火墙：
sudo systemctl start firewalld
停止防火墙：
sudo systemctl stop firewalld
重新加载防火墙配置：
sudo firewall-cmd --reload
查看已开放的端口：
sudo firewall-cmd --list-ports
查看已开放的服务：
sudo firewall-cmd --list-services
```

```shell
以太网适配器 VMware Network Adapter VMnet8:

   连接特定的 DNS 后缀 . . . . . . . :
   本地链接 IPv6 地址. . . . . . . . : fe80::1124:747b:6094:a80a%14
   IPv4 地址 . . . . . . . . . . . . : 172.16.245.1
   子网掩码  . . . . . . . . . . . . : 255.255.255.0
   默认网关. . . . . . . . . . . . . :
子网 172.16.245.0
网关 172.16.245.20

gitlab搭建
# 拉取Gitlab镜像
docker pull gitlab/gitlab-ce:latest
# 启动容器
docker run \
 -itd  \
 -p 9980:80 \
 -p 9922:22 \
 -v /home/gitlab/etc:/etc/gitlab  \
 -v /home/gitlab/log:/var/log/gitlab \
 -v /home/gitlab/opt:/var/opt/gitlab \
 --restart always \
 --privileged=true \
 --name gitlab \
 gitlab/gitlab-ce

修改root密码
# 进入容器内部
docker exec -it gitlab /bin/bash

# 进入控制台
gitlab-rails console -e production

# 查询id为1的用户，id为1的用户是超级管理员
user = User.where(id:1).first
# 修改密码为lhx123456
user.password='lujiale'
# 保存
user.save!
# 退出
exit
```

备份镜像和容器

```
保存镜像

docker save ID > xxx.tar

docker load < xxx.tar

保存容器

docker export ID >xxx.tar

docker import xxx.tar containr:v1

然后再docker run -it containr:v1 bash
```

```
搭建mysql
docker pull mysql:latest
docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql
```

```搭建Nginx

第一步：下载Nginx最新版本的镜像

docker pull nginx:latest
第二步：运行Nginx镜像
docker run --name nginx -p 80:80 -d nginx
--name nginx 指定容器的名称
-p 80:80 映射端口，前面的80是指宿主机的端口，后面的80是指Docker容器里的端口
-d 守护进程运行
第三步：从nginx容器中映射核心文件。为什么要这样做呢？目的就是以后修改宿主机上对应的文件后，重启nginx容器后就生效了。不然每次都要进入nginx容器里面去操作，很麻烦的嘛。

A，在宿主机本地创建以下文件目录

mkdir -p /opt/docker/nginx/conf.d
mkdir -p /opt/docker/nginx/html
mkdir -p /opt/docker/nginx/logs
mkdir -p /opt/docker/nginx/conf
B，将nginx容器里对应的文件拷贝到上面创建的宿主机目录里

docker cp nginx:/etc/nginx/nginx.conf /opt/docker/nginx/conf
docker cp nginx:/etc/nginx/conf.d /opt/docker/nginx/conf.d
docker cp nginx:/usr/share/nginx/html /opt/docker/nginx
第四步：停止并删除原先的nginx容器

docker stop nginx
docker rm nginx
第五步：创建一个新的nginx容器

docker run  -p 80:80 --name nginx --restart=always
-v /opt/docker/nginx/conf/nginx.conf:/etc/nginx/nginx.conf
-v /opt/docker/nginx/conf.d:/etc/nginx/conf.d
-v /opt/docker/nginx/html:/usr/share/nginx/html
-v /opt/docker/nginx/logs:/var/log/nginx
-d  nginx

docker run  -p 80:80 --name nginx --restart=always \
-v /data1/nginx/conf/nginx.conf:/etc/nginx/nginx.conf \
-v /data1/nginx/conf.d:/etc/nginx/conf.d \
-v /data1/nginx/html:/usr/share/nginx/html \
-v /data1/nginx/logs:/var/log/nginx \
-d  nginx
```
