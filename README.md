# go工具包

## 下载
```shell
go get github.com/loveyu233/goutils/imagecode@v2.0.0
```

### imagecode
```go
import (
	"fmt"
	"github.com/loveyu233/goutils/imagecode"
)

func main() {
	resource := imagecode.NewCaptchaResource(imagecode.WithPath("./code"))
	code := resource.CreateCode()
	fmt.Println(code.Y)
	fmt.Println(code.X)
	fmt.Println(code.SliderImg)
	fmt.Println(code.BgImg)
}

```
## conn
```go
    addr: ip:端口
```

## email
```go
IsValidEmail: 校验email格式
SendEmailBody： 只发送
SendEmailBodyAndSaveRedis： 发送并把body内容保存到redis
```

## jwt
```go
CreateToken: 创建token
GetPayload： 获取负载
VerifyToken： 验证token
```

## log
```go
InitZapLogger: 初始化zap
```

## random
```go
GeneratesSpecifiedNumberOfDigitsRandom: 生成指定位数的随机数
GeneratesRandomNumbersOfSpecifiedSize: 生成指定大小的随机数
```