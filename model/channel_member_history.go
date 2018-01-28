
package model

type ChannelMemberHistory struct {
	ChannelId string
	UserId    string
	UserEmail string `db:"Email"`
	JoinTime  int64
	LeaveTime *int64
}
