package qqlt

import "strconv"

type Message struct {
	msgType int
	qq      string
	groupId string
	msg     string
}

func NewMessage() *Message {
	return &Message{}
}

func (m *Message) Send(bot *BotApiClient) {
	switch m.msgType {
	case 0:
		panic("Message type is necessary")
	case MT_Group:
		bot.SendGroupMsg(m.groupId, m.ToString())
	case MT_Friend:
		bot.SendPrivateMsg(m.qq, m.ToString())
	}
}

func (m *Message) Dice(bot *BotApiClient, msgType int, sendId ...string) {
	switch msgType {
	case MT_Discuss:
		fallthrough
	case MT_Group:
		bot.SendMsg(msgType, sendId[0], "", m.ToString())
	case MT_QQTmp:
		fallthrough
	case MT_Friend:
		bot.SendMsg(msgType, "", sendId[0], m.ToString())
	case MT_GroupTmp:
		fallthrough
	case MT_DiscussTmp:
		bot.SendMsg(msgType, sendId[0], sendId[1], m.ToString())
	}
}

func (m *Message) QQ(qq string) *Message {
	m.qq = qq
	m.msgType = MT_Friend
	return m
}

func (m *Message) Group(groupId string) *Message {
	m.groupId = groupId
	m.msgType = MT_Group
	return m
}

func (m *Message) SendPrivate(bot *BotApiClient, qq string) {
	bot.SendPrivateMsg(qq, m.ToString())
}

func (m *Message) SendGroup(bot *BotApiClient, groupId string) {
	bot.SendGroupMsg(groupId, m.ToString())
}

func (m *Message) At(qq string) *Message {
	m.msg += "[QQ:at=" + qq + "] "
	return m
}

func (m *Message) AtAll() *Message {
	m.msg += "[QQ:at=all] "
	return m
}

func (m *Message) Text(text string) *Message {
	m.msg += text
	return m
}

func (m *Message) Image(url string) *Message {
	m.msg += "[QQ:pic=" + url + "]"
	return m
}

func (m *Message) Line() *Message {
	m.msg += "\r\n"
	return m
}

func (m *Message) Voice(url string) *Message {
	m.msg += "[QQ:voice=" + url + "]"
	return m
}

func (m *Message) Emoji(id int) *Message {
	m.msg += "[QQ:emoji=" + strconv.Itoa(id) + "]"
	return m
}

func (m *Message) Face(id int) *Message {
	m.msg += "[QQ:face=" + strconv.Itoa(id) + "]"
	return m
}

func (m *Message) FlashImg(url string) *Message {
	m.msg += "[QQ:flash,pic=" + url + "]"
	return m
}

func (m *Message) ShowImg(url string) *Message {
	m.msg += "[QQ:show,type=1,pic=" + url + "]"
	return m
}

func (m *Message) Url(url string) *Message {
	m.msg += "[QQ:url=" + url + "]"
	return m
}

func (m *Message) ToString() string {
	return m.msg
}
