// 部门管理
package weworkapi_golang

import (
	"errors"
	"github.com/lennon7c7/weworkapi_golang/core"
	"net/url"
)

type DepartmentReq struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	NameEn   string `json:"name_en"`
	Parentid int32  `json:"parentid"`
	Order    int    `json:"order"`
}

type DepartmentResp struct {
	core.Error
	Department []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		NameEn   string `json:"name_en"`
		Parentid int    `json:"parentid"`
		Order    int    `json:"order"`
	} `json:"department"`
}

// 获取部门列表
// https://open.work.weixin.qq.com/api/doc/90000/90135/90208
func (s *Server) DepartmentList(req *DepartmentReq) (resp *DepartmentResp, err error) {
	var (
		u = OpenApiUrl + "department/list"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &DepartmentResp{}

	v := make(url.Values)
	v.Set("access_token", token)
	if req != nil {
		v.Set("id", string(req.ID))
	}
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

// 创建部门
// https://open.work.weixin.qq.com/api/doc/90000/90135/90205
func (s *Server) DepartmentAdd(req *DepartmentReq) (resp *DepartmentResp, err error) {
	var (
		u = OpenApiUrl + "department/create"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &DepartmentResp{}

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

// 更新部门
// https://open.work.weixin.qq.com/api/doc/90000/90135/90206
func (s *Server) DepartmentEdit(req *DepartmentReq) (resp *DepartmentResp, err error) {
	var (
		u = OpenApiUrl + "department/update"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &DepartmentResp{}

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

// 删除部门
// https://open.work.weixin.qq.com/api/doc/90000/90135/90207
func (s *Server) DepartmentDel(req *DepartmentReq) (resp *DepartmentResp, err error) {
	var (
		u = OpenApiUrl + "department/delete"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &DepartmentResp{}

	v := make(url.Values)
	v.Set("access_token", token)
	v.Set("department_id", string(req.ID))
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
