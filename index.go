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
	Buffer  int
	BaseApi string
}

type params map[string]interface{}

func NewBotApiClient() (*BotApiClient, error) {
	bot := &BotApiClient{
		Client:  &http.Client{},
		Buffer:  100,
		BaseApi: "http://127.0.0.1:36524/api/v1/QQLight",
	}

	return bot, nil
}

func NerParams() params {
	return make(params)
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
	_, _ = bot.Post(bot.BaseApi+"/Api_SendShake", params{"qqid": qq})
}

// GetBkn 取得网页操作时需要用到的bkn/Gtk
//func (bot *BotApiClient) GetBkn(cookies string)  {
//	_, _ = bot.Post(bot.BaseApi+"/Api_SendShake", map[string]string{"cookies": cookies})
//}

// QuitGroup 退群
func (bot *BotApiClient) QuitGroup(groupId string) {
	_, _ = bot.Post(bot.BaseApi+"/Api_QuitGroup", params{"groupID": groupId})
}

// QuitGroup 退出讨论组
func (bot *BotApiClient) QuitDiscussGroup(groupId string) {
	_, _ = bot.Post(bot.BaseApi+"/Api_QuitDiscussGroup", params{"groupID": groupId})
}

// SetDiscussName 修改讨论组名称
func (bot *BotApiClient) SetDiscussName(discussGroupID, name string) {
	_, _ = bot.Post(bot.BaseApi+"/Api_SetDiscussName", params{"discussGroupID": discussGroupID, "name": name})
}

// SetGroupName 修改群名称 测试通过 无返回值
func (bot *BotApiClient) SetGroupName(groupId, name string) {
	_, _ = bot.Post(bot.BaseApi+"/Api_SetGroupName", params{"groupID": groupId, "name": name})
}

// BanSomeBody 禁言某群某人 time为禁言时间, 单位: 秒, 0为解除禁言
func (bot *BotApiClient) BanSomeBody(groupId, qq string, time int32) {
	_, _ = bot.Post(bot.BaseApi+"/Api_Ban", params{"groupID": groupId, "qq": qq, "time": time})
}

// SetAnony 修改群匿名聊天权限
func (bot *BotApiClient) SetAnony(groupId string, isAnony bool) {
	_, _ = bot.Post(bot.BaseApi+"/Api_SetAnony", params{"groupID": groupId, "isSetAnony": isAnony})
}

// RemoveMember 群踢人 isBan不再接受此人加群
func (bot *BotApiClient) RemoveMember(groupId, qq string, isBan bool) {
	_, _ = bot.Post(bot.BaseApi+"/Api_RemoveMember", params{"groupID": groupId, "qq": qq, "isBan": isBan})
}

// RemoveGroup 解散群
func (bot *BotApiClient) RemoveGroup(groupId string) {
	_, _ = bot.Post(bot.BaseApi+"/Api_RemoveGroup", params{"groupID": groupId})
}

// GetQzoneToken 取得QQ空间Token
func (bot *BotApiClient) GetQzoneToken() string {
	res, _ := bot.Post(bot.BaseApi+"/Api_GetQzoneToken", nil)
	return res
}

// GetQzoneCookies 取得QQ空间Cookies
func (bot *BotApiClient) GetQzoneCookies() string {
	res, _ := bot.Post(bot.BaseApi+"/Api_GetQzoneCookies", nil)
	return res
}

// SendTaotao 发表说说
func (bot *BotApiClient) SendTaotao(text string) {
	_, _ = bot.Post(bot.BaseApi+"/Api_SendTaotao", params{"str": text})
}

// GetLoginQQ 取得所登录的qq
func (bot *BotApiClient) GetLoginQQ() string {
	res, _ := bot.Post(bot.BaseApi+"/Api_GetLoginQQ", nil)
	return res
}

// GetPath 框架为插件所创建的插件目录
func (bot *BotApiClient) GetPath() string {
	res, _ := bot.Post(bot.BaseApi+"/Api_GetPath", nil)
	return res
}

// PluginError 置插件错误管理 调用会导致出错
// Deprecated: 调用会导致qqlight异常退出
func (bot *BotApiClient) PluginError(code, text string) {
	res, _ := bot.Post(bot.BaseApi+"/Api_Error", params{"code": code, "str": text})
	fmt.Println(res)
}

// SendMsg 机器人发送消息，返回值为该条消息的ID 实际上无返回值
// @param msgType 消息类型 1.好友消息 2.群消息 3.群临时消息 4.讨论组消息 5.讨论组临时消息 6.QQ临时消息
// @param qqId 消息类型为2，4时可留空 即空字符串""
func (bot *BotApiClient) SendMsg(msgType int, groupId, qqId, msg string) {
	_, _ = bot.Post(bot.BaseApi+"/Api_SendMsg", params{"type": msgType, "groupID": groupId, "qqid": qqId, "msg": msg})
}

// SetGroupAdd 处理加群信息
// @param handleType 处理类型 1.同意 2.拒绝 3.忽略
// @param reason 原因, 仅拒绝时有效, 其他情况下
func (bot *BotApiClient) SetGroupAdd(groupId, qqId, seq string, handleType int32, arg ...string) {
	// 如果拒绝加群 忘记加原因, 默认为空
	if len(arg) == 0 && handleType == 2 {
		arg = append(arg, "")
	}
	_, _ = bot.Post(bot.BaseApi+"/Api_SendMsg", params{"groupID": groupId, "qqid": qqId, "seq": seq, "type": handleType, "reason": arg[0]})
}

// SetFriendAdd 处理加群信息
func (bot *BotApiClient) SetFriendAdd(qqId string, handleType int32, arg ...string) {
	if len(arg) == 0 && handleType == 2 {
		arg = append(arg, "")
	}
	_, _ = bot.Post(bot.BaseApi+"/Api_SetFriendAdd", params{"qqid": qqId, "type": handleType, "reason": arg[0]})
}

// SendLog 向框架推送一条日志
// 颜色为可变参数, 默认为0 即黑色
func (bot *BotApiClient) SendLog(logType, msg string, color ... int32) {
	if len(color) == 0 {
		color = append(color, 0)
	}
    _, _ = bot.Post(bot.BaseApi+"/Api_SendLog", params{"type": logType, "msg": msg, "fontColor": color[0]})
}
