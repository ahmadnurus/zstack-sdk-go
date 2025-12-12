package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// PageVmInstance Paginate VM instances
func (cli *ZSClient) QueryAccountResourceRef(params param.QueryParam) ([]view.AccountResourceRefInventoryView, error) {
	var resp []view.AccountResourceRefInventoryView

	return resp, cli.List("v1/accounts/resources/refs", &params, &resp)
}

func (cli *ZSClient) QuerryIAMProject(param param.QueryParam) ([]view.IMA2ProjectInventoryView, error) {
	var resp []view.IMA2ProjectInventoryView
	return resp, cli.List("v1/iam2/projects", &param, &resp)
}

func (cli *ZSClient) QueryAccount(param param.QueryParam) ([]view.AccountInventoryView, error) {
	var resp []view.AccountInventoryView
	return resp, cli.List("v1/accounts", &param, &resp)
}
