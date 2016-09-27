package documentdb

import "github.com/jen20/riviera/azure"

type GetDatabaseResponse struct {
	ID   *string `mapstructure:"id"`
	Name *string `mapstructure:"name"`
	Type *string `mapstructure:"type"`
}

type GetDocumentDBDatabase struct {
	DatabaseName      string             `json:"-"`
	ResourceGroupName string             `json:"-"`
	FailoverPolicies  []FailoverPolicies `json:"failoverPolicies"`
}

type FailoverPolicies struct {
	LocationName     string `json:"locationName"`
	FailoverPriority int    `json:"failoverPriority"`
}

func (s GetDocumentDBDatabase) APIInfo() azure.APIInfo {
	return azure.APIInfo{
		APIVersion:  apiVersion,
		Method:      "POST",
		URLPathFunc: documentDBDatabasePath(s.ResourceGroupName, s.DatabaseName),
		ResponseTypeFunc: func() interface{} {
			return &GetDatabaseResponse{}
		},
	}
}
