package middleware

import (
	"log"
	"net/http"
	"strings"
)

type Middleware func(http.Handler) http.Handler


func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Do stuff here
        log.Println(r.Method," ",r.Host," ",r.RequestURI," ",r.RemoteAddr)
        // Call the next handler, which can be another middleware in the chain, or the final handler.
        next.ServeHTTP(w, r)
    })
}

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		authorizationHeader:=r.Header.Get("Authorization")		
		if authorizationHeader != "" {
			token:=strings.Split(authorizationHeader,"Bearer ")[1]
			if(len(token)>0){
				log.Println("User withjwt token")
			}
		}else{
			log.Println("User anonymous")
		}
	// Call the next handler, which can be another middleware in the chain, or the final handler.
	next.ServeHTTP(w, r)
	})
}