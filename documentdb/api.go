package documentdb

import "fmt"

const apiVersion = "2015-04-08"
const apiProvider = "Microsoft.DocumentDB"

func documentDBPath(resourceGroupName, dbName string) func() string {
	return func() string {
		return fmt.Sprintf("resourceGroups/%s/providers/%s/databaseAccounts/%s", resourceGroupName, apiProvider, dbName)
	}
}

func documentDBFailoverPath(resourceGroupName, dbName string) func() string {
	return func() string {
		return fmt.Sprintf("resourceGroups/%s/providers/%s/databaseAccounts/%s/failoverPriorityChange", resourceGroupName, apiProvider, dbName)
	}
}
