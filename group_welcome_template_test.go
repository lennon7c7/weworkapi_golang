package weworkapi_golang_test

import (
	"github.com/go-redis/redis/v8"
	"github.com/lennon7c7/weworkapi_golang"
	"log"
	"testing"
)

func TestGroupWelcomeTemplateAdd(t *testing.T) {
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

	req := weworkapi_golang.GroupWelcomeTemplateReq{
		Image:       &weworkapi_golang.GroupWelcomeTemplateImage{PicURL: "https://wework.qpic.cn/wwpic/981282_5gXAwYbpQzugZnE_1622808766/0"},
		Link:        nil,
		Miniprogram: nil,
		Text:        &weworkapi_golang.GroupWelcomeTemplateText{Content: "亲爱的%NICKNAME%，你好.."},
		Notify:      0,
	}
	resp, err := service.GroupWelcomeTemplateAdd(&req)
	if err != nil {
		t.Error("GroupWelcomeTemplateAdd err：" + err.Error())
		return
	}
	log.Println(resp)
}

func TestGroupWelcomeTemplateGet(t *testing.T) {
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

	req := weworkapi_golang.GroupWelcomeTemplateReq{
		TemplateID: "msguA1sCQAADDSXBRSsR94WMSdyJjbl1Q",
	}
	resp, err := service.GroupWelcomeTemplateGet(&req)
	if err != nil {
		t.Error("GroupWelcomeTemplateGet err：" + err.Error())
		return
	}
	log.Println(resp)
}

func TestGroupWelcomeTemplateEdit(t *testing.T) {
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

	req := weworkapi_golang.GroupWelcomeTemplateReq{
		TemplateID: "msguA1sCQAADDSXBRSsR94WMSdyJjbl1Q",
		Text:       &weworkapi_golang.GroupWelcomeTemplateText{Content: "亲爱的%NICKNAME%，你好2.."},
		Notify:     0,
	}
	resp, err := service.GroupWelcomeTemplateEdit(&req)
	if err != nil {
		t.Error("GroupWelcomeTemplateEdit err：" + err.Error())
		return
	}
	log.Println(resp)
}

func TestGroupWelcomeTemplateDel(t *testing.T) {
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

	req := weworkapi_golang.GroupWelcomeTemplateReq{
		TemplateID: "msguA1sCQAADDSXBRSsR94WMSdyJjbl1Q",
	}
	resp, err := service.GroupWelcomeTemplateDel(&req)
	if err != nil {
		t.Error("GroupWelcomeTemplateDel err：" + err.Error())
		return
	}
	log.Println(resp)
}
