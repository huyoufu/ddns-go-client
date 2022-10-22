###一个简单的由go语言实现的ddns自动更新客户端工具

### 一、使用方式

#### 1.1 linux系统

```shell
mkdir -p /usr/local/ddns-go-client
cd -p /usr/local/ddns-go-client
wget linux版本下载路径
#初始化配置文件
./ddns-go-client -s init
#在当前目录下可以找到一个 (.ddns_go_config.json) 配置文件 注意带有.开头的文件 默认是隐藏文件

#编辑配置文件 填入想要自动报告的 域名和 对应ddns服务商的 访问秘钥

#安装为服务
./ddns-go-client -s install

#设置为开机自启动
systemctl enable ddns-go-client

#启动服务
systemctl start ddns-go-client

#查看服务状态
systemctl status ddns-go-client

#可以查看日志文件
tail -f ddns-go-client.log


```

#### 1.2 windows系统

```yacas
#windows系统制作系统服务太过麻烦 请自行搜索如何将应用程序添加到 开机启动计划中
```

#### 1.3 macos系统

```shell
#首先不太鼓励 做这个玩意儿 很少有人使用mac系统做服务器
#如果必须要玩

#创建软件存放目录
mkdir -p ~/ddns
cd ~/ddns
wget macos版本下载路径
#初始化配置文件
./ddns-go-client -s init
#在当前目录下可以找到一个 (.ddns_go_config.json) 配置文件 注意带有.开头的文件 默认是隐藏文件

#编辑配置文件 填入想要自动报告的 域名和 对应ddns服务商的 访问秘钥

#安装为服务
./ddns-go-client -s install


#可以查看日志文件
tail -f ddns-go-client.log
```

