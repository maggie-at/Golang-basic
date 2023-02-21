### Extension
- [x] GO
- [x] Code Runner


### 命令
```bash
go version
go env
go mod init <package>
go build xx.go
go run xx.go
go clean
go help
```

### 结构
> 入口: `main`包下的`main()`函数

```GO
package main

import "Golang/basic"

func main() {
	basic.Array_()
}
```