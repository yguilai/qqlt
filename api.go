package qqlt

const (
	Api_SendShake          = "/Api_SendShake"
	Api_GetBkn             = "/Api_Getbkn"
	Api_QuitGroup          = "/Api_QuitGroup"
	Api_QuitDiscussGroup   = "/Api_QuitDiscussGroup"
	Api_SetDiscussName     = "/Api_SetDiscussName"
	Api_SetGroupName       = "/Api_SetGroupName"
	Api_BanSomeBody        = "/Api_Ban"
	Api_SetAnony           = "/Api_SetAnony"
	Api_RemoveMember       = "/Api_RemoveMember"
	Api_RemoveGroup        = "/Api_RemoveGroup"
	Api_GetQzoneToken      = "/Api_GetQzoneToken"
	Api_GetQzoneCookies    = "/Api_GetQzoneCookies"
	Api_SendTaotao         = "/Api_SendTaotao"
	Api_GetLoginQQ         = "/Api_GetLoginQQ"
	Api_GetPath            = "/Api_GetPath"
	Api_PluginError        = "/Api_Error"
	Api_SendMsg            = "/Api_SendMsg"
	Api_SetGroupAdd        = "/Api_SetGroupAdd"
	Api_SetFriendAdd       = "/Api_SetFriendAdd"
	Api_SendLog            = "/Api_SendLog"
	Api_SetFriendName      = "/Api_SetFriendName"
	Api_DeleteFriend       = "/Api_DeleteFriend"
	Api_AddGroup           = "/Api_AddGroup"
	Api_AddFriend          = "/Api_AddFriend"
	Api_GetJsonMusic       = "/Api_JsonMusic"
	Api_GetGroupCard       = "/Api_GetGroupCard"
	Api_GetNick            = "/Api_GetNick"
	Api_GetGroupName       = "/Api_GetGroupName"
	Api_GetSoftVersion     = "/Api_GetSoftVersion"
	Api_GetGroupIntroduce  = "/Api_GetGroupIntroduce"
	Api_GetGroupOwner      = "/Api_GetGroupOwner"
	Api_SetGroupCard       = "/Api_SetGroupCard"
	Api_GetPraiseNum       = "/Api_GetPraiseNum"
	Api_GetQQLevel         = "/Api_GetQQLevel"
	Api_SetNick            = "/Api_SetNick"
	Api_GetQQAge           = "/Api_GetQQAge"
	Api_GetQQSex           = "/Api_GetQQSex"
	Api_GetFriendList      = "/Api_GetFriendList"
	Api_GetGroupList       = "/Api_GetGroupList"
	Api_GetGroupMemberList = "/Api_GetGroupMemberList"
	Api_GetQQInfo          = "/Api_GetQQInfo"
	Api_GetGroupInfo       = "/Api_GetGroupInfo"
	Api_DeleteMsg          = "/Api_DeleteMsg"
	Api_SetQQState         = "/Api_SetQQState"
	Api_InviteFriend       = "/Api_InviteFriend"
	Api_GetQQInfoV2        = "/Api_GetQQInfo_v2"
	Api_UpLoadPic          = "/Api_UpLoadPic"
	Api_SetPluginState     = "/Api_SetPluginState"
	Api_DeleteFile         = "/Api_DeleteFile"
	Api_RepeatFile         = "/Api_RepeatFile"

	// 消息上报类型
	GetNewMsg = "GetNewMsg"

)

const (
	// 消息类型 Msg Type
	// 1.好友消息 2.群消息 3.群临时消息 4.讨论组消息 5.讨论组临时消息 6.QQ临时消息
	MT_Friend = iota + 1
	MT_Group
	MT_GroupTmp
	MT_Discuss
	MT_DiscussTmp
	MT_QQTmp
)