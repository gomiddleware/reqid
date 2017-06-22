// Copyright (c) 2016, Apps Attic Ltd (https://appsattic.com/) <chilts@appsattic.com>.

// Permission to use, copy, modify, and/or distribute this software for any purpose with or without fee is hereby
// granted, provided that the above copyright notice and this permission notice appear in all copies.

// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH REGARD TO THIS SOFTWARE INCLUDING ALL
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN
// AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
// PERFORMANCE OF THIS SOFTWARE.

/*
Package reqid provides standard Go middleware to read the X-Request-ID header on the incoming request and store it in
the request's context. If no header exists, then one is generated. You may optionally scrub the incoming header so that
one is always generated (such as on externally facing servers).

Also provided are some convenience functions to be able to read this RequestId back out in later middleware.

Like every other https://gomiddleware.github.io/ middleware, it uses both the http.Handler interface and the context
package to store each request's ID. Since Go v1.7, this has been in the Go standard library and therefore we should
reduce our dependence on external packages such as gorilla/context.

All of the functions such as RandomId, exported by this package work in the same way as noted above. Explicitly:

* if the incoming request already has a X-Request-Id header, it is read and stored in the request's context
* if there is no X-Request-Id header, then one is generated and stored in the request's context (as above)

See https://golang.org/pkg/context/ for more information on the context package. And see
https://blog.golang.org/context for a run-down of some code that uses Contexts.

*/
package reqid
