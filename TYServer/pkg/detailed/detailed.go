package detailed

import (
	"TYServer/model/message"
)

//ExecFuncs 执行函数群
type ExecFuncs struct {
	registeredFunc map[string]message.FuncType
}

var execFuncs ExecFuncs

//MainDetailed 函数注册
func MainDetailed() {

	if len(execFuncs.registeredFunc) == 0 {
		execFuncs.registeredFunc = make(map[string]message.FuncType)
	}

	execFuncs.AddFuncMap()

}

//AddFuncMap 添加函数到执行列表中
func (execfuncs *ExecFuncs) AddFuncMap() {
	//用户注册函数
	execfuncs.registeredFunc["SYS0001"] = UserRegistered
	//用户登录函数
	execfuncs.registeredFunc["SYS0002"] = UserLogin

}

//GetFuncByCode 获取值
func (execfuncs *ExecFuncs) GetFuncByCode(code string) (funcName message.FuncType) {

	return execfuncs.registeredFunc[code]
}

//FindFuncByCode 根据业务码查找函数
func FindFuncByCode(code string) (funcName message.FuncType) {

	return execFuncs.GetFuncByCode(code)
}
