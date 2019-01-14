package sql_curd

import (
	"strings"
	"fmt"

	"http_sql_api/config"

	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

//连接类型
func (m *Models) Db(connect interface{}) *Models{
	m.PrimaryKey="id"
	m.Fieldes="*"
	m.WhereFrequency=0
	m.QuoteIdentifier="`"
	switch connect := connect.(type) {
	case string:
		m.LibraryName=connect
	default:
		m.LibraryName=""
	}
	m.ParamIdentifier = config.AppConfig.DataBaseType
	return m
}
//指定数据库连接
func (m *Models) AppointDataBase(libraryName string) *Models {
	if libraryName != "" {
		m.LibraryName=libraryName
	}
	return m
}
//表名称
func (m *Models) TableNames(name string) *Models{
	m.TableName=fmt.Sprintf("%v%v%v",m.QuoteIdentifier,name,m.QuoteIdentifier)
	return m
}
//表别名
func (m *Models) Alias(name interface{}) *Models {
	switch name := name.(type){
	case string:
		m.AliasName=name
	default:
		m.AliasName=""
	}
	return m
}
//字段
func (m *Models) Field(fields string) *Models {
	if fields == "*" {
		// if len(m.AliasName) == 0 {
		// 	m.Fieldes = fields
		// }else{
		// 	m.Fieldes = fmt.Sprintf("%v.%v",m.Alias,fields)
		// }
		m.Fieldes = m.judgeAliasName(fields)
	}else{
		field_ary := strings.Split(fields,",")
		field_str := []string{}
		if len(field_ary[0]) == 0 {
			m.Fieldes = m.judgeAliasName("*")
		}else{
			for _,val := range field_ary {
				field_ary_dian := strings.Split(val,".")
				if len(field_ary_dian) != 2{
					field_ary_d := strings.Split(val,"`")
					if len(field_ary_d) == 3 {
						val=m.judgeAliasName(val)
						field_str = append(field_str,val)
					}else{
						for i:=0; i<len(field_ary_d); i++ {
							if field_ary_d[i] != "`" &&  field_ary_d[i] != "" {
								field_ary_str := fmt.Sprintf("`%v`",field_ary_d[i])
								val=m.judgeAliasName(field_ary_str)
								field_str = append(field_str,field_ary_str)
							}
						}
					}
				}else{
					field_ary_zhi := strings.Split(field_ary_dian[1],"`")
					if len(field_ary_zhi) == 3 {
						field_str = append(field_str,val)
					}else{
						for i:=0; i<len(field_ary_zhi); i++ {
							if field_ary_zhi[i] != "`" &&  field_ary_zhi[i] != "" {
								field_ary_str := fmt.Sprintf("`%v`",field_ary_zhi[i])
								if len(field_ary_dian[0]) != 0 {
									field_ary_str=fmt.Sprintf("%v.%v",field_ary_dian[0],field_ary_str)
								}
								field_str = append(field_str,field_ary_str)
							}
						}
					}
				}
			}
			m.Fieldes=strings.Join(field_str,",")
		}
	}
	return m
}

//判断是否有数据别名
func  (m *Models)judgeAliasName(val string) string {
	if len(m.AliasName) != 0 {
		val=fmt.Sprintf("%v.%v",m.Alias,val)
	}
	return val
}
//查询条件
func (m *Models) Where(wheStr interface{}) *Models {
	wheAry := []string{}
	switch wheStr	:= wheStr.(type) {
	case string:
		wheAry = append(wheAry,wheStr)
	case int:
		wheAry = append(wheAry,fmt.Sprintf("%v%v%v = %v",m.QuoteIdentifier,m.PrimaryKey,m.QuoteIdentifier,wheStr))
	case map[string]interface{}:
		for key, val := range wheStr {
			switch value := val.(type) {
				case string:
					wheAry = append(wheAry,fmt.Sprintf("%v%v%v = '%v'", m.QuoteIdentifier, key, m.QuoteIdentifier, value))
				case int:
					wheAry = append(wheAry,fmt.Sprintf("%v%v%v = %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value))
				case int8:
					wheAry = append(wheAry,fmt.Sprintf("%v%v%v = %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value))
				case int16:
					wheAry = append(wheAry,fmt.Sprintf("%v%v%v = %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value))
				case int32:
					wheAry = append(wheAry,fmt.Sprintf("%v%v%v = %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value))
				case int64:
					wheAry = append(wheAry,fmt.Sprintf("%v%v%v = %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value))
				case Setwhere:
					switch valStr := value.Result.(type) {
						case int:
							wheAry = append(wheAry,fmt.Sprintf("%v%v%v %v %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value.Equation,valStr))
						case int8:
							wheAry = append(wheAry,fmt.Sprintf("%v%v%v %v %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value.Equation,valStr))
						case int16:
							wheAry = append(wheAry,fmt.Sprintf("%v%v%v %v %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value.Equation,valStr))
						case int32:
							wheAry = append(wheAry,fmt.Sprintf("%v%v%v %v %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value.Equation,valStr))
						case int64:
							wheAry = append(wheAry,fmt.Sprintf("%v%v%v %v %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value.Equation,valStr))
						default:
							wheAry = append(wheAry,fmt.Sprintf("%v%v%v %v '%v'", m.QuoteIdentifier, key, m.QuoteIdentifier, value.Equation,valStr))
					}
				default:
			}
		}
	default:
	}
	if len(wheAry) > 1{
		m.WhereStr=strings.Join(wheAry," AND ")
	}else {
		if(wheAry[0] != ""){
			m.WhereStr= wheAry[0]
		}
	}
	return m
}
//过滤条件
func (m *Models) Having(havStr interface{}) *Models {
	havAry := []string{}
	switch havStr := havStr.(type) {
		case string:
			havAry = append(havAry,havStr)
		case int:
			havAry = append(havAry,fmt.Sprintf("%v%v%v = %v", m.QuoteIdentifier, m.PrimaryKey, m.QuoteIdentifier, havStr))
		case map[string]interface{}:
			for key, val := range havStr {
				switch value := val.(type) {
					case string:
						havAry = append(havAry,fmt.Sprintf("%v%v%v = '%v'", m.QuoteIdentifier, key, m.QuoteIdentifier, value))
					case int:
						havAry = append(havAry,fmt.Sprintf("%v%v%v = %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value))
					case int8:
						havAry = append(havAry,fmt.Sprintf("%v%v%v = %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value))
					case int16:
						havAry = append(havAry,fmt.Sprintf("%v%v%v = %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value))
					case int32:
						havAry = append(havAry,fmt.Sprintf("%v%v%v = %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value))
					case int64:
						havAry = append(havAry,fmt.Sprintf("%v%v%v = %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value))
					case Setwhere:
						switch valStr := value.Result.(type) {
							case int:
								havAry = append(havAry,fmt.Sprintf("%v%v%v %v %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value.Equation,valStr))
							case int8:
								havAry = append(havAry,fmt.Sprintf("%v%v%v %v %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value.Equation,valStr))
							case int16:
								havAry = append(havAry,fmt.Sprintf("%v%v%v %v %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value.Equation,valStr))
							case int32:
								havAry = append(havAry,fmt.Sprintf("%v%v%v %v %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value.Equation,valStr))
							case int64:
								havAry = append(havAry,fmt.Sprintf("%v%v%v %v %v", m.QuoteIdentifier, key, m.QuoteIdentifier, value.Equation,valStr))
							default:
								havAry = append(havAry,fmt.Sprintf("%v%v%v %v '%v'", m.QuoteIdentifier, key, m.QuoteIdentifier, value.Equation,valStr))
						}
					default:
				}
			}
		default:
	}
	if len(havAry) > 1{
		m.HavingStr=strings.Join(havAry," AND ")
	}else {
		if(havAry[0] != ""){
			m.HavingStr= havAry[0]
		}
	}
	return m
}
//查询排序
func (m *Models) OrderBy(ordStr interface{}) *Models{
	ordAry := []string{}
	switch ordStr := ordStr.(type) {
	case string:
		switch ordStr{
		case "desc":
			ordAry = append(ordAry,fmt.Sprintf(" %v DESC",m.PrimaryKey))
		case "DESC":
			ordAry = append(ordAry,fmt.Sprintf(" %v DESC",m.PrimaryKey))
		case "asc":
			ordAry = append(ordAry,fmt.Sprintf(" %v ASC",m.PrimaryKey))
		case "ASC":
			ordAry = append(ordAry,fmt.Sprintf(" %v ASC",m.PrimaryKey))
		default:
			ordAry = append(ordAry,ordStr)
		}
	case int:
		if ordStr > 0 {
			ordAry = append(ordAry,fmt.Sprintf(" %v DESC",m.PrimaryKey))
		}else{
			ordAry = append(ordAry,fmt.Sprintf(" %v ASC",m.PrimaryKey))
		}
	case map[string]interface{}:
		for key, val := range ordStr{
			switch value := val.(type) {
				case string:
					ordAry = append(ordAry,fmt.Sprintf(" %v %v",key,value))
				case int:
					if value > 0 {
						ordAry = append(ordAry,fmt.Sprintf(" %v DESC",key))
					}else{
						ordAry = append(ordAry,fmt.Sprintf(" %v ASC",key))
					}
				default:
			}
		}
	default:
	}
	if len(ordAry) > 1{
		m.OrderStr = " ORDER BY" + strings.Join(ordAry,",")
	}else {
		if(ordAry[0] != ""){
			ord_ary := strings.Split(ordAry[0],"order by")
			ord_ary_big := strings.Split(ordAry[0],"ORDER BY")

			if len(ord_ary) < 2 && len(ord_ary_big) < 2 {
				m.OrderStr = " ORDER BY" + ordAry[0]
			}else{
				if len(ord_ary) > 1{
					for i :=0;i<len(ord_ary);i++ {
						if ord_ary[i] != "" && ord_ary[i] != "order by" {
							m.OrderStr = " ORDER BY" + ord_ary[i]
						}
					}
				}else if  len(ord_ary_big) > 1{
					for i :=0;i<len(ord_ary_big);i++ {
						if ord_ary_big[i] != "" && ord_ary_big[i] != "ORDER BY" {
							m.OrderStr = " ORDER BY" + ord_ary_big[i]
						}
					}
				}
			}

			
		}
	}
	
	return m
}
//归纳分组
func (m *Models) GroupBy(groupby string) *Models {
	m.GroupStr = fmt.Sprintf(" GROUP BY %v", groupby)
	return m
}
//查询多少条
func (m *Models) Limit(star int,size ...int) *Models {
	m.LimitInt = star
	if len(size) > 0 {
		m.PageSize = size[0]
	}
	return m
}

//多表链接
func (m *Models) Join(tablename string,condition string,method string) *Models {
	if m.WhereFrequency == 1 {
		m.JoinStr = fmt.Sprintf("%v JOIN %v ON %v", strings.ToUpper(method),tablename,condition)
	}else{
		m.JoinStr = m.JoinStr + fmt.Sprintf(" %v JOIN %v ON %v ", strings.ToUpper(method),tablename,condition)
	}
	m.WhereFrequency++
	return m
}
//解析拼接sql语句
func (m *Models) analysis() (sqlstr string) {
	switch m.ParamIdentifier {
	case "mysql":
		sqlstr = fmt.Sprintf("SELECT %v FROM %v %v", m.Fieldes, m.TableName,m.AliasName)
		
		if m.JoinStr != "" {
			sqlstr = fmt.Sprintf("%v %v", sqlstr, m.JoinStr)
		}
		if m.WhereStr != "" {
			sqlstr = fmt.Sprintf("%v WHERE %v", sqlstr, m.WhereStr)
		}
		if m.GroupStr != "" {
			sqlstr = fmt.Sprintf("%v %v", sqlstr, m.GroupStr)
		}
		if m.HavingStr != "" {
			sqlstr = fmt.Sprintf("%v HAVING %v", sqlstr, m.HavingStr)
		}
		if m.OrderStr != "" {
			sqlstr = fmt.Sprintf("%v %v", sqlstr, m.OrderStr)
		}
		if m.PageSize > 0 {
			sqlstr = fmt.Sprintf("%v LIMIT %v OFFSET %v", sqlstr, m.LimitInt, m.PageSize)
		} else if m.LimitInt > 0 {
			sqlstr = fmt.Sprintf("%v LIMIT %v", sqlstr, m.LimitInt)
		}
	}
	m.choiceDatabase("q")
	return
}
//选择数据库
func (m *Models) choiceDatabase(runType string) *Models {
	sql_str := ""
	switch runType {
	case "q":
		sql_str =m.judgeSqlDb(runType)
	case "w":
		sql_str =m.judgeSqlDb(runType)
	default:
		sql_str =m.judgeSqlDb(runType)
	}
	db, err := sql.Open("mysql",sql_str)
	if err != nil {
		return m
	}
	m.DB=db
	return m
}

func (m *Models) judgeSqlDb(runType string) (db string) {
	if config.AppConfig.ReadWriteSeparation == false {
		db = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",config.MysqlConfig.DB_User,config.MysqlConfig.DB_Pwds,config.MysqlConfig.DB_Host,config.MysqlConfig.DB_Port,config.MysqlConfig.DB_Name)
	}
	return db
}