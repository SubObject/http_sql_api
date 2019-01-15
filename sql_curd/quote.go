package sql_curd

import (
	"errors"
	"reflect"
	//"strconv"
	"strings"
	//"time"
)
//扫描结构以映射
func scanStructIntoMap(object interface{}) (map[string]interface{}, error) {
	dataStruct := reflect.Indirect(reflect.ValueOf(object))
	if dataStruct.Kind() != reflect.Struct {
		return nil, errors.New("数据引用错误，该数据类型应为指向数据结构的指针！")
	}
	dataStructType := dataStruct.Type()

	mapdata := make(map[string]interface{})

	for i := 0; i < dataStructType.NumField(); i++ {
		field := dataStructType.Field(i)
		fieldval := dataStruct.Field(i)
		fieldname := field.Name
		fieldtag := field.Tag
		sqlTag := fieldtag.Get("json")
		sqlTagstr := strings.Split(sqlTag, ":")

		var mapKey string

		setStatus := false

		if fieldtag.Get("sql_curd") == "-" || sqlTag == "-" || reflect.ValueOf(fieldtag).String() == "-" {
			continue
		} else if len(sqlTag) > 0 {
			if sqlTagstr[0] == "-" {
				continue
			}
			mapKey = sqlTagstr[0]
		} else {
			mapKey = fieldname
		}

		if len(sqlTagstr) > 1 {
			if stringArrayContains("setStatus", sqlTagstr[1:]) {
				setStatus = true
			}
		}

		if setStatus {
			map2, err2 := scanStructIntoMap(fieldval.Interface())
			if err2 != nil {
				return mapdata, err2
			}
			for k, v := range map2 {
				mapdata[k] = v
			}
		} else {
			value := dataStruct.FieldByName(fieldname).Interface()
			mapdata[mapKey] = value
		}
	}

	return mapdata,nil
}
//字符串数组
func stringArrayContains(needle string, haystack []string) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}
	return false
}
//获取数据表名称
func getTableName(tables interface{}) string {
	val := reflect.TypeOf(tables)
	if val.Kind() == reflect.String {
		tableing, _ := tables.(string)
		return casedName(tableing)
	}
	tablename := scanTableName(tables)
	if len(tablename) > 0 {
		return tablename
	}
	return getTableName(StructName(tables))
}
//model明珠
func casedName(name string) string {
	namestr := make([]rune, 0)
	firstTime := true

	for _, chr := range name {
		if isUpper := 'A' <= chr && chr <= 'Z'; isUpper {
			if firstTime == true {
				firstTime = false
			} else {
				namestr = append(namestr, '_')
			}
			chr -= ('A' - 'a')
		}
		namestr = append(namestr, chr)
	}

	return string(namestr)
}
//扫描模型名字
func scanTableName(str interface{}) string {
	if reflect.TypeOf(reflect.Indirect(reflect.ValueOf(str)).Interface()).Kind() == reflect.Slice {
		sliceValue := reflect.Indirect(reflect.ValueOf(str))
		sliceElementType := sliceValue.Type().Elem()
		for i := 0; i < sliceElementType.NumField(); i++ {
			tabelTag := sliceElementType.Field(i).Tag
			if len(tabelTag.Get("models")) > 0 {
				return tabelTag.Get("models")
			}
		}
	} else {
		tt := reflect.TypeOf(reflect.Indirect(reflect.ValueOf(str)).Interface())
		for i := 0; i < tt.NumField(); i++ {
			tabelTag := tt.Field(i).Tag
			if len(tabelTag.Get("models")) > 0 {
				return tabelTag.Get("models")
			}
		}
	}
	return ""

}
func StructName(str interface{}) string {
	val := reflect.TypeOf(str)
	for val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val.Name()
}