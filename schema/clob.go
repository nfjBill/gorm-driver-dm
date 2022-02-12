package dmSchema

import (
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/nfjBill/gorm-driver-dm/dmr"
)

type Clob string

// 写入数据库之前，对数据做类型转换
func (clob Clob) Value() (driver.Value, error) {
	if len(clob) == 0 {
		return nil, nil
	}
	return string(clob), nil
}

// 将数据库中取出的数据，赋值给目标类型
func (clob *Clob) Scan(v interface{}) error {
	switch v.(type) {
	case *dmr.DmClob:
		tmp := v.(*dmr.DmClob)
		le, err := tmp.GetLength()
		if err != nil {
			return errors.New(fmt.Sprint("err：", err))
		}

		str, err := tmp.ReadString(1, int(le))
		*clob = Clob(str)
		break

	//非clob，当成字符串，兼容oracle
	default:
		*clob = Clob(v.(string))
	}
	return nil
}
