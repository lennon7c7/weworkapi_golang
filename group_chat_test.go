package weworkapi_golang_test

import (
	"github.com/go-redis/redis/v8"
	"github.com/lennon7c7/weworkapi_golang"
	"log"
	"testing"
)

func TestGroupChatList(t *testing.T) {
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

	req := weworkapi_golang.GroupChatListReq{
		//StatusFilter: 0,
		OwnerFilter: weworkapi_golang.OwnerFilter{
			UseridList: []string{"OuYangXiongDi"},
		},
		//Cursor:       "",
		Limit: 1000,
	}
	resp, err := service.GroupChatList(&req)
	if err != nil {
		t.Error("GroupChatList err：" + err.Error())
		return
	}
	log.Println(resp)
}

func TestGroupChatGet(t *testing.T) {
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

	req := weworkapi_golang.GroupChatGetReq{
		//ChatID: "wruA1sCQAAo2P_ecegKW_uQyyWkR7cKQ",
		ChatID: "wruA1sCQAATQb6Q82BQy17BX2LOs0Wcw",
		//ChatID: "wruA1sCQAAK7plV1CiOvDDMlxLcXhDqA",
	}
	resp, err := service.GroupChatGet(&req)
	if err != nil {
		t.Error("GroupChatGet err：" + err.Error())
		return
	}
	log.Println(resp)
}
