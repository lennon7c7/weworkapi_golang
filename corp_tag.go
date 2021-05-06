// 客户标签管理
package weworkapi_golang

import (
	"errors"
	"github.com/lennon7c7/weworkapi_golang/core"
)

type CorpTagReq struct {
	GroupID    string `json:"group_id"`
	GroupName  string `json:"group_name"`
	CreateTime int    `json:"create_time"`
	Order      int    `json:"order"`
	Deleted    bool   `json:"deleted"`
	Tag        []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		CreateTime int    `json:"create_time"`
		Order      int    `json:"order"`
		Deleted    bool   `json:"deleted"`
	} `json:"tag"`
}

type CorpTag struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CreateTime int    `json:"create_time"`
	Order      int    `json:"order"`
	Deleted    bool   `json:"deleted"`
}

type CorpTagListResp struct {
	core.Error
	TagGroup []CorpTagReq `json:"tag_group"`
}

type CorpTagAddResp struct {
	core.Error
	TagGroup CorpTagReq `json:"tag_group"`
}

type CorpTagDelReq struct {
	TagID   []string `json:"tag_id"`
	GroupID []string `json:"group_id"`
}

// 获取企业标签库
// https://work.weixin.qq.com/api/doc/90000/90135/92116#%E8%8E%B7%E5%8F%96%E4%BC%81%E4%B8%9A%E6%A0%87%E7%AD%BE%E5%BA%93
func (s *Server) CorpTagList() (resp *CorpTagListResp, err error) {
	var (
		u = OpenApiUrl + "externalcontact/get_corp_tag_list"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &CorpTagListResp{}

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

// 添加企业客户标签
// https://work.weixin.qq.com/api/doc/90000/90135/92117#%E6%B7%BB%E5%8A%A0%E4%BC%81%E4%B8%9A%E5%AE%A2%E6%88%B7%E6%A0%87%E7%AD%BE
func (s *Server) CorpTagAdd(req *CorpTagReq) (resp *CorpTagAddResp, err error) {
	var (
		u = OpenApiUrl + "externalcontact/add_corp_tag"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &CorpTagAddResp{}

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

// 编辑企业客户标签
// https://work.weixin.qq.com/api/doc/90000/90135/92117#%E7%BC%96%E8%BE%91%E4%BC%81%E4%B8%9A%E5%AE%A2%E6%88%B7%E6%A0%87%E7%AD%BE
func (s *Server) CorpTagEdit(req *CorpTag) (resp *core.Error, err error) {
	var (
		u = OpenApiUrl + "externalcontact/edit_corp_tag"
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

// 删除企业客户标签
// https://work.weixin.qq.com/api/doc/90000/90135/92117#%E5%88%A0%E9%99%A4%E4%BC%81%E4%B8%9A%E5%AE%A2%E6%88%B7%E6%A0%87%E7%AD%BE
func (s *Server) CorpTagDel(req *CorpTagDelReq) (resp *core.Error, err error) {
	var (
		u = OpenApiUrl + "externalcontact/del_corp_tag"
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
