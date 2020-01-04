package qqlt

import (
	"fmt"
	"testing"
	"time"
)

func TestGetLoginQQ(t *testing.T) {
	bot := NewBotApiClient()
	//data, err := bot.Post("http://127.0.0.1:36524/api/v1/QQLight/Api_GetSoftVersion", nil)
	//fmt.Println(data, err)
	//bot.SendShake("")
	//bot.QuitGroup("")
	//bot.SetGroupName("", "测试1")
	//bot.BanSomeBody("", "", 0)
	//bot.SetAnony("", true)
	//bot.SendLog("66", "测试一下")
	//bot.SetFriendName("", "大号")

	//p := NewParams()
	//p.Add("11", 11)
	//p.Set("111", 111)
	fmt.Println(bot.GetFriendList(false))
	time.Sleep(100 * time.Second)
}
