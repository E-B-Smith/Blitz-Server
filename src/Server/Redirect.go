//  Redirect.go  -  Handle Redirects social data.
//
//  E.B.Smith  -  May, 2015.


package main


import (
    "fmt"
    "path"
    "html"
    "strings"
    "net/url"
    "net/http"
)

// Redirect replies to the request with a redirect to url,
// which may be a path relative to the request path.


func RedirectWithQuery(w http.ResponseWriter, r *http.Request, urlStr string, code int) {
    if u, err := url.Parse(urlStr); err == nil {
        // If url was relative, make absolute by
        // combining with request path.
        // The browser would probably do this for us,
        // but doing it ourselves is more reliable.

        // NOTE(rsc): RFC 2616 says that the Location
        // line must be an absolute URI, like
        // "http://www.google.com/redirect/",
        // not a path like "/redirect/".
        // Unfortunately, we don't know what to
        // put in the host name section to get the
        // client to connect to us again, so we can't
        // know the right absolute URI to send back.
        // Because of this problem, no one pays attention
        // to the RFC; they all send back just a new path.
        // So do we.
        oldpath := r.URL.Path
        if oldpath == "" { // should not happen, but avoid a crash if it does
            oldpath = "/"
        }
        if u.Scheme == "" {
            // no leading http://server
            if urlStr == "" || urlStr[0] != '/' {
                // make relative path absolute
                olddir, _ := path.Split(oldpath)
                urlStr = olddir + urlStr
            }

            var query string
            if i := strings.Index(urlStr, "?"); i != -1 {
                urlStr, query = urlStr[:i], urlStr[i:]
            }

            if len(r.URL.RawQuery) > 0 {
                query = "?"+r.URL.RawQuery
            }

            // clean up but preserve trailing slash
            trailing := strings.HasSuffix(urlStr, "/")
            urlStr = path.Clean(urlStr)
            if trailing && !strings.HasSuffix(urlStr, "/") {
                urlStr += "/"
            }
            urlStr += query
        }
    }

    w.Header().Set("Location", urlStr)
    w.WriteHeader(code)

    // RFC2616 recommends that a short note "SHOULD" be included in the
    // response because older user agents may not understand 301/307.
    // Shouldn't send the response for POST or HEAD; that leaves GET.
    if r.Method == "GET" {
        note := "<a href=\"" + html.EscapeString(urlStr) + "\">" + http.StatusText(code) + "</a>.\n"
        fmt.Fprintln(w, note)
    }
}


// Redirect to a fixed URL
type redirectWithQueryHandler struct {
    url  string
    code int
}

func (rh *redirectWithQueryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    RedirectWithQuery(w, r, rh.url, rh.code)
}

// RedirectHandler returns a request handler that redirects
// each request it receives to the given url using the given
// status code.
func RedirectWithQueryHandler(url string, code int) http.Handler {
    return &redirectWithQueryHandler{url, code}
}

