package handler

type Result struct {
	Code    int32       `json:"code" description:"响应码"`
	Status  bool        `json:"status" description:"状态"`
	Message string      `json:"message" description:"响应消息"`
	Data    interface{} `json:"data" description:"响应数据"`
}

func ResultNew() Result {
	return Result{
		Code:   200,
		Status: true,
	}
}

func ResultError(code int32, message string) Result {
	return Result{
		Code:    code,
		Status:  false,
		Message: message,
	}
}

func ResultSucc(data interface{}) Result {
	result := ResultNew()
	result.Data = data
	return result
}

type KeyValue struct {
	Key  string `yaml:"key" json:"key" description:"键"`
	Name string `yaml:"name" json:"name" description:"值"`
}

type Group struct {
	Key      string     `yaml:"key" json:"key" description:"键"`
	Name     string     `yaml:"name" json:"name" description:"值"`
	Projects []KeyValue `yaml:"projects" json:"projects" description:"项目列表"`
}

type AppConfig struct {
	Groups []Group `yaml:"groups" json:"groups" description:"分组列表"`
}

type DevOpsType string

const (
	Java   DevOpsType = "Java"
	NodeJs DevOpsType = "nodejs"
)

type DevOps struct {
	Type       DevOpsType `json:"type" binding:"required" description:"类型"`
	ImageGroup string     `json:"imageGroup" binding:"required" description:"镜像仓库分组"`
	Image      string     `json:"image" binding:"required" description:"镜像名称"`
	GitGroup   string     `json:"gitGroup" binding:"required" description:"Git分组"`
	GitName    string     `json:"gitName" binding:"required" description:"Git项目名"`
	Project    string     `json:"project" binding:"required" description:"项目名称"`
	App        string     `json:"app" binding:"required" description:"应用名称"`
	Workload   string     `json:"workload" binding:"required" description:"工作负载名称"`
	Desc       string     `json:"desc" description:"工作负载描述"`
}

func DevOpsNew() DevOps {
	return DevOps{}
}

type DbType int32

const (
	MySQL      DbType = 0
	PostgreSQL DbType = 1
	Oracal     DbType = 2
	SqlServer  DbType = 3
	Sqlite     DbType = 4
)

type Database struct {
	Type   DbType `json:"type" description:"数据库类型"`
	DbHost string `json:"host" description:"数据库主机地址"`
	DbPort int32  `json:"port" description:"数据库端口"`
	DbName string `json:"name" description:"数据库名称"`
	DbUser string `json:"user" description:"数据库用户"`
	DbPwd  string `json:"pwd" description:"数据库密码"`
}

func DatabaseNew() Database {
	return Database{
		Type:   MySQL,
		DbHost: "localhost",
		DbPort: 3306,
		DbName: "",
		DbUser: "root",
		DbPwd:  "",
	}
}

type Springboot struct {
	GroupId   string   `json:"group" binding:"required" description:"分组ID"`
	Project   string   `json:"project" binding:"required" description:"项目名"`
	Version   string   `json:"version" description:"版本号"`
	Framework string   `json:"framework" binding:"required" description:"框架版本"`
	Database  Database `json:"database" description:"数据库配置"`
	Devops    DevOps   `json:"devops" description:"DevOps配置"`
}

func SpringbootNew() Springboot {
	return Springboot{
		GroupId:   "com.yunzhicloud",
		Project:   "",
		Version:   "1.0-SNAPSHOT",
		Framework: "1.1.9",
		Database:  DatabaseNew(),
		Devops:    DevOpsNew(),
	}
}
