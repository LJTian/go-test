//用户相关的执行函数
package business

import (
	"TYServer/model/message"
)

type UserFunc struct {
	//NULL
}

//Login 登录
func (userfunc *UserFunc) Login(inMsg *[]byte) (outMsg *message.ExecPackInfo, err error) {

	outMsg.Adder = string(*inMsg)
	outMsg.Sender = string(*inMsg)
	outMsg.Data = []byte("田利军~~~~~~~~~~")
	//确切的业务处理
	return
}

//Login Registered
func (userfunc *UserFunc) Registered(inMsg *[]byte) (outMsg *message.ExecPackInfo, err error) {
	//确切的业务处理
	var outMsgob message.ExecPackInfo

	outMsgob.Adder = string(*inMsg)
	outMsgob.Sender = string(*inMsg)
	outMsgob.Data = []byte("田利军~~~~~~~~~~")

	outMsg = &outMsgob
	return
}
