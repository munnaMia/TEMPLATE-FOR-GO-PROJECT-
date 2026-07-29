package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
)

func (mdlw *Middleware) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// parse JWT
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Extract header payload
		headerArr := strings.Split(header, " ")

		if len(headerArr) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		accessToken := headerArr[1] // secound elem contain the jwt

		// now split the token
		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		jwtSignature := tokenParts[2]

		// using Hmac sha256 -> generate signeture using env secret key
		massage := jwtHeader + "." + jwtPayload

		byteArrSecret := []byte(mdlw.cnf.Service.SecretKey) // convet secret also base64.
		byteArrMsg := []byte(massage)

		h := hmac.New(sha256.New, byteArrSecret) // algorithm & secret key as byte slice
		h.Write(byteArrMsg)

		hash := h.Sum(nil)

		newSignature := base64UrlEncode(hash)

		if jwtSignature != newSignature {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// match user jwt with our generated signeture and validate user
		// if valid user allow to create product
		// else unauthorized 401 error send

		next.ServeHTTP(w, r)
	})

}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
