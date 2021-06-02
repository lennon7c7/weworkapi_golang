package weworkapi_golang_test

import (
	"github.com/go-redis/redis/v8"
	"github.com/lennon7c7/weworkapi_golang"
	"log"
	"testing"
)

func TestExternalDepartmentList(t *testing.T) {
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

	//req := weworkapi_golang.DepartmentReq{
	//	ID: 99,
	//}
	//resp, err := service.DepartmentList(&req)
	resp, err := service.DepartmentList(nil)
	if err != nil {
		t.Error("DepartmentList err：" + err.Error())
		return
	}
	log.Println(resp)
}

func TestExternalDepartmentAdd(t *testing.T) {
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

	req := weworkapi_golang.DepartmentReq{
		//ID:       9,
		Name: "技术2",
		//Parentid: 0,
	}
	resp, err := service.DepartmentAdd(&req)
	if err != nil {
		t.Error("DepartmentList err：" + err.Error())
		return
	}
	log.Println(resp)
}

func TestExternalDepartmentEdit(t *testing.T) {
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

	req := weworkapi_golang.DepartmentReq{
		ID:   2,
		Name: "技术2",
	}
	resp, err := service.DepartmentEdit(&req)
	if err != nil {
		t.Error("DepartmentList err：" + err.Error())
		return
	}
	log.Println(resp)
}
