// Copyright 2016 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Dirserver is a wrapper for a directory implementation that presents it as an
// HTTP interface.
package main // import "github.com/palager/upspin/cmd/dirserver"

import (
	"github.com/palager/upspin/cloud/https"
	"github.com/palager/upspin/serverutil/dirserver"

	// TODO: Which of these are actually needed?

	// Load useful packers
	_ "github.com/palager/upspin/pack/ee"
	_ "github.com/palager/upspin/pack/eeintegrity"
	_ "github.com/palager/upspin/pack/plain"

	// Load required transports
	_ "github.com/palager/upspin/transports"
)

func main() {
	ready := dirserver.Main()
	https.ListenAndServeFromFlags(ready)
}
