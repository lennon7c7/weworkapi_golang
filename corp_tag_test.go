package weworkapi_golang_test

import (
	"github.com/go-redis/redis/v8"
	"github.com/lennon7c7/weworkapi_golang"
	"log"
	"testing"
)

func TestCorpTagList(t *testing.T) {
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

	resp, err := service.CorpTagList()
	if err != nil {
		t.Error("CorpTagList err：" + err.Error())
		return
	}
	log.Println(resp)
}

func TestCorpTagAdd(t *testing.T) {
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

	// 添加新的标签组
	req := weworkapi_golang.CorpTagReq{
		GroupName: "客户购买意愿等级",
		Tag: []struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			CreateTime int    `json:"create_time"`
			Order      int    `json:"order"`
			Deleted    bool   `json:"deleted"`
		}{
			{Name: "不感兴趣", Order: 1},
			{Name: "有意向", Order: 2},
			{Name: "已购买", Order: 3},
		},
	}
	// 添加新的标签
	req = weworkapi_golang.CorpTagReq{
		GroupID:   "etuA1sCQAA97QopllG8bWWqNEIuEOPtQ",
		GroupName: "客户购买意愿等级",
		Tag: []struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			CreateTime int    `json:"create_time"`
			Order      int    `json:"order"`
			Deleted    bool   `json:"deleted"`
		}{
			{Name: "有意向", Order: 2},
		},
	}
	resp, err := service.CorpTagAdd(&req)
	if err != nil {
		t.Error("CorpTagAdd err：" + err.Error())
		return
	}
	log.Println(resp)
}

func TestCorpTagEdit(t *testing.T) {
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

	req := weworkapi_golang.CorpTag{
		ID:    "etuA1sCQAA0xOFcaWrpCTqxAwI5wpMUw",
		Name:  "有意向购买",
		Order: 2,
	}
	resp, err := service.CorpTagEdit(&req)
	if err != nil {
		t.Error("CorpTagEdit err：" + err.Error())
		return
	}
	log.Println(resp)
}

func TestCorpTagDel(t *testing.T) {
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

	req := weworkapi_golang.CorpTagDelReq{
		TagID: []string{"etuA1sCQAA0xOFcaWrpCTqxAwI5wpMUw"},
	}
	resp, err := service.CorpTagDel(&req)
	if err != nil {
		t.Error("CorpTagDel err：" + err.Error())
		return
	}
	log.Println(resp)
}
