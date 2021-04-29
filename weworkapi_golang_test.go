package weworkapi_golang_test

import (
	"github.com/go-redis/redis/v8"
	"github.com/lennon7c7/weworkapi_golang"
	"log"
	"testing"
)

const (
	CorpID     = "" // 公司的id
	CorpSecret = "" // 企业应用secret
)

const (
	RedisAddr     = "127.0.0.1:6379"
	RedisPassword = "root"
	RedisDB       = 0
)

func TestGetToken(t *testing.T) {
	// 除Config外的其它参数传nil则使用默认配置.  该处代码你应该使用单例模式或服务池方式来管理
	service, err := weworkapi_golang.NewService(weworkapi_golang.Config{
		CorpID:     CorpID,
		CorpSecret: CorpSecret,
	}, redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPassword,
		DB:       RedisDB,
	}), nil, nil)
	if err != nil {
		t.Error("NewService err：" + err.Error())
		return
	}

	token, err := service.Token()
	if err != nil {
		t.Error("Token err：" + err.Error())
		return
	}
	log.Println(token)
}
