package weworkapi_golang_test

import (
	"github.com/go-redis/redis/v8"
	"github.com/lennon7c7/weworkapi_golang"
	"log"
	"testing"
)

func TestExternalContactList(t *testing.T) {
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

	req := weworkapi_golang.ExternalContactReq{
		UserID: "OuYangXiongDi",
	}
	resp, err := service.ExternalContactList(&req)
	if err != nil {
		t.Error("ExternalContactList err：" + err.Error())
		return
	}
	log.Println(resp)
}

func TestExternalContactGet(t *testing.T) {
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

	req := weworkapi_golang.ExternalContactGetReq{
		ExternalUserid: "wmuA1sCQAAUsh9IpSJeSRWweYqs2ekpA",
	}
	resp, err := service.ExternalContactGet(&req)
	if err != nil {
		t.Error("ExternalContactGet err：" + err.Error())
		return
	}
	log.Println(resp)
}

// 批量获取客户详情
func TestExternalContactBatchGet(t *testing.T) {
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

	req := weworkapi_golang.ExternalContactReq{
		UserID: "OuYangXiongDi",
		Limit:  100,
	}
	resp, err := service.ExternalContactBatchGet(&req)
	if err != nil {
		t.Error("ExternalContactBatchGet err：" + err.Error())
		return
	}
	log.Println(resp)
}

func TestExternalContactEditCorpTag(t *testing.T) {
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

	// 要标记的标签
	req := weworkapi_golang.ExternalContactEditCorpTagReq{
		Userid:         "OuYangXiongDi",
		ExternalUserid: "wmuA1sCQAAUsh9IpSJeSRWweYqs2ekpA",
		AddTag:         []string{"etuA1sCQAAxCyvnByidABzGxiHwuwqpw"},
	}
	// 要移除的标签
	req = weworkapi_golang.ExternalContactEditCorpTagReq{
		Userid:         "OuYangXiongDi",
		ExternalUserid: "wmuA1sCQAAUsh9IpSJeSRWweYqs2ekpA",
		RemoveTag:      []string{"etuA1sCQAAxCyvnByidABzGxiHwuwqpw"},
	}
	resp, err := service.ExternalContactEditCorpTag(&req)
	if err != nil {
		t.Error("ExternalContactEditCorpTag err：" + err.Error())
		return
	}
	log.Println(resp)
}
