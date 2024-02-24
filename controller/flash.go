package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/blue-army-2017/knight/view"
)

const flashCookie = "flash"

func setFlash(w http.ResponseWriter, t string, msg string) {
	cookie := http.Cookie{
		Name:     flashCookie,
		Value:    fmt.Sprintf("%s_%s", t, msg),
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

func getFlash(w http.ResponseWriter, r *http.Request) *view.Flash {
	cookie, _ := r.Cookie(flashCookie)
	if cookie == nil {
		return nil
	}

	cookie.MaxAge = -1
	http.SetCookie(w, cookie)

	vals := strings.SplitN(cookie.Value, "_", 2)
	return &view.Flash{
		Type:    vals[0],
		Message: vals[1],
	}
}
