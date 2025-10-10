package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/config"
	"encoding/base64"
	"net/http"
	"strings"
)

func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		headerArray := strings.Split(header, " ")

		if len(headerArray) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := headerArray[1]
		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		jwtSignature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload

		byteArraySecret := []byte(config.GetConfig().JwtSecretKey)
		byteArrayMessage := []byte(message)

		h := hmac.New(sha256.New, byteArraySecret)
		h.Write(byteArrayMessage)

		hash := h.Sum(nil)
		newSinature := base64UrlEncode(hash)

		if newSinature != jwtSignature {
			http.Error(w, "Unauthorized, might be hacker", http.StatusUnauthorized)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

func base64UrlEncode(data []byte) string {
		return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}