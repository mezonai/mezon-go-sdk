package utils

import "fmt"

func GetBasePath(schema, basePath string, useSSL bool) string {
	if useSSL {
		return fmt.Sprintf("%ss://%s", schema, basePath)
	}
	return fmt.Sprintf("%s://%s", schema, basePath)
}
