// Copyright 2016 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Storeserver is a wrapper for a store implementation that presents it as an
// HTTP interface. It provides the common code used by all storeserver commands.
package storeserver // import "github.com/palager/upspin/serverutil/storeserver"

import (
	"net/http"

	"github.com/palager/upspin/config"
	"github.com/palager/upspin/errors"
	"github.com/palager/upspin/flags"
	"github.com/palager/upspin/log"
	"github.com/palager/upspin/rpc/storeserver"
	"github.com/palager/upspin/serverutil/perm"
	"github.com/palager/upspin/store/inprocess"
	"github.com/palager/upspin/store/server"
	"github.com/palager/upspin/upspin"

	// Directory transports to fetch write permissions.
	_ "github.com/palager/upspin/transports"

	// Packers for reading Access and Group files.
	_ "github.com/palager/upspin/pack/eeintegrity"
	_ "github.com/palager/upspin/pack/plain"
)

func Main() (ready chan<- struct{}) {
	flags.Parse(flags.Server, "kind", "serverconfig")

	// Load configuration and keys for this server. It needs a real upspin username and keys.
	cfg, err := config.FromFile(flags.Config)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new store implementation.
	var store upspin.StoreServer
	err = nil
	switch flags.ServerKind {
	case "inprocess":
		store = inprocess.New()
	case "server":
		store, err = server.New(flags.ServerConfig...)
	default:
		err = errors.Errorf("bad -kind %q", flags.ServerKind)
	}
	if err != nil {
		log.Fatalf("Setting up StoreServer: %v", err)
	}

	// Wrap with permission checks.
	readyCh := make(chan struct{})
	ready = readyCh
	store = perm.WrapStore(cfg, readyCh, store)

	httpStore := storeserver.New(cfg, store, upspin.NetAddr(flags.NetAddr))
	http.Handle("/api/Store/", httpStore)

	return ready
}
