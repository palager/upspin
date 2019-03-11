// Copyright 2016 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Storeserver is a wrapper for a store implementation that presents it as an
// HTTP interface.
package main // import "github.com/palager/upspin/cmd/storeserver"

import (
	"github.com/palager/upspin/cloud/https"
	"github.com/palager/upspin/serverutil/storeserver"

	// Storage implementation.
	_ "github.com/palager/upspin/cloud/storage/disk"
)

func main() {
	ready := storeserver.Main()
	https.ListenAndServeFromFlags(ready)
}
