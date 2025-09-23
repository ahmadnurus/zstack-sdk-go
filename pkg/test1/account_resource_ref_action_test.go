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
	fmt.Println("Query All===============================")
	accountResources, err := accountLoginCli.QueryAccountResourceRef(queryParam)
	if err != nil {
		golog.Errorf("TestAccountResourceRef %v", err)
	}

	fmt.Printf("data : %v", accountResources)

	//条件查询
	fmt.Println("Query condition===============================")
	queryParam.Set("q", "accountUuid=ab9d2b974f0d445781197ee2d15d3f27")
	accountResources, err = accountLoginCli.QueryAccountResourceRef(queryParam)
	if err != nil {
		golog.Errorf("TestAccountResourceRef by AccontUuid %v", err)
	}
	fmt.Printf("data : %v", accountResources)
}

func TestQueryAccount(t *testing.T) {
	accountLoginCli.Login()
	defer accountLoginCli.Logout()
	queryParam := param.NewQueryParam()
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
