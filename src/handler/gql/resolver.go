package gql

import "github.com/irdaislakhuafa/octacat-app-backend/src/business/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Usecase *usecase.Usecase
}
