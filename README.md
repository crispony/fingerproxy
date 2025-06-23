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
  - [x] 通过 header 实现请求级别的动态指纹伪装
  - [x] 对 header 指纹进行模拟
  - [x] 通过 header 实现请求级别的代理切换
  - [ ] 对 QUIC 指纹进行模拟
  - [ ] 对 websocket 指纹进行模拟
  - [ ] 开放插件功能,允许开发者对app 进行定制化开发

## app 目录
```
dist 目录下面是编译好的可执行文件
```
## key table

| key | value |
| --- | --- |
| gospider007_fingerproxy_spec | fp 指纹 |
| gospider007_proxy_spec | 代理 |

## quick start
```bash
./dist/myapp_darwin_arm64_v8.0/myapp 
```
```bash
listening on: 127.0.0.1:8080
```

# python sdk
```python
import requests
response=requests.get("https://tools.scrapfly.io/api/fp/anything",proxies={"http": "http://127.0.0.1:8080","https": "http://127.0.0.1:8080"},verify=False,headers={
    "gospider007_fingerproxy_spec":"16030106f2010006ee03039a2b98d81139db0e128ea09eff6874549c219b543fb6dbaa7e4dbfe9e31602c620ce04c4026f019442affade7fed8ba66e022e186f77f1c670fd992f33c0143f120020aaaa130113021303c02bc02fc02cc030cca9cca8c013c014009c009d002f0035010006851a1a00000010000e000c02683208687474702f312e31002b000706dada03040303002d00020101000d0012001004030804040105030805050108060601001b0003020002ff0100010000230000000a000c000afafa11ec001d001700180000000e000c0000096c6f63616c686f7374003304ef04edfafa00010011ec04c06903195f3660633741f9f5ae64d05a316ac8e717582adb9e58c5c242ba306c99ca8f68c15f261245f9141812383c9265f7d0b5c44a5e0d7633f4f40ab8e820ec01bb6cc74e6ab3168e66b40cc4cd37e96e286b4080552b8b0217f786b7c1a0088fb613cc84471b17a33fbcc68db151df387907a1cf3fb14a0f45c6b84608db5b103131b537255c09559cf1940c3980a7f37959a7f95d32c49923600c76c616af238c579361e1c6a20c251d3d42a50e182ad54b1e5d54a57fe6986e64142815b9478cca8066c9bdcc0eb9022ebe05b0ebd53e7d146761d81aee41cd377611699c536a444d300c152994bfeef3cfb4736cb3d57d683269c1e3c001fba2ac220b2c993cec410f6fa5104d1bcde21c46a9b0be8ab7b51aba15f1745ebbd0a0d3a5224170b3bce456c157937e43390b733375bb96601667f5b36888f9520931bead48bae4723d9ed40af2680746e27eee4328503a280ee8846b37803d6206e0f5248bea4ca4a53ca4e1afdd2b84bce0c83260333bc9b38b86486fb48e18d1ce9187b1b6332b2f4145eca38122ac363210137f52140a57b7976b609d739844fa61f21c1e3c300f3434bc3f8b6856994847e2c9b0f20a24b976f9d552b153246db69f30d4a95301b933b0ba9d48402ddb7863cb4f1923a2c33021fd68634a387bf0f76d87f01b35b6182dd10fe9c14b8e548d7988388308b08ff1585ba18a7615737857a7c23c24ee9b3a2ab1915be18b233acd354c7c6513b8ea617a5cf299f34139756cc1df524292c43ee3364990961b2490ae204634a4461b53c95a11d214503985f27ab85bc7c179ab1ba37a828312cabea8cfc5088616386a83e566279a0a5517b60aca4ec6c30a32191dca3cbb7d33ae5087bbdbab5c42e6293b63ad8e35311d459e1ce57037e65e96283e449c3e012051d247653197834c42613ac377a950607c98cbd5a79fef948d18e99758e12d31d13231e638cfb183623346b231a443f56533d444c7204a63479e4efb34ef97597d858915a8a10a32aa78824c9d2993741176da643be1c6c4d91b6511055b098477e7a3c5b5c312cf7bb4ef5905fd741375e62c8ce942a2117bbe707fcc9871e59c0687507862f3634c871885949fce97612793a30a155e84ac503dc519816a13772c50e1167a7031cb2c8187913108f9a26e55958fd19a6e1ef18ed53a70f1c13d01d71d3b40a22413852c9982daac4ae8071966016c38a60bd5c0258c32b882740c6ed5252093c91e51c50cf037d4f5cb6ca610672710f7ca77b0a76039a9968e368a6b243ee4ca7632855cb568c73f01764c4944fc5879d2c52d7992840863c057db2efec658eeb2a73e02bd62617438d9192911bac1f6b0e55cb38255417af20000d69378c857bb278156f16a684200125906b6c22f3d505bc9e76d75fac3a009332ff98fe6baabe3941cab5271c6d2c0ebc993b944c49bd437353019d1b24d10390e45fa87ad77b329a9025933a11af2af0da44d3ed761722c94d8053242f537624113d7bc0155600573301bd2217c6c481ce63b0944b052c97bcb9d3349258257ff33cccf963a6945119ecab21c25051ce02548f642e0ec1ffd392d60facfdf76bfb7274363b62979231f4996362c85d5ba19d2cab7019750b3443565436867a53b71d875eba3282e6d0ee22076d6b97b7c6c556ae216e8bc1bc9f202ce94c763bfe9afc105fca9372dec2e286a001d00200ee8ac33f1ea3153f6b4a06ab71d21b7ce7955ce64ccfc66b7ec8077d02ffd18fe0d00fa00000100018b00206cada2aa48ee4478c40adad21f147d6bc90d13f6889a9b8a58a02536585a261f00d09306c85aab2a6e424b658f3cd9d1c46f35020839287259d3be605ff97faea0d87b9f7f96529661f08cf3f3899db8e805ee7405e2f9b6abd99bc4f6fa5f99b1ed442ebe53c5b10451c93d1221f662783efc3cc8fcf135ed935bcf02ec32251dd09705f191bd7959afbab5619d8e63cb634a259dd63d1b0e42225ae8c08b5b1620cd59d914857e9f1e8a3b7b892863bdaa05429922d75583059641468d8fc51c01e977a69d3a51d714cd5cceea9a5f404ce4a285fd6647931ed8b1c12db027328f214afdbe2c8102b46fe041b553f8670b00050005010000000044cd0005000302683200170000000b0002010000120000caca000100@@505249202a20485454502f322e300d0a0d0a534d0d0a0d0a00001804000000000000010001000000020000000000040060000000060004000000000408000000000000ef00010001d401250000000180000000ff82418aa0e41d139d09b8f3efbf87845887a47e561cc5801f40874148b1275ad1ffb9fe749d3fd4372ed83aa4fe7efbc1fcbefff3f4a7f388e79a82a97a7b0f497f9fbef07f21659fe7e94fe6f4f61e935b4ff3f7de0fe42cb3fcff408b4148b1275ad1ad49e33505023f30408d4148b1275ad1ad5d034ca7b29f07226d61634f53224092b6b9ac1c8558d520a4b6c2ad617b5a54251f01317ad9d07f66a281b0dae053fad0321aa49d13fda992a49685340c8a6adca7e28104416e277fb521aeba0bc8b1e632586d975765c53facd8f7e8cff4a506ea5531149d4ffda97a7b0f49580b2cae05c0b814dc394761986d975765cf53e5497ca589d34d1f43aeba0c41a4c7a98f33a69a3fdf9a68fa1d75d0620d263d4c79a68fbed00177fe8d48e62b03ee697e8d48e62b1e0b1d7f46a4731581d754df5f2c7cfdf6800bbdf43aeba0c41a4c7a9841a6a8b22c5f249c754c5fbef046cfdf6800bbbf408a4148b4a549275906497f83a8f517408a4148b4a549275a93c85f86a87dcd30d25f408a4148b4a549275ad416cf023f31408a4148b4a549275a42a13f8690e4b692d49f50929bd9abfa5242cb40d25fa523b3e94f684c9f518cf73ad7b4fd7b9fefb4005dff4086aec31ec327d785b6007d286f"
})
print(response.text)
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
