// 群发管理
package weworkapi_golang

import (
	"errors"
	"github.com/lennon7c7/weworkapi_golang/core"
	"time"
)

type GroupMsgVideo struct {
	MediaID string `json:"media_id"`
}

type GroupMsgAttachment struct {
	Msgtype     string                           `json:"msgtype"`
	Image       *GroupWelcomeTemplateImage       `json:"image,omitempty"`
	Link        *GroupWelcomeTemplateLink        `json:"link,omitempty"`
	Miniprogram *GroupWelcomeTemplateMiniprogram `json:"miniprogram,omitempty"`
	Video       *GroupMsgVideo                   `json:"video,omitempty"`
}

type GroupMsgListReq struct {
	ChatType   string `json:"chat_type"`
	StartTime  int    `json:"start_time"`
	EndTime    int    `json:"end_time"`
	Creator    string `json:"creator"`
	FilterType int    `json:"filter_type"`
	Limit      int    `json:"limit"`
	Cursor     string `json:"cursor"`
}

type GroupMsgListResp struct {
	core.Error
	NextCursor   string `json:"next_cursor"`
	GroupMsgList []struct {
		Msgid       string                    `json:"msgid"`
		Creator     string                    `json:"creator"`
		CreateTime  string                    `json:"create_time"`
		CreateType  int                       `json:"create_type"`
		Text        *GroupWelcomeTemplateText `json:"text"`
		Attachments []GroupMsgAttachment      `json:"attachments"`
	} `json:"group_msg_list"`
}

type GroupMsgTaskListReq struct {
	Msgid  string `json:"msgid"`
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type GroupMsgTaskListResp struct {
	core.Error
	NextCursor string `json:"next_cursor"`
	TaskList   []struct {
		Userid   string `json:"userid"`
		Status   int    `json:"status"`
		SendTime int    `json:"send_time"`
	} `json:"task_list"`
}

type GroupMsgResultListReq struct {
	Msgid string `json:"msgid"`
}

type GroupMsgResultListResp struct {
	core.Error
	DetailList []struct {
		ExternalUserid string `json:"external_userid"`
		ChatID         string `json:"chat_id"`
		Userid         string `json:"userid"`
		Status         int    `json:"status"`
		SendTime       int    `json:"send_time"`
	} `json:"detail_list"`
}

type GroupMsgSendResultListReq struct {
	Msgid  string `json:"msgid"`
	Userid string `json:"userid"`
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type GroupMsgSendResultListResp struct {
	core.Error
	NextCursor string `json:"next_cursor"`
	TaskList   []struct {
		Userid   string `json:"userid"`
		Status   int    `json:"status"`
		SendTime int    `json:"send_time"`
	} `json:"task_list"`
}

type GroupMsgAddReq struct {
	ChatType       string                    `json:"chat_type"`
	ExternalUserid []string                  `json:"external_userid"`
	Sender         string                    `json:"sender"`
	Text           *GroupWelcomeTemplateText `json:"text"`
	Attachments    []*GroupMsgAttachment     `json:"attachments"`
}

type GroupMsgAddResp struct {
	core.Error
	FailList []string `json:"fail_list"`
	Msgid    string   `json:"msgid"`
}

// 获取群发记录列表
// 企业和第三方应用可通过此接口获取企业与成员的群发记录。
// https://work.weixin.qq.com/api/doc/90000/90135/93338#%E8%8E%B7%E5%8F%96%E7%BE%A4%E5%8F%91%E8%AE%B0%E5%BD%95%E5%88%97%E8%A1%A8
func (s *Server) GroupMsgList(req *GroupMsgListReq) (resp *GroupMsgListResp, err error) {
	var (
		u = OpenApiUrl + "externalcontact/get_groupmsg_list_v2"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	now := time.Now()
	if req.StartTime == 0 {
		// 开始时间默认：一个月前
		d, _ := time.ParseDuration("-720h")
		d1 := now.Add(d)
		req.StartTime = int(d1.Unix())
	}
	if req.EndTime == 0 {
		// 结束时间默认：当前时间
		req.EndTime = int(now.Unix())
	}

	resp = &GroupMsgListResp{}

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

// 获取群发成员发送任务列表
// https://work.weixin.qq.com/api/doc/90000/90135/93338#%E8%8E%B7%E5%8F%96%E7%BE%A4%E5%8F%91%E6%88%90%E5%91%98%E5%8F%91%E9%80%81%E4%BB%BB%E5%8A%A1%E5%88%97%E8%A1%A8
func (s *Server) GroupMsgTaskList(req *GroupMsgTaskListReq) (resp *GroupMsgTaskListResp, err error) {
	var (
		u = OpenApiUrl + "externalcontact/get_groupmsg_task"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &GroupMsgTaskListResp{}

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

// 获取企业群发成员执行结果
// https://work.weixin.qq.com/api/doc/16251
func (s *Server) GroupMsgResultList(req *GroupMsgResultListReq) (resp *GroupMsgResultListResp, err error) {
	var (
		u = OpenApiUrl + "externalcontact/get_group_msg_result"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &GroupMsgResultListResp{}

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

// 获取企业群发成员执行结果
// https://work.weixin.qq.com/api/doc/90000/90135/93338#%E8%8E%B7%E5%8F%96%E4%BC%81%E4%B8%9A%E7%BE%A4%E5%8F%91%E6%88%90%E5%91%98%E6%89%A7%E8%A1%8C%E7%BB%93%E6%9E%9C
func (s *Server) GroupMsgSendResultList(req *GroupMsgSendResultListReq) (resp *GroupMsgSendResultListResp, err error) {
	var (
		u = OpenApiUrl + "externalcontact/get_groupmsg_send_result"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &GroupMsgSendResultListResp{}

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

// 创建企业群发
// 调用该接口并不会直接发送消息给客户/客户群，需要成员确认后才会执行发送
// https://open.work.weixin.qq.com/api/doc/90000/90135/92135
func (s *Server) GroupMsgAdd(req *GroupMsgAddReq) (resp *GroupMsgAddResp, err error) {
	var (
		u = OpenApiUrl + "externalcontact/add_msg_template"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &GroupMsgAddResp{}

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
