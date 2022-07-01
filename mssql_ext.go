package sqlmock

import (
	"database/sql/driver"
	"reflect"
)

const returnStatusParamType = "*mssql.ReturnStatus"

func checkReturnStatusParam(v driver.Value) bool {
	return reflect.TypeOf(v).String() == returnStatusParamType
}

func setReturnStatus(args []driver.NamedValue, expectedReturnStatus int32) {
	argsCount := len(args)
	if argsCount > 0 && checkReturnStatusParam(args[argsCount-1].Value) {
		reflect.ValueOf(args[argsCount-1].Value).Elem().SetInt(int64(expectedReturnStatus))
	}
}
