// Copyright 2016 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package transports is a helper package that aggregates
// the key, store, and directory imports.
// It can be imported by Upspin programs as a convenient
// way to link with all the transport implementations.
package transports // import "github.com/palager/upspin/transports"

import (
	"sync"

	"github.com/palager/upspin/bind"
	"github.com/palager/upspin/dir/inprocess"
	"github.com/palager/upspin/upspin"

	_ "github.com/palager/upspin/key/transports"
	_ "github.com/palager/upspin/store/transports"

	_ "github.com/palager/upspin/dir/remote"
	_ "github.com/palager/upspin/dir/unassigned"
)

var bindOnce sync.Once

// Init initializes the transports for the given configuration.
// It is a no-op if passed a nil config or called more than once.
//
// It should be called only by client programs, directly after parsing a
// config. This handles the case where a config specifies an inprocess
// directory server and configures that server to talk to the specified store
// server.
func Init(cfg upspin.Config) {
	if cfg == nil {
		return
	}
	if cfg.DirEndpoint().Transport == upspin.InProcess {
		bindOnce.Do(func() {
			bind.RegisterDirServer(upspin.InProcess, inprocess.New(cfg))
		})
	}
}
