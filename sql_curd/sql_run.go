package sql_curd

import (
	"fmt"
	// "reflect"
	// "strconv"
	// "time"

	"http_sql_api/config"
)

//查询单一
// func (m *Models) Find(kjs ...interface{}) (resultsSlice []map[string][]byte,resul []map[string]string, err error) {
func (m *Models) Find(kjs ...interface{})  {
	//var resultslices []map[string][]byte
	

	sql_str := m.analysis()

	//db, err := m.DB.Query(sql_str)

	fmt.Println(sql_str);
	//defer m.DB.Close()
	defer m.InitModel()
	//return resultSlice,nil
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