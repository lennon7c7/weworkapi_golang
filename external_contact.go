// 客户管理
package weworkapi_golang

import (
	"errors"
	"github.com/lennon7c7/weworkapi_golang/core"
	"net/url"
)

type ExternalContactReq struct {
	UserID string `json:"userid"`
}

type ExternalContactResp struct {
	core.Error
	ExternalUserID []string `json:"external_userid"`
}

type ExternalContactGetReq struct {
	ExternalUserid string `json:"external_userid"`
}

type ExternalContactGetResp struct {
	core.Error
	ExternalContact struct {
		ExternalUserid  string `json:"external_userid"`
		Name            string `json:"name"`
		Position        string `json:"position"`
		Avatar          string `json:"avatar"`
		CorpName        string `json:"corp_name"`
		CorpFullName    string `json:"corp_full_name"`
		Type            int    `json:"type"`
		Gender          int    `json:"gender"`
		Unionid         string `json:"unionid"`
		ExternalProfile struct {
			ExternalAttr []struct {
				Type int    `json:"type"`
				Name string `json:"name"`
				Text struct {
					Value string `json:"value"`
				} `json:"text,omitempty"`
				Web struct {
					URL   string `json:"url"`
					Title string `json:"title"`
				} `json:"web,omitempty"`
				Miniprogram struct {
					Appid    string `json:"appid"`
					Pagepath string `json:"pagepath"`
					Title    string `json:"title"`
				} `json:"miniprogram,omitempty"`
			} `json:"external_attr"`
		} `json:"external_profile"`
	} `json:"external_contact"`
	FollowUser []struct {
		Userid      string `json:"userid"`
		Remark      string `json:"remark"`
		Description string `json:"description"`
		Createtime  int    `json:"createtime"`
		Tags        []struct {
			GroupName string `json:"group_name"`
			TagName   string `json:"tag_name"`
			TagID     string `json:"tag_id"`
			Type      int    `json:"type"`
		} `json:"tags,omitempty"`
		RemarkCorpName string   `json:"remark_corp_name,omitempty"`
		RemarkMobiles  []string `json:"remark_mobiles,omitempty"`
		OperUserid     string   `json:"oper_userid"`
		AddWay         int      `json:"add_way"`
		State          string   `json:"state,omitempty"`
	} `json:"follow_user"`
	NextCursor string `json:"next_cursor"`
}

// 获取客户列表
// https://open.work.weixin.qq.com/api/doc/90000/90135/92113
func (s *Server) ExternalContactList(req *ExternalContactReq) (resp *ExternalContactResp, err error) {
	var (
		u = OpenApiUrl + "externalcontact/list"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &ExternalContactResp{}

	v := make(url.Values)
	v.Set("access_token", token)
	v.Set("userid", req.UserID)
	err = core.GetRequest(u, v, resp)
	if err != nil {
		return
	}
	if resp.ErrCode != 0 || resp.ErrMsg != "ok" {
		err = errors.New(resp.ErrMsg)
		return
	}

	return
}

// 获取客户详情
// https://open.work.weixin.qq.com/api/doc/90000/90135/92114
func (s *Server) ExternalContactGet(req *ExternalContactGetReq) (resp *ExternalContactGetResp, err error) {
	var (
		u = OpenApiUrl + "externalcontact/get"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &ExternalContactGetResp{}

	v := make(url.Values)
	v.Set("access_token", token)
	v.Set("external_userid", req.ExternalUserid)
	err = core.GetRequest(u, v, resp)
	if err != nil {
		return
	}
	if resp.ErrCode != 0 || resp.ErrMsg != "ok" {
		err = errors.New(resp.ErrMsg)
		return
	}

	return
}
