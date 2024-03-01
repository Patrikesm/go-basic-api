package middleware

import "net/http"

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})

	//aqui eu preciso renderizar uma outra função por isso utilizamos um HandlerFunc
	//pra isso criamos uma função anonima que utiliza o response e o request setando o header
	//como content type json
}
