package qqlt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Msg struct {
	Type    int32  `json:"type"`
	GroupId string `json:"groupId"`
	QQId    string `json:"qqId"`
	Msg     string `json:"msg"`
}

type BotApiClient struct {
	Client  *http.Client
	BaseApi string
}

// NewBotApiClient 获取一个机器人操作客户端
func NewBotApiClient() (*BotApiClient) {
	bot := &BotApiClient{
		Client:  &http.Client{},
		BaseApi: "http://127.0.0.1:36524/api/v1/QQLight",
	}

	return bot
}

func (bot *BotApiClient) SetBaseApi(baseApi string) {
	bot.BaseApi = baseApi
}

// Post send a post request to mahua plugin with yourself api and data
func (bot *BotApiClient) Post(api string, data interface{}) (string, error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// contentType必须为"application/json", 否则mahua插件不处理
	resp, err := bot.Client.Post(api, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), nil
}

// SendPraise 抖动好友窗口
func (bot *BotApiClient) SendShake(qq string) {
	_, _ = bot.Post(bot.BaseApi+Api_SendShake, Params{"qqid": qq})
}

// GetBkn 取得网页操作时需要用到的bkn/Gtk
func (bot *BotApiClient) GetBkn(cookies string) string {
	res, _ := bot.Post(bot.BaseApi+Api_GetBkn, map[string]string{"cookies": cookies})
	return res
}

// QuitGroup 退群
func (bot *BotApiClient) QuitGroup(groupId string) {
	_, _ = bot.Post(bot.BaseApi+Api_QuitGroup, Params{"groupID": groupId})
}

// QuitGroup 退出讨论组
func (bot *BotApiClient) QuitDiscussGroup(groupId string) {
	_, _ = bot.Post(bot.BaseApi+Api_QuitDiscussGroup, Params{"groupID": groupId})
}

// SetDiscussName 修改讨论组名称
func (bot *BotApiClient) SetDiscussName(discussGroupID, name string) {
	_, _ = bot.Post(bot.BaseApi+Api_SetDiscussName, Params{"discussGroupID": discussGroupID, "name": name})
}

// SetGroupName 修改群名称 测试通过 无返回值
func (bot *BotApiClient) SetGroupName(groupId, name string) {
	_, _ = bot.Post(bot.BaseApi+Api_SetGroupName, Params{"groupID": groupId, "name": name})
}

// BanSomeBody 禁言某群某人 time为禁言时间, 单位: 秒, 0为解除禁言
func (bot *BotApiClient) BanSomeBody(groupId, qq string, time int32) {
	_, _ = bot.Post(bot.BaseApi+Api_BanSomeBody, Params{"groupID": groupId, "qq": qq, "time": time})
}

// SetAnony 修改群匿名聊天权限
func (bot *BotApiClient) SetAnony(groupId string, isAnony bool) {
	_, _ = bot.Post(bot.BaseApi+Api_SetAnony, Params{"groupID": groupId, "isSetAnony": isAnony})
}

// RemoveMember 群踢人 isBan不再接受此人加群
func (bot *BotApiClient) RemoveMember(groupId, qq string, isBan bool) {
	_, _ = bot.Post(bot.BaseApi+Api_RemoveMember, Params{"groupID": groupId, "qq": qq, "isBan": isBan})
}

// RemoveGroup 解散群
func (bot *BotApiClient) RemoveGroup(groupId string) {
	_, _ = bot.Post(bot.BaseApi+Api_RemoveGroup, Params{"groupID": groupId})
}

// GetQzoneToken 取得QQ空间Token
func (bot *BotApiClient) GetQzoneToken() string {
	res, _ := bot.Post(bot.BaseApi+Api_GetQzoneToken, nil)
	return res
}

// GetQzoneCookies 取得QQ空间Cookies
func (bot *BotApiClient) GetQzoneCookies() string {
	res, _ := bot.Post(bot.BaseApi+Api_GetQzoneCookies, nil)
	return res
}

// SendTaotao 发表说说
func (bot *BotApiClient) SendTaotao(text string) {
	_, _ = bot.Post(bot.BaseApi+Api_SendTaotao, Params{"str": text})
}

// GetLoginQQ 取得所登录的qq
func (bot *BotApiClient) GetLoginQQ() string {
	res, _ := bot.Post(bot.BaseApi+Api_GetLoginQQ, nil)
	return res
}

// GetPath 框架为插件所创建的插件目录
func (bot *BotApiClient) GetPath() string {
	res, _ := bot.Post(bot.BaseApi+Api_GetPath, nil)
	return res
}

// PluginError 置插件错误管理 调用会导致出错
// Deprecated: 调用会导致qqlight异常退出
func (bot *BotApiClient) PluginError(code, text string) {
	res, _ := bot.Post(bot.BaseApi+Api_PluginError, Params{"code": code, "str": text})
	fmt.Println(res)
}

