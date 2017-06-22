// Copyright (c) 2016, Apps Attic Ltd (https://appsattic.com/) <chilts@appsattic.com>.

// Permission to use, copy, modify, and/or distribute this software for any purpose with or without fee is hereby
// granted, provided that the above copyright notice and this permission notice appear in all copies.

// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH REGARD TO THIS SOFTWARE INCLUDING ALL
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN
// AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
// PERFORMANCE OF THIS SOFTWARE.

package reqid

import (
	"context"
	"net/http"

	"github.com/chilts/sid"
)

type key int

const reqIdKey key = 42

// randomId just returns a string of a new Id.
func randomId(len int) string {
	return sid.Id()
}

// ScrubRequestIdHeader should be used on externally facing servers when you want to add your own X-Request-ID. Use
// this before RandomId. If an externally facing server is hitting internal microservices and you want the X-Request-ID
// to be passed along to the other services, then just use the RandomId since it will keep any existing value.
func ScrubRequestIdHeader(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r.Header.Del("X-Request-ID")
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// RandomId will read the incoming X-Request-ID header and store it in the request's context and can be used in later
// middleware. If the incoming request does not have an associated X-Request-ID, then one will be generated instead.
//
// Note: for internal services, it is safe to just use this middleware. However, for externally facing servers you must
// use ScrubRequestIdHeader first so that external clients can't set something which is an internal artifact of your
// services.
func RandomId(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		reqId := r.Header.Get("X-Request-ID")
		if reqId == "" {
			reqId = randomId(12)
		}

		// get a new context and pass to the next middleware using it
		ctx := context.WithValue(r.Context(), reqIdKey, reqId)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

// ReqIdFromRequest can be used to obtain the ReqId from the request. This is given as a convenience method, though at
// the same time, you must use it because the key value used to store the ReqId in the context is not exported from
// this package.
func ReqIdFromRequest(r *http.Request) string {
	return r.Context().Value(reqIdKey).(string)
}

// ReqIdFromContext can be used to obtain the ReqId from the context (if you already have it handy). You are most
// likely to use ReqIdFromRequest rather than this, but either does the same job.
func ReqIdFromContext(ctx context.Context) string {
	return ctx.Value(reqIdKey).(string)
}
