<p align="center">
  <a href="https://github.com/gospider007/requests"><img src="https://go.dev/images/favicon-gopher.png"></a>
</p>
<p align="center"><strong>Proxy</strong> <em>- A next-generation fingerprint proxy for Golang.</em></p>
<p align="center">
<a href="https://github.com/gospider007/requests">
    <img src="https://img.shields.io/github/last-commit/gospider007/requests">
</a>
<a href="https://github.com/gospider007/requests">
    <img src="https://img.shields.io/badge/build-passing-brightgreen">
</a>
<a href="https://github.com/gospider007/requests">
    <img src="https://img.shields.io/badge/language-golang-brightgreen">
</a>
</p>

---
<h2 align="center">A crawler-focused forward proxy with fingerprint spoofing</h2>

<h2 align="center">Fingerproxy is a fully featured Golang-based forward proxy app with browser fingerprint spoofing.
With just a few lines of code, you can enable powerful fingerprint proxy capabilities.
It provides unified support for HTTP/1, HTTP/2, HTTP/3, uTLS, and uQUIC protocols.</h2>

## Features
  - [x] 同一个端口同时支持 http,https,socks5 代理协议
  - [x] 所有支持http,https,socks5中任意一个代理协议的应用,都能受到指纹保护
  - [x] 通过配置文件,实现指纹的伪装
  - [x] 使所有不支持http2协议的http客户端支持http2协议
  - [x] 使所有不支持http3协议的http客户端支持http3协议
  - [x] 对tls 指纹进行模拟
  - [x] 对http2 指纹进行模拟
  - [ ] 通过 header 实现请求级别的动态指纹伪装
  - [ ] 通过 header 实现请求级别的代理切换
  - [ ] 对 header 指纹进行模拟
  - [ ] 对 QUIC 指纹进行模拟
  - [ ] 对 websocket 指纹进行模拟
  - [ ] 开放插件功能,允许开发者对app 进行定制化开发
## 项目说明
```
任务清单未完成的部分已经经过技术验证通过了，后续将逐步完善，敬请期待。
```

## app 目录
```
dist 目录下面是编译好的可执行文件
```

## 配置文件
```yaml
users: "" #代理用户名
passwords: "" #代理密码
address: ":8080" #监听地址
proxy: "" #代理
spec: "" #指纹
```
## quick start
```bash
./dist/myapp_darwin_arm64_v8.0/myapp --config config.yaml 
./myapp -config config.yaml #配置文件的启动方式
./myapp  #默认配置文件启动方式
```


# Contributing
If you have a bug report or feature request, you can [open an issue](../../issues/new)
# Contact
If you have questions, feel free to reach out to us in the following ways:
* QQ Group (Chinese): 939111384 - <a href="http://qm.qq.com/cgi-bin/qm/qr?_wv=1027&k=yI72QqgPExDqX6u_uEbzAE_XfMW6h_d3&jump_from=webapi"><img src="https://pub.idqqimg.com/wpa/images/group.png"></a>
* WeChat (Chinese): gospider007

## Sponsors
If you like and it really helps you, feel free to reward me with a cup of coffee, and don't forget to mention your github id.
<table>
    <tr>
        <td align="center">
            <img src="https://github.com/gospider007/tools/blob/master/play/wx.jpg?raw=true" height="200px" width="200px"   alt=""/>
            <br />
            <sub><b>Wechat</b></sub>
        </td>
        <td align="center">
            <img src="https://github.com/gospider007/tools/blob/master/play/qq.jpg?raw=true" height="200px" width="200px"   alt=""/>
            <br />
            <sub><b>Alipay</b></sub>
        </td>
    </tr>
</table>

## License  
This project is licensed under the Mozilla Public License 2.0 (MPL-2.0) with additional author attribution requirements.  
See the [LICENSE](./LICENSE) file for details.  
