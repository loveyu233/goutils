# go工具包

## 下载
```shell
go get github.com/loveyu233/goutils/imagecode@v1.0.0
```

## imagecode
> 图片滑块验证码

### 使用
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