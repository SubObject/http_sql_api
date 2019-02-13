package sql_curd

import (
	"strings"
	"fmt"
)

//查询条件
func (m *Models) Wheres(wheStr interface{},exp ...interface{}) *Models {
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
	if m.WhereStr == "" {
		if len(wheAry) > 1{
			m.WhereStr=strings.Join(wheAry," AND ")
		}else {
			if(wheAry[0] != ""){
				m.WhereStr= wheAry[0]
			}
		}
	}else{
		wheStr := strings.Split(m.WhereStr,"(")
		var wheStrVal string
		if len(wheAry) > 1{
			wheStrVal=strings.Join(wheAry," AND ")
		}else {
			if(wheAry[0] != ""){
				wheStrVal = wheAry[0]
			}
		}
		if len(wheStr) > 1 {
			m.WhereStr = fmt.Sprintf("( %v ) AND ( %v )",m.WhereStr ,wheStrVal)
		}else{
			m.WhereStr = fmt.Sprintf("%v AND %v",m.WhereStr ,wheStrVal)
		}
	}
	
	return m
}