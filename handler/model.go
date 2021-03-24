package handler

type Result struct {
	Code    int         `json:"code" description:"响应码"`
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

func ResultError(code int, message string) Result {
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

type DevOps struct {
	Type       string `json:"type" binding:"required" description:"类型"`
	ImageGroup string `json:"imageGroup" binding:"required" description:"镜像仓库分组"`
	Image      string `json:"image" binding:"required" description:"镜像名称"`
	GitGroup   string `json:"gitGroup" binding:"required" description:"Git分组"`
	GitName    string `json:"gitName" binding:"required" description:"Git项目名"`
	Project    string `json:"project" binding:"required" description:"项目名称"`
	App        string `json:"app" binding:"required" description:"应用名称"`
	Workload   string `json:"workload" binding:"required" description:"工作负载名称"`
	Desc       string `json:"desc" description:"工作负载描述"`
}
