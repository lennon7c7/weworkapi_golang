package weworkapi_golang

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/lennon7c7/weworkapi_golang/util"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

// open api 配置
type Config struct {
	CorpID     string
	CorpSecret string
}

func (c *Config) check() error {
	if c.CorpID == "" {
		return errors.New("CorpID was not set for Server")
	}

	if c.CorpSecret == "" {
		return errors.New("CorpSecret was not set for Server")
	}
	return nil
}

//type Handler func(c *MixedMsg)

type Server struct {
	sync.Mutex
	cfg          Config
	errorHandler WechatErrorer // 错误处理
	// 获取token
	AccessTokenServer
}

const (
	OpenApiUrl = "https://qyapi.weixin.qq.com/cgi-bin/"
)

type cipherRequestHttpBody struct {
	XMLName            struct{} `xml:"xml"`
	ToUserName         string   `xml:"ToUserName"`
	AppId              string   `xml:"AppId"` // openapi use
	Base64EncryptedMsg []byte   `xml:"Encrypt"`
}

func NewService(cfg Config, cache *redis.Client, tokenService AccessTokenServer, errHandler WechatErrorer) (s *Server, err error) {
	err = cfg.check()
	if err != nil {
		return nil, err
	}

	if cache != nil {
		util.Cache = cache
	}
	if errHandler == nil {
		errHandler = DefaultErrorHandler
	}
	if tokenService == nil {
		tokenService = &DefaultAccessTokenServer{CorpID: cfg.CorpID, CorpSecret: cfg.CorpSecret}
	}
	s = &Server{
		cfg:               cfg,
		errorHandler:      errHandler,
		AccessTokenServer: tokenService,
	}

	return s, nil
}

func (s *Server) ServeHTTP(r *http.Request) (resp *MixedMsg, err error) {
	var (
		query = r.URL.Query()

		//wantSignature string
		haveSignature = query.Get("signature")
		timestamp     = query.Get("timestamp")
		nonce         = query.Get("nonce")

		//get
		echostr = query.Get("echostr")

		//post
		wantMsgSignature string
		haveMsgSignature = query.Get("msg_signature")
		encryptType      = query.Get("encrypt_type")

		//handle vars
		data                         []byte
		requestHttpBody              = &cipherRequestHttpBody{}
		encryptedMsg                 []byte
		encryptedMsgLen              int
		msgPlaintext, haveAppIdBytes []byte
		//hand Handler
		//exist bool
	)

	if haveSignature == "" {
		err = errors.New("not found signature query parameter")
		return
	}
	if timestamp == "" {
		err = errors.New("not found timestamp query parameter")
		return
	}
	if nonce == "" {
		err = errors.New("not found nonce query parameter")
		return
	}

	//如果是验证url有效性 则echo即可
	if r.Method == "GET" {
		if echostr == "" {
			err = errors.New("not found echostr query parameter")
			return
		}
		resp = &MixedMsg{EchoStr: echostr}
		return
	}

	//进入事件执行
	if encryptType != "aes" {
		err = errors.New("unknown encrypt_type: " + encryptType)
		return
	}
	if haveMsgSignature == "" {
		err = errors.New("not found msg_signature query parameter")
		return
	}

	data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = xml.Unmarshal(data, requestHttpBody)
	if err != nil {
		return
	}

	wantMsgSignature = util.MsgSign(timestamp, nonce, string(requestHttpBody.Base64EncryptedMsg))
	if haveMsgSignature != wantMsgSignature {
		err = errors.New("check msg_signature failed, have: " + haveMsgSignature + ", want: " + wantMsgSignature)
		return
	}

	encryptedMsg = make([]byte, base64.StdEncoding.DecodedLen(len(requestHttpBody.Base64EncryptedMsg)))
	encryptedMsgLen, err = base64.StdEncoding.Decode(encryptedMsg, requestHttpBody.Base64EncryptedMsg)
	if err != nil {
		return
	}
	encryptedMsg = encryptedMsg[:encryptedMsgLen]

	//_, msgPlaintext, haveAppIdBytes, err = util.AESDecryptMsg(encryptedMsg, srv.getAESKey())
	//if err != nil {
	//	return
	//}

	if string(haveAppIdBytes) != s.cfg.CorpID {
		err = errors.New("the message AppId mismatch, have: " + string(haveAppIdBytes) + ", want: " + s.cfg.CorpID)
		return
	}
	resp = &MixedMsg{}
	if err = xml.Unmarshal(msgPlaintext, resp); err != nil {
		return
	}
	//
	//hand, exist = srv.handlerMap[resp.InfoType]
	//if !exist {
	//	err = errors.New("match handler failed :"+resp.InfoType)
	//	return
	//}
	//hand(resp)
	return
}

//用于解密数据
//func (srv *Server) AESDecryptData(cipherText, iv []byte) (rawData []byte, err error) {
//	return util.AESDecryptData(cipherText, srv.getAESKey(), iv)
//}

//url增加后缀
func (s *Server) AccessToken2url(u string) (string, error) {
	token, err := s.Token()
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(u, "?") {
		u += "?"
	}
	u += "access_token=" + token
	return u, nil
}

func (s *Server) AuthToken2url(u string, authToken string) string {
	if !strings.HasSuffix(u, "?") {
		u += "?"
	}
	u += "access_token=" + authToken
	return u
}
