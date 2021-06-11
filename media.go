// 媒体素材管理
package weworkapi_golang

import (
	"errors"
	"github.com/lennon7c7/weworkapi_golang/core"
	"net/url"
	"strings"
)

type MediaReq struct {
	Filename string `json:"filename"`
}

type MediaResp struct {
	core.Error
	URL string `json:"url"`
}

type MediaGetReq struct {
	MediaID string `json:"media_id"`
}

// 获取素材的下载链接
func (s *Server) MediaGetDownloadURLByMediaID(req *MediaGetReq) (downloadURL string, err error) {
	var (
		u = OpenApiUrl + "media/get"
	)
	token, err := s.Token()
	if err != nil {
		return
	}
	if !strings.HasSuffix(u, "?") {
		u += "?"
	}

	v := make(url.Values)
	v.Set("access_token", token)
	v.Set("media_id", req.MediaID)
	downloadURL = u + v.Encode()

	return
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
