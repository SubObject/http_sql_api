# Go  Mysql CURD 操作练手

***

## 一、基础配置

### 1、App基础应用。

所有的配置文件都在 conf 目录下

> App配置(appConfig.ini)

~~~
	[App]
	;监听端口
	ListenPort          = :12148
	;App密钥
	AppKey              =   qindong@163.com
	;数据库类型
	DataBaseType        =   mysql
	;是否读写分离
	ReadWriteSeparation =   false
	;主库配置名
	MasterDsn           =   "MysqlConfig"
	;从库名
	SlaveDsn            =   ""
~~~

> Mysql配置(mysqlConfig.ini)

~~~
	[Mysql]
	;数据库地址
	DB_Host         =   127.0.0.1
	;数据库用户名
	DB_User         =   root
	;数据库密码
	DB_Pwds         =   ""
	;数据库端口
	DB_Port         =   3306
	;数据库名称
	DB_Name         =   app_go
	;主从（1：主库；0：从库）
	Master_Slave    =   1
~~~

> Redis配置（redisConfig.ini）

~~~
	[Redis]
	;Redis地址
	Redis_Host =   127.0.0.1
	;Redis端口
	Redis_Port =   6379
	;Redis密码
	Redis_pwds =  
~~~

## 二、数据库

### 1、Mysql数据库模型类定义

> 模型定义

 App配置文件在c'onf 目录下
 
~~~
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
	}
~~~

> 查询构造器
>> 执行所有数据库操作前需要引入库Model

~~~
	var sql_model sql_curd.Models
~~~

	1、查询数据
	2、添加数据
	3、更新数据
	4、删除数据
	5、链式操作

> # 1、查询数据

## 基本查询
	
查询单个数据使用Find()方法
	
~~~
	//第一种方法
	//TableNames指定要查询的表 &databest 为表模型
	list, err :=sql_model.Db().TableNames("system_admin").Find(&databest)
	
	//第一种方法
	//当存在Field的时候Find可不传递任何值
	list, err :=sql_model.Db().TableNames("system_admin").Field("id,username,pwd").Find()
~~~

最终生成的SQL语句可能是：

~~~
	//第一种方法
	SELECT * FROM `system_admin` WHERE  `id` = 1 LIMIT 1
	
	//第一种方法
	SELECT `id`,`username`,`pwd` FROM `system_admin` WHERE  `id` = 1 LIMIT 1
~~~

查询多个数据（数据集）使用select方法：

~~~
	list, err :=sql_model.Db().TableNames("system_admin").Select()
~~~

最终生成的SQL语句可能是：

~~~
	SELECT * FROM `system_admin`
~~~

如果设置了数据表前缀参数的话，可以使用

~~~
	username := "admin"
	Whe_ary := sql_curd.SetMapOut()
	Whe_ary["username"]=sql_curd.Setwhere{"like",username}

	list, err :=sql_model.Db().TableNames("system_admin").Where(Whe_ary).Find(&databest)
	list, err :=sql_model.Db().TableNames("system_admin").Where(Whe_ary).Select()
~~~

最终生成的SQL语句可能是：

~~~
	SELECT * FROM `system_admin` WHERE  `username` LIKE "admin" LIMIT 1
	
	SELECT * FROM `system_admin` WHERE  `username` LIKE "admin"
~~~

#### 在find和select方法之前可以使用所有的链式操作（参考链式操作章节）方法。

> # 2、添加数据

### 添加一条数据

 使用 Db 的 insert 方法向数据库提交数据
~~~
	userModel.UserName = "username"
	userModel.Pwd = outputformat.Md5("123456")
	userModel.Creade = time.Now().Unix()
	userModel.UpDate = time.Now().Unix()
	userModel.FullName = "dullname"
	
	cont, err := sql_model.Db().Insert(&userModel)
~~~

或者使用data方法配合insert使用。

~~~
	userModel.UserName = "username"
	userModel.Pwd = outputformat.Md5("123456")
	userModel.Creade = time.Now().Unix()
	userModel.UpDate = time.Now().Unix()
	userModel.FullName = "fullname"
	
	cont, err := sql_model.Db().TableNames("system_admin").Data(&userModel).Insert()
