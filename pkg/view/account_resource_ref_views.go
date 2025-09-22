package view

type AccountResourceRefInventoryView struct {
	BaseTimeView
	AccountUuid          string `json:"accountUuid"`
	ConcreteResourceType string `json:"concreteResourceType"`
	ResourceType         string `json:"resourceType"`
	ResourceUuid         string `json:"resourceUuid"`
}
