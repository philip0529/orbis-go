package main

import (
	"context"
	"fmt"

	"github.com/sourcenetwork/orbis-go/app"
	"github.com/sourcenetwork/orbis-go/config"
	p2pbb "github.com/sourcenetwork/orbis-go/pkg/bulletin/p2p"
	"github.com/sourcenetwork/orbis-go/pkg/dkg/rabin"
	"github.com/sourcenetwork/orbis-go/pkg/host"
	"github.com/sourcenetwork/orbis-go/pkg/pre/elgamal"
	"github.com/sourcenetwork/orbis-go/pkg/pss/avpss"
	p2ptp "github.com/sourcenetwork/orbis-go/pkg/transport/p2p"
)

func setupApp(ctx context.Context, cfg config.Config) (*app.App, error) {

	// testingStreamHandler := func(msg *transport.Message) error {
	// 	log.Infof("Transport Received message: %s", msg.Payload)
	// 	return nil
	// }
	// tp.AddHandler(p2ptp.ProtocolID, testingStreamHandler)

	opts := []app.Option{
		app.DefaultOptions(),
		app.WithHost(host.Factory{}),
		app.WithTransport(p2ptp.Factory),
		app.WithBulletin(p2pbb.Factory),

		app.WithDistKeyGenerator(rabin.Factory),
		app.WithProxyReencryption(elgamal.Factory),

		// Enable support the AVPSS, ECPSS, and CHURP based PSS systems.
		app.WithProactiveSecretSharing(avpss.Factory),

		// Also enable basic VSS for testing (no networking/bulleting required).
		// app.WithProactiveSecretSharing(vss.Provider),

		// mount DB Tables
		// app.WithTable()
	}

	app, err := app.New(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("create app: %w", err)
	}

	return app, nil
}

// 	rid := types.RingID("40b086ef") // cid/multihash of the ring config
// 	err = node.JoinRing(rid, config.Ring, ringPeers)

// 	ring := node.GetRing(rid)

// 	select {
// 	case <-time.NewTimer(time.Minute):
// 		// timeout
// 	case <-ring.WaitForState(ptypes.STATE_INITIALIZED):
// 		// ready
// 	}

// 	sid = types.SecretID("mySecretIdentifier")
// 	secretVal := []byte("My Super Secret Value or Private Key or Symmetric Key or w.e")
// 	ring.Store(ctx, sid, secretVal)

// 	secret, err := ring.Get(ctx, auth.NilToken, sid)
// 	share, err := ring.GetLocalShare(ctx, auth.NilToken, sid)
// 	shares, err := ring.GetShares(ctx, auth.NilToken, sid)

// 	recovered := orbis.RecoverFromShares(shares)
