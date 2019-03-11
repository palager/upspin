// Copyright 2016 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package transports is a helper package that aggregates the user imports.
// It is meant to be imported, using an "underscore" import, as a convenient
// way to link with all the transport implementations.
package transports // import "github.com/palager/upspin/key/transports"

import (
	"github.com/palager/upspin/bind"
	"github.com/palager/upspin/key/inprocess"
	"github.com/palager/upspin/key/usercache"
	"github.com/palager/upspin/upspin"

	_ "github.com/palager/upspin/key/remote"
	_ "github.com/palager/upspin/key/unassigned"
)

func init() {
	bind.RegisterKeyServer(upspin.InProcess, usercache.Global(inprocess.New()))
}
