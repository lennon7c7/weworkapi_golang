// 客户朋友圈管理
package weworkapi_golang

import (
	"errors"
	"github.com/lennon7c7/weworkapi_golang/core"
	"time"
)

type ExternalContactMomentImage struct {
	MediaID string `json:"media_id"`
}

type ExternalContactMomentLink struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type ExternalContactMomentVideo struct {
	MediaID      string `json:"media_id"`
	ThumbMediaID string `json:"thumb_media_id"`
}

type ExternalContactMomentLocation struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Name      string `json:"name"`
}

type ExternalContactMoment struct {
	MomentID    string                         `json:"moment_id"`
	Creator     string                         `json:"creator"`
	CreateTime  int                            `json:"create_time"`
	CreateType  int                            `json:"create_type"`
	VisibleType int                            `json:"visible_type"`
	Text        *GroupWelcomeTemplateText      `json:"text,omitempty"`
	Image       []*ExternalContactMomentImage  `json:"image,omitempty"`
	Video       *ExternalContactMomentVideo    `json:"video,omitempty"`
	Link        *ExternalContactMomentLink     `json:"link,omitempty"`
	Location    *ExternalContactMomentLocation `json:"location,omitempty"`
}

type ExternalContactMomentListReq struct {
	StartTime  int    `json:"start_time"`
	EndTime    int    `json:"end_time"`
	Creator    string `json:"creator,omitempty"`
	FilterType int    `json:"filter_type,omitempty"`
	Cursor     string `json:"cursor,omitempty"`
	Limit      int    `json:"limit,omitempty"`
}

type ExternalContactMomentListResp struct {
	core.Error
	NextCursor                string                   `json:"next_cursor"`
	ExternalContactMomentList []*ExternalContactMoment `json:"moment_list"`
}

type ExternalContactMomentTask struct {
	UserID        string `json:"userid"`
	PublishStatus int    `json:"publish_status"`
}

type ExternalContactMomentTaskListReq struct {
	MomentID string `json:"moment_id"`
	Limit    int    `json:"limit"`
	Cursor   string `json:"cursor"`
}

type ExternalContactMomentTaskListResp struct {
	core.Error
	NextCursor string                       `json:"next_cursor"`
	TaskList   []*ExternalContactMomentTask `json:"task_list"`
}

type ExternalContactMomentCustomer struct {
	UserID         string `json:"userid"`
	ExternalUserID string `json:"external_userid"`
}

type ExternalContactMomentCustomerListReq struct {
	MomentID string `json:"moment_id"`
	UserID   string `json:"userid"`
	Limit    int    `json:"limit"`
	Cursor   string `json:"cursor"`
}

type ExternalContactMomentCustomerListResp struct {
	core.Error
	NextCursor   string                           `json:"next_cursor"`
	CustomerList []*ExternalContactMomentCustomer `json:"customer_list"`
}

type ExternalContactMomentSendResult struct {
	ExternalUserID string `json:"external_userid"`
}

type ExternalContactMomentSendResultListReq struct {
	MomentID string `json:"moment_id"`
	UserID   string `json:"userid"`
	Limit    int    `json:"limit"`
	Cursor   string `json:"cursor"`
}

type ExternalContactMomentSendResultListResp struct {
	core.Error
	NextCursor     string                             `json:"next_cursor"`
	SendResultList []*ExternalContactMomentSendResult `json:"customer_list"`
}

type ExternalContactMomentComment struct {
	ExternalUserID string `json:"external_userid,omitempty"`
	CreateTime     int    `json:"create_time"`
	UserID         string `json:"userid,omitempty"`
}

type ExternalContactMomentCommentListReq struct {
	MomentID string `json:"moment_id"`
	UserID   string `json:"userid"`
}

type ExternalContactMomentCommentListResp struct {
	core.Error
	NextCursor  string                          `json:"next_cursor"`
	CommentList []*ExternalContactMomentComment `json:"comment_list"`
	LikeList    []*ExternalContactMomentComment `json:"like_list"`
}