// SendMsg 机器人发送消息，返回值为该条消息的ID 实际上无返回值
// @param msgType 消息类型 1.好友消息 2.群消息 3.群临时消息 4.讨论组消息 5.讨论组临时消息 6.QQ临时消息
// @param qqId 消息类型为2，4时可留空 即空字符串""
func (bot *BotApiClient) SendMsg(msgType int, groupId, qqId, msg string) {
	_, _ = bot.Post(bot.BaseApi+Api_SendMsg, Params{"type": msgType, "groupID": groupId, "qqid": qqId, "msg": msg})
}

// SendPrivateMsg 发送私聊消息
func (bot *BotApiClient) SendPrivateMsg(qq, msg string) {
	_, _ = bot.Post(bot.BaseApi+Api_SendMsg, Params{"type": MT_Friend, "groupID": "", "qqid": qq, "msg": msg})
}

// SendGroupMsg 发送群消息
func (bot *BotApiClient) SendGroupMsg(groupId, msg string) {
	_, _ = bot.Post(bot.BaseApi+Api_SendMsg, Params{"type": MT_Group, "groupID": groupId, "qqid": "", "msg": msg})
}

// SetGroupAdd 处理加群信息
// @param handleType 处理类型 1.同意 2.拒绝 3.忽略
// @param reason 原因, 仅拒绝时有效, 其他情况下
func (bot *BotApiClient) SetGroupAdd(groupId, qqId, seq string, handleType int32, arg ...string) {
	// 如果拒绝加群 忘记加原因, 默认为空
	if len(arg) == 0 && handleType == 2 {
		arg = append(arg, "")
	}
	_, _ = bot.Post(bot.BaseApi+Api_SetGroupAdd, Params{"groupID": groupId, "qqid": qqId, "seq": seq, "type": handleType, "reason": arg[0]})
}

// SetFriendAdd 处理加群信息
func (bot *BotApiClient) SetFriendAdd(qqId string, handleType int32, arg ...string) {
	if len(arg) == 0 && handleType == 2 {
		arg = append(arg, "")
	}
	_, _ = bot.Post(bot.BaseApi+Api_SetFriendAdd, Params{"qqid": qqId, "type": handleType, "reason": arg[0]})
}

// SendLog 向框架推送一条日志
// 颜色为可变参数, 默认为0 即黑色
func (bot *BotApiClient) SendLog(logType, msg string, color ...int32) {
	if len(color) == 0 {
		color = append(color, 0)
	}
	_, _ = bot.Post(bot.BaseApi+Api_SendLog, Params{"type": logType, "msg": msg, "fontColor": color[0]})
}

// SetFriendName 修改好友备注
func (bot *BotApiClient) SetFriendName(qqId, name string) {
	_, _ = bot.Post(bot.BaseApi+Api_SetFriendName, Params{"qqid": qqId, "name": name})
}

// DeleteFriend 删除好友
func (bot *BotApiClient) DeleteFriend(qqId string) {
	_, _ = bot.Post(bot.BaseApi+Api_DeleteFriend, Params{"qqid": qqId})
}

// AddGroup 主动添加群
func (bot *BotApiClient) AddGroup(groupID, text string) {
	_, _ = bot.Post(bot.BaseApi+Api_AddGroup, Params{"groupID": groupID, "附加信息": text})
}

// AddFriend 主动添加好友
func (bot *BotApiClient) AddFriend(qqId, text string) {
	_, _ = bot.Post(bot.BaseApi+Api_AddFriend, Params{"qqid": qqId, "info": text})
}

// GetJsonMusic 返回卡片点歌JSON代码
func (bot *BotApiClient) GetJsonMusic(songId string) string {
	res, _ := bot.Post(bot.BaseApi+Api_GetJsonMusic, Params{"songID": songId})
	return res
}

// GetGroupCard 取得群内成员的名片
func (bot *BotApiClient) GetGroupCard(groupId string, qqId string) string {
	res, _ := bot.Post(bot.BaseApi+Api_GetGroupCard, Params{"groupID": groupId, "qqid": qqId})
	return res
}

// GetNick 取指定qq名称
func (bot *BotApiClient) GetNick(qq string) string {
	res, _ := bot.Post(bot.BaseApi+Api_GetNick, Params{"qqid": qq})
	return res
}

// GetGroupName 获取群名
func (bot *BotApiClient) GetGroupName(groupId string) string {
	res, _ := bot.Post(bot.BaseApi+Api_GetGroupName, Params{"groupID": groupId})
	return res
}

// GetSoftVersion 获取软件版本
func (bot *BotApiClient) GetSoftVersion() string {
	res, _ := bot.Post(bot.BaseApi+Api_GetSoftVersion, nil)
	return res
}

// GetGroupIntroduce 获取群介绍
func (bot *BotApiClient) GetGroupIntroduce() string {
	res, _ := bot.Post(bot.BaseApi+Api_GetGroupIntroduce, nil)
	return res
}

// GetGroupOwner 获取群主
func (bot *BotApiClient) GetGroupOwner(groupId string) string {
	res, _ := bot.Post(bot.BaseApi+Api_GetGroupOwner, Params{"groupID": groupId})
	return res
}

// SetGroupCard 修改群成员的名片
func (bot *BotApiClient) SetGroupCard(groupId, qq, newCard string) string {
	res, _ := bot.Post(bot.BaseApi+Api_SetGroupCard, Params{
		"groupID": groupId,
		"qqid":    qq,
		"newCard": newCard,
	})
	return res
}

