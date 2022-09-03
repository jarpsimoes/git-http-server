package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type CustomPath struct {
	path    string
	target  string
	rewrite bool
}

// GetPath - return path value
func (cp CustomPath) GetPath() string {
	return cp.path
}

// GetTarget - return target value
func (cp CustomPath) GetTarget() string {
	return cp.target
}

// IsRewrite - return true if rewrite is available
func (cp CustomPath) IsRewrite() bool {
	return cp.rewrite
}

var baseCustomPathsInstance *[]CustomPath

// GetAllCustomPaths it's a function to get all custom paths
// started with GHS_CUSTOM_PATH_
// Recommended to be used on redirects for cloud storages
func GetAllCustomPaths() *[]CustomPath {
	envVars := os.Environ()
	var result []CustomPath

	for index, variable := range envVars {
		if strings.HasPrefix(variable, "GHS_CUSTOM_PATH_") {
			log.Println(fmt.Sprintf("I: %X - V: %s", index, variable))
			splitVar := strings.Split(variable, "=")

			newPath := strings.ReplaceAll(splitVar[0], "GHS_CUSTOM_PATH_", "")

			if strings.HasPrefix(newPath, "/") {
				newPath = newPath[1:len(newPath)]
			}
			if len(splitVar) > 1 {
				varObject := CustomPath{
					path:    strings.ReplaceAll(newPath, ".", "/"),
					target:  splitVar[1],
					rewrite: checkIfEnvVariableIsSet(envVars, fmt.Sprintf("GHS_CUSTOM_REWRITE_%s", newPath)),
				}

				result = append(result, varObject)
			}
		}

	}

	return &result
}
func checkIfEnvVariableIsSet(envVars []string, envVar string) bool {

	for _, variable := range envVars {
		if strings.HasPrefix(variable, envVar) {
			return true
		}
	}

	return false

}

// FindPath - Find path on custom path instance
func FindPath(pathRequested string) (bool, *CustomPath) {

	paths := GetCustomPathsInstance()

	for _, path := range *paths {
		prefixPath := fmt.Sprintf("/%s", path.GetPath())
		if strings.HasPrefix(pathRequested, prefixPath) {
			return true, &path
		}
	}

	return false, nil

}
func GetCustomPathsInstance() *[]CustomPath {
	if baseCustomPathsInstance == nil {

		lock.Lock()
		defer lock.Unlock()

		if baseCustomPathsInstance == nil {
			log.Println("[CustomPaths] Creating new instance")

			baseCustomPathsInstance = GetAllCustomPaths()
		}
	}

	return baseCustomPathsInstance
}
