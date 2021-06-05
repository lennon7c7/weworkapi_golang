package weworkapi_golang_test

import (
	"github.com/go-redis/redis/v8"
	"github.com/lennon7c7/weworkapi_golang"
	"log"
	"testing"
)

func TestGroupMsgList(t *testing.T) {
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

	// 企业群发 给 客户
	req := weworkapi_golang.GroupMsgListReq{
		ChatType: "single",
		//StartTime: 1622881345,
		//EndTime:   1622908799,
		Creator: "OuYangXiongDi",
		//FilterType: 0,
		//Limit:      100,
		//Cursor:     "",
	}
	// 企业群发 给 客户群
	//req = weworkapi_golang.GroupMsgListReq{
	//	ChatType:  "group",
	//	//StartTime: 1622881345,
	//	//EndTime:   1622908799,
	//	Creator:   "OuYangXiongDi",
	//	//FilterType: 0,
	//	Limit:      100,
	//	//Cursor:     "",
	//}
	resp, err := service.GroupMsgList(&req)
	if err != nil {
		t.Error("GroupMsgList err：" + err.Error())
		return
	}
	log.Println(resp)
}

func TestGroupMsgTaskList(t *testing.T) {
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

	req := weworkapi_golang.GroupMsgTaskListReq{
		Msgid: "msguA1sCQAAfD4CYq3uMABb_A6HumAbog",
		//Limit:  0,
		//Cursor: "",
	}
	resp, err := service.GroupMsgTaskList(&req)
	if err != nil {
		t.Error("GroupMsgTaskList err：" + err.Error())
		return
	}
	log.Println(resp)
}

func TestGroupMsgResultList(t *testing.T) {
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

	req := weworkapi_golang.GroupMsgResultListReq{
		Msgid: "msguA1sCQAAfD4CYq3uMABb_A6HumAbog",
	}
	resp, err := service.GroupMsgResultList(&req)
	if err != nil {
		t.Error("GroupMsgResultList err：" + err.Error())
		return
	}
	log.Println(resp)
}

func TestGroupMsgSendResultList(t *testing.T) {
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

	req := weworkapi_golang.GroupMsgSendResultListReq{
		Msgid:  "msguA1sCQAAfD4CYq3uMABb_A6HumAbog",
		Userid: "OuYangXiongDi",
		//Limit:  0,
		//Cursor: "",
	}
	resp, err := service.GroupMsgSendResultList(&req)
	if err != nil {
		t.Error("GroupMsgSendResultList err：" + err.Error())
		return
	}
	log.Println(resp)
}

func TestGroupMsgAdd(t *testing.T) {
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

	// 创建 企业群发 给 客户
	req := weworkapi_golang.GroupMsgAddReq{
		ChatType: "single",
		ExternalUserid: []string{
			"wmuA1sCQAAUsh9IpSJeSRWweYqs2ekpA",
			"wmuA1sCQAAdGyPSa04j8RJgsP71TV86Q",
			"wmuA1sCQAAFvSnuhRsUbc5JZH7aqZmFA",
		},
		Sender: "OuYangXiongDi",
	}
	req.Text = &weworkapi_golang.GroupWelcomeTemplateText{Content: "稻花香里说丰年。听取蛙声一片。"}
	// 创建 企业群发 给 客户群
	//req = weworkapi_golang.GroupMsgAddReq{
	//	ChatType: "group",
	//	Text: struct {
	//		Content string `json:"content"`
	//	}{Content: "峨眉山月半轮秋，影入平羌江水流。 夜发清溪向三峡，思君不见下渝州。"},
	//	Sender: "OuYangXiongDi",
	//}
	//req = weworkapi_golang.GroupMsgAddReq{
	//	ChatType: "group",
	//	Text: struct {
	//		Content string `json:"content"`
	//	}{Content: "山有木兮木有枝，心悦君兮君不知。——佚名《越人歌》"},
	//	Sender: "OuYangXiongDi",
	//}
	resp, err := service.GroupMsgAdd(&req)
	if err != nil {
		t.Error("GroupMsgAdd err：" + err.Error())
		return
	}
	log.Println(resp)
}