// 获取企业全部的发表列表
// 企业和第三方应用可通过该接口获取企业全部的发表内容。
// https://work.weixin.qq.com/api/doc/90000/90135/93333#%E8%8E%B7%E5%8F%96%E4%BC%81%E4%B8%9A%E5%85%A8%E9%83%A8%E7%9A%84%E5%8F%91%E8%A1%A8%E5%88%97%E8%A1%A8
func (s *Server) ExternalContactMomentList(req *ExternalContactMomentListReq) (resp *ExternalContactMomentListResp, err error) {
	var (
		u = OpenApiUrl + "externalcontact/get_moment_list"
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

	resp = &ExternalContactMomentListResp{}

	var externalContactMoments []*ExternalContactMoment
	for {
		err = core.PostJson(s.AuthToken2url(u, token), req, resp)
		if err != nil {
			return
		}
		if resp.ErrCode != 0 {
			err = errors.New(resp.ErrMsg)
			return
		}

		externalContactMoments = append(externalContactMoments, resp.ExternalContactMomentList...)
		req.Cursor = resp.NextCursor
		if resp.NextCursor == "" {
			break
		}
	}

	resp.ExternalContactMomentList = externalContactMoments

	return
}

// 获取客户朋友圈企业发表的列表
// 企业和第三方应用可通过该接口获取企业发表的朋友圈成员执行情况
// https://work.weixin.qq.com/api/doc/90000/90135/93333#%E8%8E%B7%E5%8F%96%E5%AE%A2%E6%88%B7%E6%9C%8B%E5%8F%8B%E5%9C%88%E4%BC%81%E4%B8%9A%E5%8F%91%E8%A1%A8%E7%9A%84%E5%88%97%E8%A1%A8
func (s *Server) ExternalContactMomentTaskList(req *ExternalContactMomentTaskListReq) (resp *ExternalContactMomentTaskListResp, err error) {
	var (
		u = OpenApiUrl + "externalcontact/get_moment_task"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &ExternalContactMomentTaskListResp{}

	var externalContactMoments []*ExternalContactMomentTask
	for {
		err = core.PostJson(s.AuthToken2url(u, token), req, resp)
		if err != nil {
			return
		}
		if resp.ErrCode != 0 {
			err = errors.New(resp.ErrMsg)
			return
		}

		externalContactMoments = append(externalContactMoments, resp.TaskList...)
		req.Cursor = resp.NextCursor
		if resp.NextCursor == "" {
			break
		}
	}

	resp.TaskList = externalContactMoments

	return
}

// 获取客户朋友圈发表时选择的可见范围
// 企业和第三方应用可通过该接口获取客户朋友圈创建时，选择的客户可见范围
// https://work.weixin.qq.com/api/doc/90000/90135/93333#%E8%8E%B7%E5%8F%96%E5%AE%A2%E6%88%B7%E6%9C%8B%E5%8F%8B%E5%9C%88%E5%8F%91%E8%A1%A8%E6%97%B6%E9%80%89%E6%8B%A9%E7%9A%84%E5%8F%AF%E8%A7%81%E8%8C%83%E5%9B%B4
func (s *Server) ExternalContactMomentCustomerList(req *ExternalContactMomentCustomerListReq) (resp *ExternalContactMomentCustomerListResp, err error) {
	var (
		u = OpenApiUrl + "externalcontact/get_moment_customer_list"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &ExternalContactMomentCustomerListResp{}

	var externalContactMoments []*ExternalContactMomentCustomer
	for {
		err = core.PostJson(s.AuthToken2url(u, token), req, resp)
		if err != nil {
			return
		}
		if resp.ErrCode != 0 {
			err = errors.New(resp.ErrMsg)
			return
		}

		externalContactMoments = append(externalContactMoments, resp.CustomerList...)
		req.Cursor = resp.NextCursor
		if resp.NextCursor == "" {
			break
		}
	}

	resp.CustomerList = externalContactMoments

	return
}

// 获取客户朋友圈发表后的可见客户列表
// 企业和第三方应用可通过该接口获取客户朋友圈发表后，可在微信朋友圈中查看的客户列表
// https://work.weixin.qq.com/api/doc/90000/90135/93333#%E8%8E%B7%E5%8F%96%E5%AE%A2%E6%88%B7%E6%9C%8B%E5%8F%8B%E5%9C%88%E5%8F%91%E8%A1%A8%E5%90%8E%E7%9A%84%E5%8F%AF%E8%A7%81%E5%AE%A2%E6%88%B7%E5%88%97%E8%A1%A8
func (s *Server) ExternalContactMomentSendResultList(req *ExternalContactMomentSendResultListReq) (resp *ExternalContactMomentSendResultListResp, err error) {
	var (
		u = OpenApiUrl + "externalcontact/get_moment_send_result"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &ExternalContactMomentSendResultListResp{}

	var externalContactMoments []*ExternalContactMomentSendResult
	for {
		err = core.PostJson(s.AuthToken2url(u, token), req, resp)
		if err != nil {
			return
		}
		if resp.ErrCode != 0 {
			err = errors.New(resp.ErrMsg)
			return
		}

		externalContactMoments = append(externalContactMoments, resp.SendResultList...)
		req.Cursor = resp.NextCursor
		if resp.NextCursor == "" {
			break
		}
	}

	resp.SendResultList = externalContactMoments

	return
}

// 获取客户朋友圈的互动数据
// 企业和第三方应用可通过此接口获取客户朋友圈的互动数据
// https://work.weixin.qq.com/api/doc/90000/90135/93333#%E8%8E%B7%E5%8F%96%E5%AE%A2%E6%88%B7%E6%9C%8B%E5%8F%8B%E5%9C%88%E7%9A%84%E4%BA%92%E5%8A%A8%E6%95%B0%E6%8D%AE
func (s *Server) ExternalContactMomentCommentList(req *ExternalContactMomentCommentListReq) (resp *ExternalContactMomentCommentListResp, err error) {
	var (
		u = OpenApiUrl + "externalcontact/get_moment_comments"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &ExternalContactMomentCommentListResp{}

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
