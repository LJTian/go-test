package peripheral

import (
	"TYServer/model/message"
	"TYServer/model/userInfo"
	"TYServer/pkg/detailed"
	. "TYServer/utils/log"
	"encoding/json"
	"fmt"
	_ "time"
)

var ReturnMsg chan *message.PackInfo

//MainPeripheral 外围解包
func MainPeripheral(inMsg *[]byte, key string) (outMsg *[]byte, err error) {

	var messageOb message.PackInfo
	TlogPrintln(LOG_DEBUG, "inMsg is ", string(*inMsg))
	err = json.Unmarshal(*inMsg, &messageOb)
	if err != nil {
		TlogPrintln(LOG_ERROR, "json.Unmarshal err is ", err)
		return

	}

	v, _ := userInfo.OnlineUser.OnlineUser.Load(key)

	userInfo.OnlineUser.UpdateNodeStatus(key, v.(userInfo.UserInfo).UUID,
		messageOb.Sender, v.(userInfo.UserInfo).IP,
		messageOb.Sender, userInfo.ONLINE)

	// TlogPrintln(LOG_DEBUG, "打印在线列表:======================================================")
	// for _, v := range userInfo.OnlineUser.OnlineUser {
	// 	TlogPrintf(LOG_DEBUG, "uuid = [%s] , ClientID = [%s] , ip = [%s] UserName = [%s] , 在线状态[%d]\n", v.UUID, v.ClientID, v.IP, v.UserName, v.Status)
	// }
	// TlogPrintln(LOG_DEBUG, "======================================================")

	TlogPrintf(LOG_DEBUG, "外围解包—— 类别码 : [%v]\t 包类型:[%v]\t 报文长度:[%v]\t 报文:[%v]\t 发送方:[%v]\t 接受方:[%v]\n",
		messageOb.Code, messageOb.PackType, messageOb.Lenth, string(messageOb.Data), messageOb.Sender, messageOb.Receer)
	TlogPrintln(LOG_DEBUG, "业务码为[%s]", messageOb.Code)

	//根据业务码获取执行函数(根据code 获取执行函数)
	funcName := detailed.FindFuncByCode(messageOb.Code)
	if funcName == nil {
		TlogPrintln(LOG_ERROR, "此业务码[%s] 未识别，请确认是否进行了注册", messageOb.Code)
	}
	outData, err := funcName(&messageOb.Data)
	if err != nil {
		TlogPrintln(LOG_ERROR, "执行函数失败[%v]", funcName)
	}

	TlogPrintln(LOG_DEBUG, "返回报文outMsg [%s]", string(outData.Data))

	//将数据发送到消息管道
	ReturnMsg <- outData

	return
}

//sendMsgFunc 发送携程
func SendMsgFunc() {

	for {
		//从管道获取数据
		select {
		case message := <-ReturnMsg:
			SendMsg2TcpClient(message)
			//fmt.Println(message)
			//default:
			//time.Sleep(100)
		}

	}

}

//SendMsg2TcpClient
func SendMsg2TcpClient(message *message.PackInfo) {

	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s \n", e)
		}
	}()

	key, err := userInfo.OnlineUser.FindNodeListByClitneID(message.Receer)
	if err != nil {
		TlogPrintln(LOG_ERROR, "未找到对应客户端Id ", message.Receer)
		return
	}
	TlogPrintf(LOG_DEBUG, "应客户端Id[%s] 对应的key为[%s] ", message.Receer, key)

	v, _ := userInfo.OnlineUser.OnlineUser.Load(key)

	//fmt.Println(v)
	TlogPrintf(LOG_DEBUG, "回复的数据为[%s] ", string(message.Data))
	_, err = v.(userInfo.UserInfo).Conn.Write(message.Data)
	if err != nil {
		fmt.Println(err)
	}

}