// GetPraiseNum 取得某个QQ的名片赞数量
// Deprecated: 导致qqlight框架闪退
func (bot *BotApiClient) GetPraiseNum(qq string) string {
	res, _ := bot.Post(bot.BaseApi+Api_GetPraiseNum, Params{"qqid": qq})
	return res
}

// GetQQLevel 取得某个QQ的等级
// 通过事件上报返回
func (bot *BotApiClient) GetQQLevel(qq string) {
	_, _ = bot.Post(bot.BaseApi+Api_GetQQLevel, Params{"qqid": qq})
}

// SetNick 修改昵称
func (bot *BotApiClient) SetNick(newNick string) {
	_, _ = bot.Post(bot.BaseApi+Api_SetNick, Params{"newNick": newNick})
}

// GetQQAge 获取q龄
// Deprecated: 无效
func (bot *BotApiClient) GetQQAge(qq string) string {
	res, _ := bot.Post(bot.BaseApi+Api_GetQQAge, Params{"qqid": qq})
	return res
}

// GetQQSex 获取性别
// Deprecated: 无效
func (bot *BotApiClient) GetQQSex(qq string) string {
	res, _ := bot.Post(bot.BaseApi+Api_GetQQSex, Params{"qqid": qq})
	return res
}

// GetFriendList 以JSON形式取得好友列表
func (bot *BotApiClient) GetFriendList(cache bool) string {
	res, _ := bot.Post(bot.BaseApi+Api_GetFriendList, Params{"cache": cache})
	return res
}

// GetGroupList 以JSON形式取得群列表
// Api_GetGroupList
func (bot *BotApiClient) GetGroupList(qq string) string {
	res, _ := bot.Post(bot.BaseApi+Api_GetGroupList, Params{"qqid": qq})
	return res
}

// GetGroupMemberList 以JSON形式取得群成员列表
// Api_GetGroupMemberList
func (bot *BotApiClient) GetGroupMemberList(qq string) string {
	res, _ := bot.Post(bot.BaseApi+Api_GetGroupMemberList, Params{"qqid": qq})
	return res
}

// GetQQInfo 以JSON形式取得某QQ个人信息
// Api_GetQQInfo
func (bot *BotApiClient) GetQQInfo(qq string) string {
	res, _ := bot.Post(bot.BaseApi+Api_GetQQInfo, Params{"qqid": qq})
	return res
}

// GetGroupInfo 以JSON形式取得某群信息
// Api_GetGroupInfo
func (bot *BotApiClient) GetGroupInfo(qq string) string {
	res, _ := bot.Post(bot.BaseApi+Api_GetGroupInfo, Params{"qqid": qq})
	return res
}

// DeleteMsg 撤回自身消息 发出消息不可以秒撤回，腾讯限制，1~2s后才可撤回
// Api_DeleteMsg
func (bot *BotApiClient) DeleteMsg(qq string) string {
	res, _ := bot.Post(bot.BaseApi+Api_DeleteMsg, Params{"qqid": qq})
	return res
}

// SetQQState 改变QQ在线状态
// @param state 1.我在线上 2.Q我吧 3.离开 4.忙碌 5.请勿打扰 6.隐身
// Api_SetQQState
func (bot *BotApiClient) SetQQState(qq string) string {
	res, _ := bot.Post(bot.BaseApi+Api_SetQQState, Params{"qqid": qq})
	return res
}

// InviteFriend 邀请好友入群
// Api_InviteFriend
func (bot *BotApiClient) InviteFriend(qq string) string {
	res, _ := bot.Post(bot.BaseApi+Api_InviteFriend, Params{"qqid": qq})
	return res
}

// GetQQInfoV2 获取qq信息 version 2
// /Api_GetQQInfo_v2
func (bot *BotApiClient) GetQQInfoV2(qq string) string {
	res, _ := bot.Post(bot.BaseApi+Api_GetQQInfoV2, Params{"qqid": qq})
	return res
}

// UpLoadPic 上传图片
// Api_UpLoadPic
func (bot *BotApiClient) UpLoadPic(qq string) string {
	res, _ := bot.Post(bot.BaseApi+Api_UpLoadPic, Params{"qqid": qq})
	return res
}

// SetPluginState 设置插件状态，开启或关闭
// Api_SetPluginState
func (bot *BotApiClient) SetPluginState(qq string) string {
	res, _ := bot.Post(bot.BaseApi+Api_SetPluginState, Params{"qqid": qq})
	return res
}

// DeleteFile 删除文件
// Api_DeleteFile
func (bot *BotApiClient) DeleteFile(qq string) string {
	res, _ := bot.Post(bot.BaseApi+Api_DeleteFile, Params{"qqid": qq})
	return res
}

// RepeatFile 转发文件
// Api_RepeatFile
func (bot *BotApiClient) RepeatFile(qq string) string {
	res, _ := bot.Post(bot.BaseApi+Api_RepeatFile, Params{"qqid": qq})
	return res
}
