// Copyright 2018 NTT Group

// Permission is hereby granted, free of charge, to any person obtaining a copy of this
// software and associated documentation files (the "Software"), to deal in the Software
// without restriction, including without limitation the rights to use, copy, modify,
// merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to the following
// conditions:

// The above copyright notice and this permission notice shall be included in all copies
// or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
// INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
// PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE
// FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package main

import (
	"net/http"

	anki "github.com/okoeth/edge-anki-base"
	"goji.io"
	"goji.io/pat"
)

type (
	// BeatController represents the controller for working with this app
	BeatController struct {
		cmdCh chan anki.Command
	}
)

// NewBeatController provides a reference to an IncomingController
func NewBeatController(ch chan anki.Command) *BeatController {
	return &BeatController{cmdCh: ch}
}

// AddHandlers inserts new greeting
func (oc *BeatController) AddHandlers(mux *goji.Mux) {
	mux.HandleFunc(pat.Get("/v1/beat/status"), oc.GetStatus)
}

// GetStatus retrieves latest status
func (oc *BeatController) GetStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status is ok"))
}
