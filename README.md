# qqlt
一个基于Newbe.Mahua.Framework项目的QQLight机器人框架SDK

[![GoDoc](https://godoc.org/github.com/yguilai/qqlt?status.svg)](https://godoc.org/github.com/yguilai/qqlt)

[Newbe.Mahua.Framework](https://github.com/newbe36524/Newbe.Mahua.Framework)

# 安装
```
go get -u github.com/yguilai/qqlt
```

# 快速开始
```go
func main() {
	bot := qqlt.NewBotApiClient()
	//发送消息
	bot.SendMsg(qqlt.MT_Friend, "", "QQ号", "")

	// 私聊消息
	bot.SendMsg(qqlt.MT_Friend, "", "somebody qq", "发送的内容")
	bot.SendPrivateMsg("somebody qq", "发送的内容")
	qqlt.NewMessage().
		QQ("qq").
		Text("aasdasd").
		Send(bot)

	qqlt.NewMessage().
		Text("aasdasd").
		Dice(bot, qqlt.MT_Friend, "qq号")

	// 发群消息
	bot.SendMsg(qqlt.MT_Group, "群号", "", "发送的内容")
	bot.SendGroupMsg("群号", "消息")

	qqlt.NewMessage().
		Group("群号").
		At("@的人的qq").
		Text("消息").
		Send(bot)

	qqlt.NewMessage().
		At("@的人的qq").
		Text("aasdasd").
		Dice(bot, qqlt.MT_Group, "群号")

	//
	s := qqlt.NewDefaultServer()
	go s.DefaultRun()

	for update := range s.Updates {
		fmt.Println(update)
	}
}
```
