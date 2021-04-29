package weworkapi_golang

import (
	"github.com/lennon7c7/weworkapi_golang/core"
	"github.com/lennon7c7/weworkapi_golang/util"
	"net/url"
	"time"
)

const (
	componentAccessTokenUrl = OpenApiUrl + "gettoken"
)

type AccessTokenServer interface {
	Token() (token string, err error)
}

type DefaultAccessTokenServer struct {
	CorpID     string
	CorpSecret string

	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"` // 当前时间 + 过期时间
}

// token不使用不获取
func (d *DefaultAccessTokenServer) Token() (token string, err error) {
	var (
		resp *AccessTokenResp
	)

	token, _ = util.CacheGetString("cachekey_of_weworkapi_golang_token")
	if token != "" {
		return
	}

	if d.ExpiresIn <= time.Now().Unix()-30 {
		resp, err = newAccessToken(&AccessTokenReq{
			CorpID:     d.CorpID,
			CorpSecret: d.CorpSecret,
		})
		if err != nil {
			return
		}

		err = util.CacheSetString("cachekey_of_weworkapi_golang_token", resp.AccessToken, time.Hour*1)
		if err != nil {
			return
		}

		d.ExpiresIn = time.Now().Unix() + resp.ExpiresIn
		d.AccessToken = resp.AccessToken
	}
	return d.AccessToken, nil
}

type AccessTokenResp struct {
	core.Error
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type AccessTokenReq struct {
	CorpID     string `json:"component_appid"`
	CorpSecret string `json:"component_appsecret"`
}

// 获取第三方应用token
func newAccessToken(r *AccessTokenReq) (*AccessTokenResp, error) {
	resp := &AccessTokenResp{}
	v := make(url.Values)
	v.Set("corpid", r.CorpID)
	v.Set("corpsecret", r.CorpSecret)

	err := core.GetRequest(componentAccessTokenUrl, v, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
