// Copyright 2012 The GAEGo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package user

import (
	"net/http"
)

var LoginURL = "/login"

// LoginRequired is a wrapper for http.HandleFunc. If the requesting
// User is not logged in, they will be redirect to the login page.
func LoginRequired(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if id, _ := CurrentUserID(r); id == "" {
			http.Redirect(w, r, LoginURL, http.StatusFound)
			return
		}
		fn(w, r)
	}
}

// AdminRequired is a wrapper for http.HandleFuc. If the requesting
// User is *not* an admin, they will redirect to the login page.
func AdminRequired(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !CurrentUserHasRole(w, r, "admin") {
			http.Redirect(w, r, LoginURL, http.StatusFound)
			return
		}
		fn(w, r)
	}
}
