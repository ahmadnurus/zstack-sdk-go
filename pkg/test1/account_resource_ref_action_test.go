package test

import (
	"fmt"
	"testing"

	"github.com/kataras/golog"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
)

func TestQueryAccountResourceRef(t *testing.T) {
	accountLoginCli.Login()
	defer accountLoginCli.Logout()
	queryParam := param.NewQueryParam()
	queryParam.AddQ("resourceType=VmInstanceVO")
	fmt.Println("Query All===============================")
	accountResources, err := accountLoginCli.QueryAccountResourceRef(queryParam)

	if err != nil {
		golog.Errorf("TestAccountResourceRef %v", err)
	}

	//fmt.Printf("data : %v", accountResources)
	for i, aaccountResources := range accountResources {
		/*
			if aaccountResources.ResourceType == "VmInstanceVO" {
				fmt.Printf("test pass\n")
				//fmt.Printf("i is %d, accountResources is %v\n", i, aaccountResources.ResourceType)
			}
		*/
		fmt.Printf("i is %d, accountResources is %v\n", i, aaccountResources.ResourceType)

	}

	//条件查询
	//fmt.Println("Query condition===============================")
	//queryParam.AddQ("accountUuid=ab9d2b974f0d445781197ee2d15d3f27")

	//queryParam.AddQ("lastOpDate>2025-08-07T13:51:51")

	//accountResources, err = accountLoginCli.QueryAccountResourceRef(queryParam)
	//if err != nil {
	//	golog.Errorf("TestAccountResourceRef by AccontUuid %v", err)
	//}
	//fmt.Printf("data : %v", accountResources)
}

func TestQueryAccount(t *testing.T) {
	accountLoginCli.Login()
	defer accountLoginCli.Logout()
	queryParam := param.NewQueryParam()
	//queryParam.AddQ("lastOpDate='2025-09-22 13:51:51'")
	fmt.Println("Query All===============================")
	account, err := accountLoginCli.QueryAccount(queryParam)
	if err != nil {
		golog.Errorf("TestAccount %v", err)
	}
	fmt.Printf("account data: %v", account)
}

func TestQueryIAM2Prject(t *testing.T) {

	accountLoginCli.Login()
	defer accountLoginCli.Logout()
	queryParam := param.NewQueryParam()
	fmt.Println("Query All===============================")
	account, err := accountLoginCli.QuerryIAMProject(queryParam)
	if err != nil {
		golog.Errorf("TestQueryIAM2Prject %v", err)
	}
	fmt.Printf("account data: %v", account)
}
