package qqlt

import (
	"testing"
)

func TestGetLoginQQ(t *testing.T) {
	bot, _ := NewBotApiClient()
	//data, err := bot.Post("http://127.0.0.1:36524/api/v1/QQLight/Api_GetSoftVersion", nil)
	//fmt.Println(data, err)
	//bot.SendShake("1007139643")
	//bot.QuitGroup("1017817543")
	//bot.SetGroupName("1017817543", "测试1")
	//bot.BanSomeBody("1017817543", "", 0)
	//bot.SetAnony("1017817543", true)
	bot.SendLog("66", "测试一下")
}
