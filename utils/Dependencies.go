package utils

import "kontest-user-stats-service/service"

type Dependencies struct {
	CodechefService *service.CodeChefService
}

func NewDependencies(codechefService *service.CodeChefService) *Dependencies {
	return &Dependencies{
		CodechefService: codechefService,
	}
}

// Global variable to hold the application dependencies
var dependencies *Dependencies

// InitializeDependencies sets the global dependencies
func InitializeDependencies() {
	dependencies = NewDependencies(service.NewCodeChefService())
}

// GetDependencies returns the global dependencies
func GetDependencies() *Dependencies {
	return dependencies
}
