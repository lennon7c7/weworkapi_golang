// 成员管理
package weworkapi_golang

import (
	"errors"
	"github.com/lennon7c7/weworkapi_golang/core"
	"net/url"
)

type UserListReq struct {
	DepartmentID string `json:"department_id"` // 获取的部门id
	FetchChild   string `json:"fetch_child"`   // 是否递归获取子部门下面的成员：1-递归获取，0-只获取本部门
}

type UserReq struct {
	UserID     string `json:"userid"`
	Name       string `json:"name"`
	Department []int  `json:"department"`
	//Alias          string `json:"alias"`
	//Mobile         string `json:"mobile"`
	//Order          []int  `json:"order"`
	//Position       string `json:"position"`
	//Gender         string `json:"gender"`
	//Email          string `json:"email"`
	//IsLeaderInDept []int  `json:"is_leader_in_dept"`
	//Enable         int    `json:"enable"`
	//AvatarMediaid  string `json:"avatar_mediaid"`
	//Telephone      string `json:"telephone"`
	//Address        string `json:"address"`
	//MainDepartment int    `json:"main_department"`
	//Extattr        struct {
	//	Attrs []struct {
	//		Type int    `json:"type"`
	//		Name string `json:"name"`
	//		Text struct {
	//			Value string `json:"value"`
	//		} `json:"text,omitempty"`
	//		Web struct {
	//			URL   string `json:"url"`
	//			Title string `json:"title"`
	//		} `json:"web,omitempty"`
	//	} `json:"attrs"`
	//} `json:"extattr"`
	//ToInvite         bool   `json:"to_invite"`
	//ExternalPosition string `json:"external_position"`
	//ExternalProfile  struct {
	//	ExternalCorpName string `json:"external_corp_name"`
	//	ExternalAttr     []struct {
	//		Type int    `json:"type"`
	//		Name string `json:"name"`
	//		Text struct {
	//			Value string `json:"value"`
	//		} `json:"text,omitempty"`
	//		Web struct {
	//			URL   string `json:"url"`
	//			Title string `json:"title"`
	//		} `json:"web,omitempty"`
	//		Miniprogram struct {
	//			Appid    string `json:"appid"`
	//			Pagepath string `json:"pagepath"`
	//			Title    string `json:"title"`
	//		} `json:"miniprogram,omitempty"`
	//	} `json:"external_attr"`
	//} `json:"external_profile"`
}

type UserListResp struct {
	core.Error
	UserList []UserResp `json:"userlist"`
}

type UserResp struct {
	core.Error
	Userid         string `json:"userid"`
	Name           string `json:"name"`
	Department     []int  `json:"department"`
	Order          []int  `json:"order"`
	Position       string `json:"position"`
	Mobile         string `json:"mobile"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	IsLeaderInDept []int  `json:"is_leader_in_dept"`
	Avatar         string `json:"avatar"`
	ThumbAvatar    string `json:"thumb_avatar"`
	Telephone      string `json:"telephone"`
	Alias          string `json:"alias"`
	Address        string `json:"address"`
	OpenUserid     string `json:"open_userid"`
	MainDepartment int    `json:"main_department"`
	Extattr        struct {
		Attrs []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				URL   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
		} `json:"attrs"`
	} `json:"extattr"`
	Status           int    `json:"status"`
	QrCode           string `json:"qr_code"`
	ExternalPosition string `json:"external_position"`
	ExternalProfile  struct {
		ExternalCorpName string `json:"external_corp_name"`
		ExternalAttr     []struct {
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
}

// 获取部门成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90200
func (s *Server) UserList(req *UserListReq) (resp *UserListResp, err error) {
	var (
		u = OpenApiUrl + "user/list"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &UserListResp{}

	v := make(url.Values)
	v.Set("access_token", token)
	if req != nil {
		v.Set("department_id", req.DepartmentID)
		v.Set("fetch_child", req.FetchChild)
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

// 创建成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90195
func (s *Server) UserAdd(req *UserReq) (resp *UserResp, err error) {
	var (
		u = OpenApiUrl + "user/create"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &UserResp{}

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

// 更新成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90197
func (s *Server) UserEdit(req *UserReq) (resp *core.Error, err error) {
	var (
		u = OpenApiUrl + "user/update"
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
	if resp.ErrCode != 0 || resp.ErrMsg != "ok" {
		err = errors.New(resp.ErrMsg)
		return
	}

	return
}

// 读取成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90196
func (s *Server) UserGet(req *UserReq) (resp *UserResp, err error) {
	var (
		u = OpenApiUrl + "user/get"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &UserResp{}

	//v := make(url.Values)
	//v.Set("userid", req.UserID)
	//err = core.GetRequest(s.AuthToken2url(u, token), v, resp)

	//err = core.GetRequest(u, core.AuthTokenUrlValues(token), resp)

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

// 删除成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90198
func (s *Server) UserDel(req *UserReq) (resp *core.Error, err error) {
	var (
		u = OpenApiUrl + "user/delete"
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
	if resp.ErrCode != 0 || resp.ErrMsg != "ok" {
		err = errors.New(resp.ErrMsg)
		return
	}

	return
}
