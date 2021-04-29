// 入群欢迎语素材管理
package weworkapi_golang

import (
	"errors"
	"github.com/lennon7c7/weworkapi_golang/core"
)

type GroupWelcomeTemplateReq struct {
	TemplateID string `json:"template_id"`
	Text       struct {
		Content string `json:"content,omitempty"`
	} `json:"text,omitempty"`
	//Image struct {
	//	MediaID string `json:"media_id,omitempty"`
	//	PicURL  string `json:"pic_url,omitempty"`
	//} `json:"image,omitempty"`
	//Link struct {
	//	Title  string `json:"title,omitempty"`
	//	Picurl string `json:"picurl,omitempty"`
	//	Desc   string `json:"desc,omitempty"`
	//	URL    string `json:"url,omitempty"`
	//} `json:"link,omitempty"`
	//Miniprogram struct {
	//	Title      string `json:"title,omitempty"`
	//	PicMediaID string `json:"pic_media_id,omitempty"`
	//	Appid      string `json:"appid,omitempty"`
	//	Page       string `json:"page,omitempty"`
	//} `json:"miniprogram,omitempty"`
	Notify int `json:"notify,omitempty"`
}

type GroupWelcomeTemplateResp struct {
	core.Error
	TemplateID string `json:"template_id"`
	Text       struct {
		Content string `json:"content"`
	} `json:"text"`
	Image struct {
		PicURL string `json:"pic_url"`
	} `json:"image"`
	Link struct {
		Title  string `json:"title"`
		Picurl string `json:"picurl"`
		Desc   string `json:"desc"`
		URL    string `json:"url"`
	} `json:"link"`
	Miniprogram struct {
		Title      string `json:"title"`
		PicMediaID string `json:"pic_media_id"`
		Appid      string `json:"appid"`
		Page       string `json:"page"`
	} `json:"miniprogram"`
}

// 添加入群欢迎语素材
// https://open.work.weixin.qq.com/api/doc/90000/90135/92366#%E6%B7%BB%E5%8A%A0%E5%85%A5%E7%BE%A4%E6%AC%A2%E8%BF%8E%E8%AF%AD%E7%B4%A0%E6%9D%90
func (s *Server) GroupWelcomeTemplateAdd(req *GroupWelcomeTemplateReq) (resp *GroupWelcomeTemplateResp, err error) {
	var (
		u = OpenApiUrl + "externalcontact/group_welcome_template/add"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &GroupWelcomeTemplateResp{}

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

// 编辑入群欢迎语素材
// https://open.work.weixin.qq.com/api/doc/90000/90135/92366#%E7%BC%96%E8%BE%91%E5%85%A5%E7%BE%A4%E6%AC%A2%E8%BF%8E%E8%AF%AD%E7%B4%A0%E6%9D%90
func (s *Server) GroupWelcomeTemplateEdit(req *GroupWelcomeTemplateReq) (resp *core.Error, err error) {
	var (
		u = OpenApiUrl + "externalcontact/group_welcome_template/edit"
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

// 获取入群欢迎语素材
// https://open.work.weixin.qq.com/api/doc/90000/90135/92366#%E8%8E%B7%E5%8F%96%E5%85%A5%E7%BE%A4%E6%AC%A2%E8%BF%8E%E8%AF%AD%E7%B4%A0%E6%9D%90
func (s *Server) GroupWelcomeTemplateGet(req *GroupWelcomeTemplateResp) (resp *GroupWelcomeTemplateResp, err error) {
	var (
		u = OpenApiUrl + "externalcontact/group_welcome_template/get"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &GroupWelcomeTemplateResp{}

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

// 删除入群欢迎语素材
// https://open.work.weixin.qq.com/api/doc/90000/90135/92366#%E5%88%A0%E9%99%A4%E5%85%A5%E7%BE%A4%E6%AC%A2%E8%BF%8E%E8%AF%AD%E7%B4%A0%E6%9D%90
func (s *Server) GroupWelcomeTemplateDel(req *GroupWelcomeTemplateResp) (resp *core.Error, err error) {
	var (
		u = OpenApiUrl + "externalcontact/group_welcome_template/del"
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
