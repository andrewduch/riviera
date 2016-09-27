package documentdb

import "fmt"

const apiVersion = "2015-04-08"
const apiProvider = "Microsoft.DocumentDB"

func documentDBDatabasePath(resourceGroupName, dbName string) func() string {
	return func() string {
		return fmt.Sprintf("resourceGroups/%s/providers/%s/databaseAccounts/%s", resourceGroupName, apiProvider, dbName)
	}
}