~~~

~~~
	add := make(map[string]interface{})
	add["username"] = "username"
	add["pwd"] = outputformat.Md5("123456")
	add["creade"] = time.Now().Unix()
	add["update"] = time.Now().Unix()
	add["fullname"] = "username"
	
	cont, err := sql_model.Db.TableNames("system_admin").Insert(add)
	
	cont1, err1 := sql_model.Db.TableNames("system_admin").Data(add).Insert(add)
~~~

> # 3、更新数据

### 更新一条数据

使用 Db 的 UpDate 方法向数据库提交数据

~~~
	Whe_ary := sql_curd.SetMapOut()
	Whe_ary["id"]=sql_curd.Setwhere{"=",id_int64}

	addData := make(map[string]interface{})
	addData["update"] = time.Now().Unix()
	addData["fullname"] = fullname

	editSystemAdmin, err := sql_model.Db().TableNames("system_admin").Where(Whe_ary).UpDate(addData)
~~~

最终生成的SQL语句类似于：

~~~
	UPDATE `system_admin` SET `fullname` = "您修改后的数据" WHERE `id` = 1
~~~

> # 4、删除数据

### 删除一条数据

使用 Db 的 Delete 方法向数据库提交数据

~~~
	id_int64, _ := strconv.ParseInt(id, 10, 64)    
	Whe_ary := sql_curd.SetMapOut()
	Whe_ary["id"]=sql_curd.Setwhere{"=",id_int64}

	delSystemAdmin, err := sql_model.Db().TableNames("system_admin").Where(Whe_ary).Delete()
~~~

最终生成的SQL语句类似于：

~~~
	DELETE FROM `system_admin` WHERE `id` = 1
~~~

> # 5、链式操作

### 数据库提供的链式操作方法，可以有效的提高数据存取的代码清晰度和开发效率，并且支持所有的CURD操作（原生查询不支持链式操作）。


使用也比较简单，假如我们现在要查询一个system_admin表的满足状态为1的前10条记录，并希望按照用户的创建时间排序 ，代码如下：

~~~
	username := "admnin"
	Whe_ary := sql_curd.SetMapOut()
	Whe_ary["username"]=sql_curd.Setwhere{"like",username}

	ord_ary :=	sql_curd.SetMapOut()
	ord_ary["id"]="desc"
	ord_ary["creade"]="desc"
	
	list, err :=sql_model.Db().TableNames("system_admin").Alias("sq").Field("*").Join("count c","sq.id = c.id","left").Where(Whe_ary).GroupBy("creade").OrderBy(ord_ary).Limit(0,20).Select()
~~~

这里的Alias、Field、Join、Where、Order、GroupBy和Limit方法就被称之为链式操作方法，除了Select方法必须放到最后一个外（因为Select方法并不是链式操作方法），链式操作的方法调用顺序没有先后.

最终生成的SQL语句可能是：
~~~
	SELECT FROM * system_admin sq LEFT JOIN count c ON sq.id = c.id WHERE `username` LIKE admnin GROUP BY creade ORDER BY id DESC, creade DESC LIMIT 0,20
~~~

> * Where
> 
>  表达式查询

查询表达式的使用格式：

~~~
	username := "admin"
	Whe_ary := sql_curd.SetMapOut()
	Whe_ary["username"]=sql_curd.Setwhere{"like",username}

	list, err :=sql_model.Db().TableNames("system_admin").Where(Whe_ary).Select()
~~~

最终生成的SQL语句可能是：

~~~
	SELECT * FROM `system_admin` WHERE  `username` LIKE "admin"
~~~

> * TableNames
>  
>  主要用于指定操作的数据表


一般在Model情况下，操作模型的时候系统能够自动识别当前对应的数据表，所以，使用table方法的情况通常是为了：

1、切换操作的数据表；

2、对多表进行操作；

> * Alias
>  
>  设置当前数据表的别名，便于使用其他的连贯操作例如join方法等。

查询表达式的使用格式：

~~~
	list, err :=sql_model.Db().TableNames("system_admin").Alias("sa").Where(Whe_ary).Select()
~~~

最终生成的SQL语句类似于：

