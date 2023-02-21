### go module

##### 初始化模块

```bash
go mod init <项目模块名称>
```

##### 依赖关系处理(根据go.mod文件)

```bash
go mod tidy
```

##### 将依赖包复制到项目的vendor目录

```bash
go mod vendor
```

##### 显示依赖关系

```bash
go list -m all
go list -m json all     # 详细
```

##### 下载依赖

```bash
go mod download [path@version]  # [path@version]为选填
```