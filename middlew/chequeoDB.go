package middlew

import (
	"net/http"

	"github.com/kevin1892/twetter/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r) // si no dio error se devuelve la funcion, es como un pasamanos
	}
}
