// Copyright (c) ZStack.io, Inc.

package param

type LoginByAccountParam struct {
	BaseParam
	LoginByAccount LoginByAccountDetailParam `json:"logInByAccount"`
}

type LoginByAccountDetailParam struct {
	AccountName string                 `json:"accountName"` // Account name
	Password    string                 `json:"password"`    // Password
	AccountType string                 `json:"accountType"` // Account type
	CaptchaUuid string                 `json:"captchaUuid"` // Captcha UUID
	VerifyCode  string                 `json:"verifyCode"`  // Verification code
	ClientInfo  map[string]interface{} `json:"clientInfo"`  // Client information
}

type LogInByUserParam struct {
	BaseParam
	LogInByUser LogInByUserDetailParam `json:"logInByUser"`
}

type LogInByUserDetailParam struct {
	AccountUuid string                 `json:"accountUuid"` // Account UUID
	AccountName string                 `json:"accountName"` // Account name
	UserName    string                 `json:"userName"`    // User name
	Password    string                 `json:"password"`    // Password
	ClientInfo  map[string]interface{} `json:"clientInfo"`  // Client information
}

type LogInIAM2VirtualIDParam struct {
	BaseParam
	LoginIAM2VirtualID LoginIAM2VirtualIDDetailParam `json:"loginIAM2VirtualID"`
}

type LoginIAM2VirtualIDDetailParam struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LogInIAM2ProjectParam struct {
	BaseParam
	LoginIAM2Project LogInIAM2ProjectDetailParam `json:"loginIAM2Project"`
}

type LogInIAM2ProjectDetailParam struct {
	ProjectName string `json:"projectName"`
}
