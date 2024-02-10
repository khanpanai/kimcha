package manager

import (
	"fmt"
	"kimcha/internal/usecase"
	"kimcha/pkg/immu"
	"os"
)

// NewSecretManager basically you don't want to run app without any of managers
func NewSecretManager() usecase.SecretManager {
	uc, err := immu.NewManager()

	if err != nil {
		fmt.Printf("error %s", err)
		os.Exit(1)
	}

	return uc
}

// NewDataManager basically you don't want to run app without any of managers
func NewDataManager() usecase.DataManager {
	uc, err := immu.NewManager()

	if err != nil {
		fmt.Printf("error %s", err)
		os.Exit(1)
	}

	return uc
}

// NewManager basically you don't want to run app without any of managers
func NewManager() usecase.Manager {
	uc, err := immu.NewManager()

	if err != nil {
		fmt.Printf("error %s", err)
		os.Exit(1)
	}

	return uc
}
