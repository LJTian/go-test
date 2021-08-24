package detailed

import (
	"TYServer/business"
	"TYServer/model/message"
	. "TYServer/utils/log"
	"TYServer/utils/tools"
	"TYServer/utils/uuid"
	"errors"
	_ "fmt"
)

var UserFunc business.UserFunc

//UserLogin 用户登录函数
func UserLogin(inMsg *[]byte) (outMsg *message.PackInfo, err error) {

	//fmt.Println(string(*inMsg))

	//fmt.Println(string(*inMsg))

	outMsg, err = TYExecFunc(UserFunc.Login, inMsg, "SYSRE0002")
	if err != nil {
		TlogPrintln(LOG_ERROR, "UserRegistered err:", err)
	}

	return
}

//UserRegistered 用户注册函数
func UserRegistered(inMsg *[]byte) (outMsg *message.PackInfo, err error) {

	//fmt.Println(string(*inMsg))

	outMsg, err = TYExecFunc(UserFunc.Registered, inMsg, "SYSRE0001")
	if err != nil {
		TlogPrintln(LOG_ERROR, "UserRegistered err:", err)
	}

	return
}

//TYExecFunc 通用执行函数，由于代码相同部分过多，所以进行函数整合
func TYExecFunc(funcName func(*[]byte) (*message.ExecPackInfo, error), inMsg *[]byte, code string) (outMsg *message.PackInfo, err error) {

	msg, err := funcName(inMsg)
	if err != nil {
		TlogPrintln(LOG_ERROR, "TYExecFunc err:", err)
	}

	//获取需要发送的地址
	adder, err := GetAdderBySenderOrAdder(msg.Adder, msg.Sender)
	if err != nil {
		return
	}

	var messageOb message.PackInfo

	messageOb.Code = code
	messageOb.Lenth = len(msg.Data)
	messageOb.Data = msg.Data
	messageOb.PackType = "Json"
	messageOb.Time = tools.LocalTime()
	messageOb.UUID = uuid.CreateUUID()
	messageOb.Receer = adder

	outMsg = &messageOb

	return
}

//GetAdderBySenderOrAdder 获取发送方地址
func GetAdderBySenderOrAdder(adder string, sender string) (reAdder string, err error) {

	//还需要修改
	if adder != "" {
		return adder, nil
	} else if sender != "" {

		return sender, nil
	}

	err = errors.New("GetAdderBySenderOrAdder err")
	return "", err

}
