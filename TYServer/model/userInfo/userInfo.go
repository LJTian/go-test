package userInfo

import (
	"errors"
	_ "fmt"
	"net"
	"sync"
)

var OnlineUser UserInfoList

//UserInfo 用户结构体
type UserInfo struct {
	UUID     string   //唯一id
	ClientID string   //客户端id
	IP       string   //客户端ip
	UserName string   //用户名称
	Conn     net.Conn //链接套接字
	Status   int      //状态
}

//status 状态标识
const (
	ONLINE   = 1 //在线
	NOONLINE = 0 //不在线
)

//UserInfoList 在线结构体列表
type UserInfoList struct {
	OnlineUser sync.Map
}

//NewMessInfo 创建一个消息结构体
func NewUserInfo(uuid string, clientID string, ip string, userName string, conn net.Conn, status int) *UserInfo {
	userInfo := &UserInfo{
		UUID:     uuid,
		ClientID: clientID,
		IP:       ip,
		UserName: userName,
		Conn:     conn,
		Status:   status,
	}
	return userInfo
}

//AddNode2List 添加节点到列表中,列表不存在将自动make
func (userInfoList *UserInfoList) AddNode2List(userinfo *UserInfo) (err error) {

	userInfoList.OnlineUser.Store(userinfo.IP, *userinfo)
	return
}

//DeleteNode2List 根据key值删除对应节点(默认不使用,使用状态改变来代替此操作)
func (userInfoList *UserInfoList) DeleteNode2List(key string) (err error) {

	userInfoList.OnlineUser.Delete(key)
	return
}

//UpdateNodeStatus 更改数据信息(整体信息修改,需要先查出原信息后,再进行修改)
func (userInfoList *UserInfoList) UpdateNodeStatus(key string, uuid string, clientID string, ip string, userName string, status int) (err error) {

	v, _ := userInfoList.OnlineUser.Load(key)
	newOb := NewUserInfo(uuid, clientID, ip, userName, v.(UserInfo).Conn, 1)
	userInfoList.OnlineUser.Store(key, *newOb)
	return
}

//FindNodeListByUuid 根据uuid返回节点的key
func (userInfoList *UserInfoList) FindNodeListByUuid(uuid string) (key string, err error) {

	userInfoList.OnlineUser.Range(func(k, v interface{}) bool {
		if v.(UserInfo).UUID == uuid && v.(UserInfo).Status == 1 {
			key = k.(string)
			return true
		}
		err = errors.New("没有找到key")
		return false
	})
	return
}

//FindNodeListByClitneID 根据clitneID返回节点的key
func (userInfoList *UserInfoList) FindNodeListByClitneID(clitneID string) (key string, err error) {

	userInfoList.OnlineUser.Range(func(k, v interface{}) bool {
		if v.(UserInfo).ClientID == clitneID && v.(UserInfo).Status == 1 {
			key = k.(string)
			return true
		}
		err = errors.New("没有找到key")
		return false
	})
	//fmt.Println("11")
	return
}

//FindNodeListByUserName 根据UserName返回节点的key
func (userInfoList *UserInfoList) FindNodeListByUserName(userName string) (key string, err error) {

	userInfoList.OnlineUser.Range(func(k, v interface{}) bool {
		if v.(UserInfo).UserName == userName && v.(UserInfo).Status == 1 {
			key = k.(string)
			return true
		}
		err = errors.New("没有找到key")
		return false
	})
	return
}
