# GORM DM8 Driver

达梦数据库`gorm`驱动

## Quick Start

```go
import (
  "github.com/nfjBill/gorm-driver-dm"
  "gorm.io/gorm"
)

// https://github.com/nfjBill/gorm-driver-dm
dsn := "dm://SYSDBA:SYSDBA@127.0.0.1:5236?autoCommit=true"
db, err := gorm.Open(dm.Open(dsn), &gorm.Config{})
```

达梦数据库用户名即模式名

## 入门

- `clone`本项目到本地，进入项目目录
- 修改`dsn`调整为本地正确信息
- 运行`go test -v`
- 如果无报错，参考`dm_test.go`使用即可

## 注意事项

- 超过4096长度字符串，需要使用`dmSchema.Clob`，不超过使用`string`即可
- 列名不要使用达梦关键字，否则会出现错误
