package message

//PackInfo 包信息
type PackInfo struct {
	UUID     string `json:"uuid"`     //唯一标识码
	Code     string `json:"code"`     //区别码
	PackType string `json:"packtype"` //报文类型
	Data     []byte `json:"data"`     //数据
	Lenth    int    `json:"lenth"`    //长度
	Sender   string `json:"sender"`   //发送方
	Receer   string `json:"receer"`   //接收方
	Time     string `json:"time"`     //发送时间
}

//FuncType 注册函数声明
type FuncType func(inMsg *[]byte) (*PackInfo, error)

//NewPackInfo 初始化PackInfo
func NewPackInfo(uuid string, code string, packType string, data *[]byte, lenth int, sender string, receer string, time string) *PackInfo {
	packinfo := &PackInfo{
		UUID:     uuid,
		Code:     code,
		PackType: packType,
		Data:     *data,
		Lenth:    lenth,
		Sender:   sender,
		Receer:   receer,
		Time:     time,
	}
	return packinfo
}
