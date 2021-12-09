package middlewares

import (
	"net/http"
	"regexp"
)

var noAuthList []string

func init() {
	noAuthList = make([]string, 0)
}

func AddToNoAuth(path string) {
	if path[len(path)-1:] == "/" {
		path = path + "*"
	}
	noAuthList = append(noAuthList, path)
}

func Redirect() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Context().Value("auth")
			noAuth := isNoAuthPath(r.URL.Path)
			if auth == false && !noAuth {
				if r.Method == "POST" {
					r.Method = "GET"
					http.Redirect(w, r, "/login", http.StatusFound)
					return
				}
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}
			if (auth == true) && noAuth {
				if r.Method == "POST" {
					http.Error(w, "No auth for this page", http.StatusMethodNotAllowed)
					return
				}
				http.Redirect(w, r, "/", http.StatusFound)
			}
			next.ServeHTTP(w, r)
		})
	}
}
func isNoAuthPath(route string) bool {
	// to refactor
	var matches bool
	for _, link := range noAuthList {
		last := matches
		matches, _ = regexp.MatchString(link, route)
		matches = last || matches
	}
	//matches, _ := regexp.Match("^/login.*|^/register.*", []byte(route))
	return matches
}
