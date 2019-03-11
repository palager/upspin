// Copyright 2016 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package transports is a helper package that aggregates the store imports.
// It has no functionality itself; it is meant to be imported, using an "underscore"
// import, as a convenient way to link with all the transport implementations.
package transports // import "github.com/palager/upspin/store/transports"

import (
	"github.com/palager/upspin/bind"
	"github.com/palager/upspin/store/inprocess"
	"github.com/palager/upspin/upspin"

	_ "github.com/palager/upspin/store/remote"
	_ "github.com/palager/upspin/store/unassigned"
)

func init() {
	bind.RegisterStoreServer(upspin.InProcess, inprocess.New())
}
