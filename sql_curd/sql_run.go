package sql_curd

import (
	// "fmt"
	"reflect"
	"strconv"
	"time"

	"http_sql_api/config"
)

//查询单一
// func (m *Models) Select(kjs ...interface{}) (resultsSlice []map[string][]byte,resul []map[string]string, err error) {
func (m *Models) Select(receiveModels ...interface{}) (resultSlice []map[string]string, err error){
// func (m *Models) Select(receiveModels ...interface{})  {
	//var resultslices []map[string][]byte
	

	// sql_str := m.analysis()

	// //db, err := m.DB.Query(sql_str)

	// fmt.Println(sql_str);
	// //defer m.DB.Close()
	// defer m.InitModel()
	//return resultSlice,nil
	var resultslices []map[string][]byte

	defer m.InitModel()

	sql_str := m.analysis()

	query_ary, err := m.DB.Query(sql_str)

	if err != nil {
		return nil,err
	}

	defer query_ary.Close()
	fields, err := query_ary.Columns()
	if err != nil {
		return nil,err
	}

	for query_ary.Next() {
		result := make(map[string][]byte)
		var result_arys []interface{}
		for i := 0; i < len(fields); i++ {
			var fields_ary interface{}
			result_arys = append(result_arys,&fields_ary)
		}
		if err := query_ary.Scan(result_arys...); err != nil {
			return nil,err
		}

		for j, key := range fields {
			res_val := reflect.Indirect(reflect.ValueOf(result_arys[j]))
			if res_val.Interface() == nil {
				continue
			}
			val_type := reflect.TypeOf(res_val.Interface())
			val_cont := reflect.ValueOf(res_val.Interface())
			var str_val string
			switch val_type.Kind() {
				case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					str_val	= strconv.FormatInt(val_cont.Int(),20)
					result[key] = []byte(str_val)
				case reflect.Float32, reflect.Float64:
					str_val = strconv.FormatFloat(val_cont.Float(),'f',-1,64)
					result[key] = []byte(str_val)
				case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					str_val = strconv.FormatUint(val_cont.Uint(),20)
					result[key] = []byte(str_val)
				case reflect.Slice:
					if val_type.Elem().Kind() == reflect.Uint8{
						result[key] = res_val.Interface().([]byte)
					}
				case reflect.String:
					str_val = val_cont.String()
					result[key] = []byte(str_val)
				case reflect.Struct:
					str_val = res_val.Interface().(time.Time).Format("1970-1-1 08:00:00.000")
					result[key] = []byte(str_val)
				case reflect.Bool:
					if val_cont.Bool() {
						result[key] = []byte("1")
					}else{
						result[key] = []byte("0")
					}

			}
		}
		resultslices = append(resultslices,result)
	}
	resultSlice, _ = uintToString(resultslices)
	defer m.DB.Close()
	return resultSlice,nil
}

//uint 转 string
func uintToString(resultsVal []map[string][]byte) (resultsString []map[string]string, err error) {
	for _, v := range resultsVal {
		val := make(map[string]string)
		for kk, vv := range v {
			val[kk] = string(vv)
		}
		resultsString = append(resultsString,val)
	}
	return resultsString,nil
}
//初始化配置
func (m *Models) InitModel() {
	m.TableName = ""
	m.AliasName = ""
	m.PrimaryKey = "id"
	m.Fieldes = "*"
	m.WhereStr=""
	m.WhereInterface = make([]interface{}, 0)
	m.OrderStr = ""
	m.LimitInt = 1
	m.PageSize = 20
	m.GroupStr = ""
	m.HavingStr = ""
	m.JoinStr = ""
	m.WhereFrequency = 0
	m.LibraryName = ""
	m.QuoteIdentifier="`"
	m.ParamIdentifier=config.AppConfig.DataBaseType
}