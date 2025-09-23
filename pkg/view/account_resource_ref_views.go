package view

type AccountResourceRefInventoryView struct {
	BaseTimeView
	AccountUuid          string `json:"accountUuid"`
	ConcreteResourceType string `json:"concreteResourceType"`
	ResourceType         string `json:"resourceType"`
	ResourceUuid         string `json:"resourceUuid"`
}

type IMA2ProjectInventoryView struct {
	BaseInfoView
	BaseTimeView
	LinkedAccountUuid string `json:"linkedAccountUuid"`
	State             string `json:"state"`
}

type AccountInventoryView struct {
	BaseTimeView
	Name string `json:"name"`
	Type string `json:"type"`
	UUID string `json:"uuid"`
}
