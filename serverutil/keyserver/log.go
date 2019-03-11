// Copyright 2016 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package keyserver

import (
	"net/http"

	"github.com/palager/upspin/key/server"
	"github.com/palager/upspin/log"
)

type logHandler struct {
	logger server.Logger
}

func (h logHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := h.logger.Log()
	if err != nil {
		log.Error.Printf("logHandler: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
