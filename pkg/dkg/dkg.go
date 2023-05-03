package dkg

import (
	"context"

	"github.com/samber/do"
	"github.com/sourcenetwork/orbis-go/pkg/bulletin"
	"github.com/sourcenetwork/orbis-go/pkg/crypto"
	"github.com/sourcenetwork/orbis-go/pkg/db"
	"github.com/sourcenetwork/orbis-go/pkg/transport"
)

const (
	ProtocolName = "dkg"
)

const (
	INITIALIZED State = iota // DKG group has initialized but not started the generation
	STARTED                  // Started the distributed key generation
	CERTIFIED                // Generated and cerified the shared key
)

// enum
type State uint8

type Node = transport.Node

type DKG interface {
	Name() string

	Init(ctx context.Context, nodes []transport.Node, n int, threshold int) error

	PublicKey() crypto.PublicKey
	Share() crypto.PriShare

	State() State

	Start(context.Context) error
	Close(context.Context) error

	ProcessMessage(*transport.Message) error

	// hooks?
}

type Factory interface {
	New(db.Repository, transport.Transport, bulletin.Bulletin, crypto.PrivateKey) (DKG, error)
	Name() string
}

type ProviderFactory func(*do.Injector) Factory
