## TCP 对接示例

本项目使用`gnet`作为TCP对接框架，[参考文档](https://github.com/panjf2000/gnet)。

> 本示例代码仅供参考。

[ASCII字符串到16进制在线转换工具](https://coding.tools/cn/ascii-to-hex)

## [TCP 内网穿透](https://github.com/fatedier/frp)
在 ``frpc.ini`` 文件中添加
```ini
[tcp]
name = "xxx"
type = "tcp"
local_ip = localhost # 本地IP
local_port = 7891 # 本地端口
remote_port = 7890 # 远程端口
server_addr = xxxx # 服务器地址
server_port = xxx # 服务器端口号
token = xxxx # token
```

## MQTT 示例

> 暂不提供

