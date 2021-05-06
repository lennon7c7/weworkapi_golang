// 客户群管理
package weworkapi_golang

import (
	"errors"
	"github.com/lennon7c7/weworkapi_golang/core"
)

type OwnerFilter struct {
	UseridList []string `json:"userid_list"`
}

type GroupChatList struct {
	ChatID string `json:"chat_id"`
	Status int    `json:"status"`
}

type Invitor struct {
	Userid string `json:"userid"`
}

type MemberList struct {
	Userid    string  `json:"userid"`
	Type      int     `json:"type"`
	JoinTime  int     `json:"join_time"`
	JoinScene int     `json:"join_scene"`
	Invitor   Invitor `json:"invitor,omitempty"`
	Unionid   string  `json:"unionid,omitempty"`
}

type AdminList struct {
	Userid string `json:"userid"`
}

type GroupChat struct {
	ChatID     string       `json:"chat_id"`
	Name       string       `json:"name"`
	Owner      string       `json:"owner"`
	CreateTime int          `json:"create_time"`
	Notice     string       `json:"notice"`
	MemberList []MemberList `json:"member_list"`
	AdminList  []AdminList  `json:"admin_list"`
}

type GroupChatListReq struct {
	StatusFilter int         `json:"status_filter"`
	OwnerFilter  OwnerFilter `json:"owner_filter"`
	Cursor       string      `json:"cursor"`
	Limit        int         `json:"limit"`
}

type GroupChatListResp struct {
	core.Error
	GroupChatList []GroupChatList `json:"group_chat_list"`
	NextCursor    string          `json:"next_cursor"`
}

type GroupChatGetReq struct {
	ChatID string `json:"chat_id"`
}

type GroupChatGetResp struct {
	core.Error
	GroupChat GroupChat `json:"group_chat"`
}

// 获取客户群列表
// https://work.weixin.qq.com/api/doc/90000/90135/92120
func (s *Server) GroupChatList(req *GroupChatListReq) (resp *GroupChatListResp, err error) {
	var (
		u = OpenApiUrl + "externalcontact/groupchat/list"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &GroupChatListResp{}

	err = core.PostJson(s.AuthToken2url(u, token), req, resp)
	if err != nil {
		return
	}
	if resp.ErrCode != 0 || resp.ErrMsg != "ok" {
		err = errors.New(resp.ErrMsg)
		return
	}

	return
}

// 获取客户群详情
// https://work.weixin.qq.com/api/doc/90000/90135/92122
func (s *Server) GroupChatGet(req *GroupChatGetReq) (resp *GroupChatGetResp, err error) {
	var (
		u = OpenApiUrl + "externalcontact/groupchat/get"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &GroupChatGetResp{}

	err = core.PostJson(s.AuthToken2url(u, token), req, resp)
	if err != nil {
		return
	}
	if resp.ErrCode != 0 || resp.ErrMsg != "ok" {
		err = errors.New(resp.ErrMsg)
		return
	}

	return
}
