package sql_curd

import (
	"database/sql"
)

type Models struct {
	DB						*sql.DB						//数据库模型
	TableName				string						//数据库名称
	AliasName				string						//表别名
	PrimaryKey				string						//主键
	Fieldes					string						//查询字段
	WhereStr				string						//查询条件
	WhereInterface			interface{}					//附属条件
	OrderStr				string						//排序
	LimitInt				int							//查询条数
	PageSize				int							//每页显示多少条
	GroupStr				string						//分类归组
	HavingStr				string						//为行分组或聚合组指定过滤条件
	JoinStr					string						//多表查询
	WhereFrequency			int							//查询拼接次数
	LibraryName				string						//指定使用数据库配置
	QuoteIdentifier			string						//识别符
	ParamIdentifier			string						//数据库顺式链接识别符
	DataKey					string						//数据值
	DataVal					[]interface{}					//数据值
	ParamIteration			int							//迭代使用
	SqlLink					string						//sql语句
	Id 						int64						
	OpenStatus				int
	Affair					*sql.Tx
}

type Setwhere struct {
	Equation	string
	Result		interface{}
}

func SetMapOut() (data map[string]interface{}) {
	data= make(map[string]interface{}) //必可不少，分配内存
	return data
}