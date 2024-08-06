//go:build client
// +build client

// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package client

import (
	_ "github.com/cosmos/gosec/v2/cmd/gosec"
	_ "github.com/ethereum/go-ethereum/cmd/abigen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/onsi/ginkgo/v2/ginkgo"
)
