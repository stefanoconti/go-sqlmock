package sqlmock

import (
	"database/sql"
	"database/sql/driver"
	"reflect"
)

type inOutValue struct {
	Name             string
	ExpectedInValue  interface{}
	ReturnedOutValue interface{}
	In               bool
}

func (t inOutValue) Match(v driver.Value) bool {
	out, ok := v.(sql.Out)

	return ok && out.In == t.In && (!t.In || reflect.DeepEqual(out.Dest, t.ExpectedInValue))
}

func NamedOutputArg(name string, returnedOutValue interface{}) interface{} {
	return inOutValue{Name: name, ReturnedOutValue: returnedOutValue, In: false}
}

func NamedInputOutputArg(name string, expectedInValue interface{}, returnedOutValue interface{}) interface{} {
	return inOutValue{Name: name, ExpectedInValue: expectedInValue, ReturnedOutValue: returnedOutValue, In: true}
}

func setOutputValues(currentArgs []driver.NamedValue, expectedArgs []driver.Value) {
	for _, expectedArg := range expectedArgs {
		if outVal, ok := expectedArg.(inOutValue); ok {
			for _, currentArg := range currentArgs {
				if currentArg.Name == outVal.Name {
					if sqlOut, ok := currentArg.Value.(sql.Out); ok {
						reflect.ValueOf(sqlOut.Dest).Elem().Set(reflect.ValueOf(outVal.ReturnedOutValue))
					}
					break
				}
			}
		}
	}
}
