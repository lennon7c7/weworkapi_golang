// 媒体素材管理
package weworkapi_golang

import (
	"errors"
	"github.com/lennon7c7/weworkapi_golang/core"
)

type MediaReq struct {
	Filename string `json:"filename"`
}

type MediaResp struct {
	core.Error
	URL string `json:"url"`
}

// 上传图片
// https://open.work.weixin.qq.com/api/doc/90000/90135/90256
func (s *Server) MediaUploadImg(req *MediaReq) (resp *MediaResp, err error) {
	var (
		u = OpenApiUrl + "media/uploadimg"
	)
	token, err := s.Token()
	if err != nil {
		return
	}

	resp = &MediaResp{}

	err = core.PostFile(s.AuthToken2url(u, token), req.Filename, resp)
	if err != nil {
		return
	}
	if resp.ErrCode != 0 || resp.ErrMsg != "ok" {
		err = errors.New(resp.ErrMsg)
		return
	}

	return
}
