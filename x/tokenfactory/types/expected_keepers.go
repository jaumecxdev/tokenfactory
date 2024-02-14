package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccountKeeper defines the expected interface for the Account module.
// /home/jaume/go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.3/x/bank/types/expected_keepers.go
type AccountKeeper interface {
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	//GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
	// Methods imported from account should be defined here
	GetModuleAddress(moduleName string) sdk.AccAddress
	//GetModuleAccount(ctx sdk.Context, moduleName string) sdk.ModuleAccountI
	GetModuleAccount(ctx context.Context, moduleName string) sdk.ModuleAccountI
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	//SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
	//SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoins(ctx context.Context, from, to sdk.AccAddress, amt sdk.Coins) error
	//MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}
