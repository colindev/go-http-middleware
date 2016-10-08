package middleware

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// CorsMiddleware handle client cors request sets
type CorsMiddleware struct {
	RejectNonCorsRequest          bool
	AllowMethods                  []string
	AllowHeaders                  []string
	AccessControlExposeHeaders    []string
	OriginValidator               func(*http.Request) (string, bool)
	AccessControlMaxAge           int
	AccessControlAllowCredentials bool
}

// Wrap implement Middleware interface
func (cm *CorsMiddleware) Wrap(fn http.HandlerFunc) http.HandlerFunc {

	// TODO 讀取當前request 檔頭設定
	whiteMethods := createWhiteMethods(cm.AllowMethods)
	accessMethodsString := strings.Join(cm.AllowMethods, ",")
	whiteHeaders := createWhiteHeaders(cm.AllowHeaders)
	accessHeadersString := strings.Join(cm.AllowHeaders, ",")

	return func(w http.ResponseWriter, r *http.Request) {

		if cm.RejectNonCorsRequest && !isCors(r) {
			corsReject(w, "Non Cors Request", http.StatusForbidden)
			return
		}

		origin, ok := cm.OriginValidator(r)
		if !ok {
			corsReject(w, "Invalid Origin Request", http.StatusForbidden)
			return
		}

		if isPreflight(r) {
			if !whiteMethods[strings.ToUpper(r.Header.Get("Access-Control-Request-Method"))] {
				corsReject(w, "Invalid Preflight Request", http.StatusForbidden)
				return
			}

			for _, hk := range readAccessHeaders(r) {
				if !whiteHeaders[http.CanonicalHeaderKey(hk)] {
					corsReject(w, "Invalid Preflight Request", http.StatusForbidden)
					return
				}
			}

			w.Header().Set("Access-Control-Request-Methods", accessMethodsString)
			w.Header().Set("Access-Control-Request-Headers", accessHeadersString)
			w.Header().Set("Access-Control-Request-Origin", origin)
			if cm.AccessControlAllowCredentials {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}
			w.Header().Set("Access-Control-Max-Age", strconv.Itoa(cm.AccessControlMaxAge))
			w.WriteHeader(http.StatusOK)
			return
		}

		for _, hk := range cm.AccessControlExposeHeaders {
			w.Header().Add("Access-Control-Expose-Headers", hk)
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		if cm.AccessControlAllowCredentials == true {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		fn(w, r)
	}
}

func readAccessHeaders(r *http.Request) []string {
	accessHeaders := []string{}
	for _, ahs := range r.Header["Access-Control-Request-Headers"] {
		s := strings.Split(ahs, ",")
		accessHeaders = append(accessHeaders, s...)
	}

	return accessHeaders
}

func createWhiteMethods(allows []string) map[string]bool {
	whites := map[string]bool{}

	for _, m := range allows {
		whites[strings.ToUpper(m)] = true
	}

	return whites
}

func createWhiteHeaders(allows []string) map[string]bool {
	whites := map[string]bool{}

	for _, k := range allows {
		whites[http.CanonicalHeaderKey(k)] = true
	}

	return whites
}

func isCors(r *http.Request) bool {
	origin := r.Header.Get("Origin")

	var originURL *url.URL

	if origin == "" {
		return false
	} else if origin == "null" {
		return true
	} else {
		var err error
		originURL, err = url.ParseRequestURI(origin)
		return err == nil && r.Host != originURL.Host
	}
}

func isPreflight(r *http.Request) bool {
	return r.Method == http.MethodOptions
}

func corsReject(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	w.Write([]byte(msg))
}
