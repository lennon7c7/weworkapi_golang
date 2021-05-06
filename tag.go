// 标签管理
package weworkapi_golang

import (
	"errors"
	"github.com/lennon7c7/weworkapi_golang/core"
	"net/url"
	"strconv"
)

type TagReq struct {
	Tagname string `json:"tagname"`
	Tagid   int    `json:"tagid"`
}

type TagListResp struct {
	core.Error
	Taglist []TagReq `json:"taglist"`
}

type TagAddResp struct {
	core.Error
	Tagid int `json:"tagid"`
}

type TagUserResp struct {
	core.Error
	Tagname  string `json:"tagname"`
	Userlist []struct {
		Userid string `json:"userid"`
		Name   string `json:"name"`
	} `json:"userlist"`
	Partylist []int `json:"partylist"`
}

type TagUserReq struct {
	Tagid     int      `json:"tagid"`
	Userlist  []string `json:"userlist"`
	Partylist []int    `json:"partylist"`
}

type TagUserAddResp struct {
	core.Error
	Invalidlist  string `json:"invalidlist"`
	Invalidparty []int  `json:"invalidparty"`
}

// 获取标签列表
// https://open.work.weixin.qq.com/api/doc/90000/90135/90216
func (s *Server) TagList() (resp *TagListResp, err error) {
	var (
		u = OpenApiUrl + "tag/list"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &TagListResp{}

	err = core.GetRequest(u, core.AuthTokenUrlValues(token), resp)
	if err != nil {
		return
	}
	if resp.ErrCode != 0 || resp.ErrMsg != "ok" {
		err = errors.New(resp.ErrMsg)
		return
	}

	return
}

// 创建标签
// https://open.work.weixin.qq.com/api/doc/90000/90135/90210
func (s *Server) TagAdd(req *TagReq) (resp *TagAddResp, err error) {
	var (
		u = OpenApiUrl + "tag/create"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &TagAddResp{}

	err = core.PostJson(s.AuthToken2url(u, token), req, resp)
	if err != nil {
		return
	}
	if resp.ErrCode != 0 {
		err = errors.New(resp.ErrMsg)
		return
	}

	return
}

// 更新标签
// https://open.work.weixin.qq.com/api/doc/90000/90135/90211
func (s *Server) TagEdit(req *TagReq) (resp *core.Error, err error) {
	var (
		u = OpenApiUrl + "tag/update"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &core.Error{}

	err = core.PostJson(s.AuthToken2url(u, token), req, resp)
	if err != nil {
		return
	}
	if resp.ErrCode != 0 {
		err = errors.New(resp.ErrMsg)
		return
	}

	return
}

// 删除标签
// https://open.work.weixin.qq.com/api/doc/90000/90135/90212
func (s *Server) TagDel(req *TagReq) (resp *TagListResp, err error) {
	var (
		u = OpenApiUrl + "tag/delete"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &TagListResp{}

	v := make(url.Values)
	v.Set("access_token", token)
	v.Set("tagid", strconv.Itoa(req.Tagid))
	err = core.GetRequest(u, v, resp)
	if err != nil {
		return
	}
	if resp.ErrCode != 0 {
		err = errors.New(resp.ErrMsg)
		return
	}

	return
}

// 获取标签成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90213
func (s *Server) TagUserGet(req *TagReq) (resp *TagUserResp, err error) {
	var (
		u = OpenApiUrl + "tag/get"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &TagUserResp{}

	v := make(url.Values)
	v.Set("access_token", token)
	v.Set("tagid", strconv.Itoa(req.Tagid))
	err = core.GetRequest(u, v, resp)
	if err != nil {
		return
	}
	if resp.ErrCode != 0 {
		err = errors.New(resp.ErrMsg)
		return
	}

	return
}

// 增加标签成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90214
func (s *Server) TagUserAdd(req *TagUserReq) (resp *TagUserAddResp, err error) {
	var (
		u = OpenApiUrl + "tag/addtagusers"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &TagUserAddResp{}

	err = core.PostJson(s.AuthToken2url(u, token), req, resp)
	if err != nil {
		return
	}
	if resp.ErrCode != 0 {
		err = errors.New(resp.ErrMsg)
		return
	}

	return
}

// 删除标签成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90215
func (s *Server) TagUserDel(req *TagUserReq) (resp *TagUserAddResp, err error) {
	var (
		u = OpenApiUrl + "tag/deltagusers"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &TagUserAddResp{}

	err = core.PostJson(s.AuthToken2url(u, token), req, resp)
	if err != nil {
		return
	}
	if resp.ErrCode != 0 {
		err = errors.New(resp.ErrMsg)
		return
	}

	return
}
