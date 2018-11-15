package service

import (
    "net/http"
    "github.com/unrolled/render"
)

type User struct {
	Username string
	Password string
}

func tableHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {

		}
		
        formatter.HTML(w, http.StatusOK, "table", struct {
			Username string
            Password string
		}{Username:req.Form.Get("username"), Password:req.Form.Get("password")})
    }
}