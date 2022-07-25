package sqlmock

import (
	"database/sql/driver"
	"reflect"
)

const returnStatusParamType = "*mssql.ReturnStatus"

// WillReturnStatus allows to specify an expected return code
// for stored procedure invocation
func (e *ExpectedQuery) WillReturnStatus(rs int32) *ExpectedQuery {
	e.rs = &rs
	return e
}

// WillReturnStatus allows to specify an expected return code
// for stored procedure invocation
func (e *ExpectedExec) WillReturnStatus(rs int32) *ExpectedExec {
	e.rs = &rs
	return e
}

func checkReturnStatusParam(v driver.Value) bool {
	return reflect.TypeOf(v).String() == returnStatusParamType
}

func setReturnStatus(args []driver.NamedValue, rs int32) {
	argsCount := len(args)
	if argsCount > 0 && checkReturnStatusParam(args[argsCount-1].Value) {
		reflect.ValueOf(args[argsCount-1].Value).Elem().SetInt(int64(rs))
	}
}

func filterReturnStatusParam(args []driver.NamedValue) (ret []driver.NamedValue) {
	for _, a := range args {
		if !checkReturnStatusParam(a.Value) {
			ret = append(ret, a)
		}
	}
	return ret
}
