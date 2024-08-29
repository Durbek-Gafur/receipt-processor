// internal/server/wire.go
//go:build wireinject

package server

import (
	db "receipt_processor/internal/dbprovider"
	"receipt_processor/internal/logic"

	"github.com/google/wire"
)

// ProviderSet is the Wire provider set for the server.
var ProviderSet = wire.NewSet(
	db.ProvideCacheDB,
	logic.ProvideOptions,
	logic.NewReceiptLogic,
	NewServer,

)

// InitializeServer initializes a new server instance.
func InitializeServer() (Server, error) {
	wire.Build(ProviderSet)
	return nil, nil
}
