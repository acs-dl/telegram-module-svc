package helpers

import (
	"net/http"

	"github.com/acs-dl/auth-svc/internal/data"
)

func SetTokensCookies(w http.ResponseWriter, access, refresh string) {
	refreshCookie := &http.Cookie{
		Name:     data.RefreshCookie,
		Value:    refresh,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, refreshCookie)

	accessCookie := &http.Cookie{
		Name:     data.AccessCookie,
		Value:    access,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, accessCookie)
}

func ClearTokensCookies(w http.ResponseWriter) {
	refreshCookie := &http.Cookie{
		Name:     data.RefreshCookie,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, refreshCookie)

	accessCookie := &http.Cookie{
		Name:     data.AccessCookie,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, accessCookie)
}
