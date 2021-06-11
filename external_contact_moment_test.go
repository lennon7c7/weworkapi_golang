package weworkapi_golang_test

import (
	"github.com/go-redis/redis/v8"
	"github.com/lennon7c7/weworkapi_golang"
	"log"
	"testing"
)

// 获取企业全部的发表列表
func TestExternalContactMomentList(t *testing.T) {
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

	req := weworkapi_golang.ExternalContactMomentListReq{
		//StartTime:  0,
		//EndTime:    0,
		//Creator:    "",
		//FilterType: 0,
		//Cursor:     "",
		//Limit:      0,
	}
	resp, err := service.ExternalContactMomentList(&req)
	if err != nil {
		t.Error("ExternalContactMomentList err：" + err.Error())
		return
	}
	log.Println(resp)
	if resp.ExternalContactMomentList != nil {
		log.Println(resp.ExternalContactMomentList[0])
		log.Println(resp.ExternalContactMomentList[0].MomentID)

		if resp.ExternalContactMomentList[0].Text != nil {
			log.Println(resp.ExternalContactMomentList[0].Text.Content)
		}

		//if resp.ExternalContactMomentList[0].Image != nil {
		//	for _, value := range resp.ExternalContactMomentList[0].Image {
		//		log.Println(value.MediaID)
		//	}
		//}
	}
}

// 获取客户朋友圈企业发表的列表
func TestExternalContactMomentTaskList(t *testing.T) {
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

	req := weworkapi_golang.ExternalContactMomentTaskListReq{
		MomentID: "mom1uA1sCQAAA6CepB5SVcmXQ5DzPjqamQ",
	}
	resp, err := service.ExternalContactMomentTaskList(&req)
	if err != nil {
		t.Error("ExternalContactMomentTaskList err：" + err.Error())
		return
	}
	log.Println(resp)
}

// 获取客户朋友圈发表时选择的可见范围
func TestExternalContactMomentCustomerList(t *testing.T) {
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

	req := weworkapi_golang.ExternalContactMomentCustomerListReq{
		MomentID: "mom1uA1sCQAAA6CepB5SVcmXQ5DzPjqamQ",
		UserID:   "OuYangXiongDi",
		//Limit:    0,
		//Cursor:   "",
	}
	resp, err := service.ExternalContactMomentCustomerList(&req)
	if err != nil {
		t.Error("ExternalContactMomentResultList err：" + err.Error())
		return
	}
	log.Println(resp)
}

// 获取客户朋友圈发表后的可见客户列表
func TestExternalContactMomentSendResultList(t *testing.T) {
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

	req := weworkapi_golang.ExternalContactMomentSendResultListReq{
		MomentID: "mom1uA1sCQAAA6CepB5SVcmXQ5DzPjqamQ",
		UserID:   "OuYangXiongDi",
		//Limit:    0,
		//Cursor:   "",
	}
	resp, err := service.ExternalContactMomentSendResultList(&req)
	if err != nil {
		t.Error("ExternalContactMomentSendResultList err：" + err.Error())
		return
	}
	log.Println(resp)
}

// 获取客户朋友圈的互动数据
func TestExternalContactMomentCommentList(t *testing.T) {
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

	req := weworkapi_golang.ExternalContactMomentCommentListReq{
		MomentID: "mom1uA1sCQAAA6CepB5SVcmXQ5DzPjqamQ",
		UserID:   "OuYangXiongDi",
		//Limit:    0,
		//Cursor:   "",
	}
	resp, err := service.ExternalContactMomentCommentList(&req)
	if err != nil {
		t.Error("ExternalContactMomentCommentList err：" + err.Error())
		return
	}
	log.Println(resp)
}