~~~
	SELECT * FROM system_admin sa WHERE `username` LIKE "admin"
~~~

> * Field
>  
>  主要作用是标识要返回或者操作的字段，可以用于查询和写入操作。

在查询操作中field方法是使用最频繁的。

~~~
	list, err :=sql_model.Db().TableNames("system_admin").Alias("sa").Field("id,username").Where(Whe_ary).Select()
~~~

最终生成的SQL语句类似于：

~~~
	SELECT `id`,`username` FROM system_admin sa WHERE `username` LIKE "admin"
~~~

> * Limit
>  
>  主要用于指定查询和操作的数量。

例如获取满足要求的10个用户，如下调用即可：

~~~
	list, err :=sql_model.Db().TableNames("system_admin").Alias("sa").Field("id,username").Where(Whe_ary).Limit(0,10).Select()
~~~

最终生成的SQL语句类似于：

~~~
	SELECT `id`,`username` FROM system_admin sa WHERE `username` LIKE "admin" LIMIT 0,10
~~~

> * OrderBy
>  
>  用于对操作的结果排序或者优先级限制。

查询表达式的使用格式：

~~~
	list, err :=sql_model.Db().TableNames("system_admin").Alias("sa").Field("id,username").Where(Whe_ary).Limit(0,10).OrderBy(ord_ary).Select()
~~~

最终生成的SQL语句类似于：

~~~
	SELECT `id`,`username` FROM system_admin sa WHERE `username` LIKE "admin" ORDER BY id DESC, creade DESC LIMIT 0,10
~~~

> * GroupBy
>  
>  通常用于结合合计函数，根据一个或多个列对结果集进行分组 。
>  
>  __方法只有一个参数，并且只能使用字符串__

查询表达式的使用格式：

~~~
	list, err :=sql_model.Db().TableNames("system_admin").Alias("sa").Field("id,username").Where(Whe_ary).Limit(0,10).OrderBy(ord_ary).GroupBy("creade").Select()
~~~
最终生成的SQL语句类似于：

~~~
	SELECT `id`,`username` FROM system_admin sa WHERE `username` LIKE "admin"  GROUP BY creade ORDER BY id DESC, creade DESC LIMIT 0,10
~~~

> * Join
>  
>  用于根据两个或多个表中的列之间的关系，从这些表中查询数据。join通常有下面几种类型，不同类型的join操作会影响返回的数据结果。


查询表达式的使用格式：

~~~
	list, err :=sql_model.Db().TableNames("system_admin").Alias("sq").Field("*").Join("count c","sq.id = c.id","left").Where(Whe_ary).GroupBy("creade").OrderBy(ord_ary).Limit(0,20).Select()
~~~
最终生成的SQL语句类似于：

~~~
	SELECT FROM * system_admin sq LEFT JOIN count c ON sq.id = c.id WHERE `username` LIKE admnin GROUP BY creade ORDER BY id DESC, creade DESC LIMIT 0,20
~~~

> # 6、事务操作

####  使用事务处理的话，需要数据库引擎支持事务处理。比如 MySQL 的 MyISAM 不支持事务处理，需要使用 InnoDB 引擎。 ####

使用方法

~~~
	//开启事务
	sql_model.BeginGo()
	
	//执行语句
	cont, err := sql_model.Db().Insert(&userModel)
	
	if err != nil {
		//回滚
		sql_model.RollbackGo()
	}
	//执行
	sql_model.CommitGo()
~~~

#### <font color="#FF0000">*</font> 注意在事务操作的时候，确保你的数据库连接使用的是同一个。 ####
####  <font color="#FF0000">*</font> 要确保你的数据表引擎为InnoDB，并且开启XA事务支持。####

## 三、表切片

~~~
	type SystemAdmin struct {
		Id			int64		`json:"id" bson:"id"`			//ID
		UserName	string		`json:"username" bson:"username"`	//用户名
		Pwd			string		`json:"pwd" bson:"pwd"`		//密码
		Creade		int64		`json:"creade" bson:"creade"`		//创建时间
		UpDate		int64		`json:"update" bson:"update"`		//更新时间
		FullName	string		`json:"fullname" bson:"fullname"`	//姓名
	}
~~~
