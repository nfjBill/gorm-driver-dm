package dmSchema

import (
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/nfjBill/gorm-driver-dm/dmr"
)

type Clob string

func (clob Clob) Value() (driver.Value, error) {
	if len(clob) == 0 {
		return nil, nil
	}
	return string(clob), nil
}

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

	default:
		*clob = Clob(v.(string))
	}
	return nil
}
