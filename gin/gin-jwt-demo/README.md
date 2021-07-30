## gin-jwt-demo

在前后端分离的项目中，越来越多的项目采用 JWT 代替传统的 cookie ，这里我们来使用 JWT 结合 Gin 来作为一个登录授权和权限校验。

### 🔑什么是 JWT
JWT 的全称叫做 JSON WEB TOKEN，在目前前后端系统中使用较多。

#### JWT 构成
JWT 是由三段构成的。分别是 HEADER，PAYLOAD，VERIFY SIGNATURE，它们生成的信息通过 . 分割。

> HEADER
header 是由一个`typ`和`alg`组成，`typ`会指明为JWT，而`alg`是所使用的加密算法
```json
{
  "alg": "HS256",
  "typ": "JWT",
}
```

> PAYLOAD

payload 是 JWT 的载体，也就是我们要承载的信息。这段信息是我们可以自定义的，可以定义我们要存放的什么信息，那些字段。该部分信息不宜过多，他会影响JWT生成的发小，还有就是请勿将敏感数据存入该部分，该端数据前端可以解析获取token内信息的。

官方给了七个默认字段

|    名称     |     含义      |
|    ----    |     ----     |
| Audience   |  表示JWT的受众  |
| ExpiresAt  |    失效时间     |
|     Id    |    签发编号     |
| IssuedAt  |    签发时间     |
| Issuer  |    签发人     |
| NotBefore  |    生效时间     |
| Subject  |    主题     |

> VERIFY SIGNATURE

这也是JWT的最后一段，该部分是由算法计算完成的

对刚刚的 header 进行 base64Url 编码，对 payload 进行 base64Url 编码，两端完成编码后通过 `.` 进行连接起来。

```json
base64UrlEncode(header).base64UrlEncode(payload)
```

完成上述步骤后，就要通过我们 header 里指定的加密算法对上部分进行加密，同时我们还要插入我们的一个密钥，来确保我的 JWT 签发是安全的。

上述完成后，通过使用`.`将三部分分割，生成上图所示的 JWT。


### 🔒Gin 生成 JWT
go语言的JWT库有很多。[jwt.io](https://jwt.io/)上也给出了很多。这里使用jwt-go

