package weworkapi_golang_test

import (
	"github.com/go-redis/redis/v8"
	"github.com/lennon7c7/weworkapi_golang"
	"log"
	"testing"
)

// 获取素材的下载链接
func TestMediaGetDownloadURLByMediaID(t *testing.T) {
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

	req := weworkapi_golang.MediaGetReq{
		//MediaID: "WWCISP_7e1I_242Jt5ssnMYeM842N703PkAVR2KO8DVphB1M87ghFpLvRbDZN17ioZFZnrWV9I0bPpEDnZ_CbG9ylPB2pVRSexzsz2uNcXzWM_-zmHfH85tNqplriQ3Y3O5s4VDccSSq11U6BJ2yKOhfz1pBUAbIjVLgVxDDBY-9aD-2h06M_hhDDytUFvOR1N38hfFSvcWcncqQq3HsUXTojIVMviu5kbQWE2r8lkkPvSpE4Rgfr0Y_hMgHMpcCyrJJtAqWU3inCefsusv8lKM6GIszS9lgGcPhfsRV1kEnOnxHMPreS1FmL_xpuZhXreYo_bc3bMn43cj1cvUUHIbnYWjHpuPvvOe-DDcSGabNzLEyqRXRcT0MdM08-PG090kFkjINgsFgqsX54Yj0n0L1_sVb2E5CvyaBjLznqkEDaPeIvfQAw-a3wmlK_4fYFuEY9GRht_X3t2rOWqpGWQjRFV4QjBrDqHUjCqcooXICUJJFbSr_0rwvp4L1SBVkw4x5mUn1A7EE0DT35oE_qOxI-RBoYo2l4jdTj5Fwq_9m0Nt_MDaAys8p-WjptSkLp-NmOSx61XCPytzdnezrtIpUY340sIakJna6CJ2vB31DhzQFOzpoBnzFH-PfoYD-f8IBsDNpCFnBe6uNwbH75AaFox3iT9jq6RDjLPxao0h8nILCBIJ9_sTFJAtEHQMTQDH5gp0nGV2CgIRtzBkxXds089BfPNloVIvSHlrFg6wi8V1OcY",
		MediaID: "WWCISP_7e1I_242Jt5ssnMYeM842N703PkAVR2KO8DVphB1M87ghFpLvRbDZN17ioZFZnrWV9I0bPpEDnZ_CbG9ylPB2pVRSexzsz2uNcXzWM_-zmEq0dpWn7BR3vG5uk3G73uW2NV83NrvNhZVKAV-_2iM-8cjIKrrvFf42FYvh2VRVa3pAOXXCgs5-4JlA6RKNKretNs333zj9IicUydlhVWpnHRzynBWpFuC4zNaPrQL3zxNOi1i24seWhMbcNkw2eLNdI8tzzGuUI05r1h9GaMHFfIJmpak3S5T4POYgrpm85cpIPQX6fwk3_-RNcVcg3rAHW6nimsDf43IiQqsPKgPjp2X5citxwMTMs82IZ_JEcbSyTgrCps-r8-qeyJUtvOIYK9FXqzIGGSehXjIDVKtIbkOMjHZmjkBEp48r_mbTVu4LAMgrhxb8DkqaRy_zOh7kRWLN8dusHCgOTpk4BYaabRXXoVglnLz_xUx6eGENNjhJVil7MTl6A7szkx0CwHhzfk1rKe1ewcTIcpejjfy5zdMUb0Fq94CIfxD-vDbK0IVjTW1W1QR4Ob1JZKTY-WjR4ZF32R3VpT0Jt9NsdwjFEl30EEhN67eMyng1UCpYcKoa_p56SmjWvbgKixjqTRhv10e0gfa7DF57gnXTksFmXDpYtiv5TZIlEYgGIdX6r2B5INx2sHAfhX2yRYSuAXC6YEVzF0fglBmDOhcHUJosFp_riXrSLkUYOdKsrM12Rw",
	}
	resp, err := service.MediaGetDownloadURLByMediaID(&req)
	if err != nil {
		t.Error("MediaGet err：" + err.Error())
		return
	}
	log.Println(resp)
}

func TestMediaUploadImg(t *testing.T) {
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

	req := weworkapi_golang.MediaReq{
		Filename: "444.jpg",
	}
	resp, err := service.MediaUploadImg(&req)
	if err != nil {
		t.Error("MediaUploadImg err：" + err.Error())
		return
	}
	log.Println(resp)
}
