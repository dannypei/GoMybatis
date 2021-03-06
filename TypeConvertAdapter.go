package GoMybatis

import (
	"bytes"
	"fmt"
	"reflect"
	"time"
)

const Adapter_DateType = `time.Time`
const Adapter_StringType = `string`
const Adapter_FormateDate = `2006-01-02 15:04:05`

var DefaultExpressionTypeConvertFunc = func(arg interface{}) interface{} {
	if reflect.TypeOf(arg).String() == Adapter_DateType {
		return arg.(time.Time).Nanosecond()
	}
	return arg
}

var DefaultSqlTypeConvertFunc = func(arg interface{}) string {
	if arg == nil {
		return ""
	}
	var t = reflect.TypeOf(arg)
	if t.String() == Adapter_DateType {
		arg = arg.(time.Time).Format(Adapter_FormateDate)
	} else if t.String() == Adapter_DateType || t.String() == Adapter_StringType {
		var argStr bytes.Buffer
		argStr.WriteString(`'`)
		argStr.WriteString(toString(arg))
		argStr.WriteString(`'`)
		return argStr.String()
	}
	return toString(arg)
}

func toString(value interface{}) string {
	if value == nil {
		return ""
	}
	return fmt.Sprint(value)
}
