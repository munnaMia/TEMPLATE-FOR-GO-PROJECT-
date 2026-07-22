package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"time"
)

type Header struct {
	ALG string `json:"alg"`
	TYP string `json:"typ"`
}

type Payload struct {
	Sub      string `json:"sub"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	IAT      int64  `json:"iat"`
	EXP      int64  `json:"exp"`
}

// Generate a JWT token.
func GenerateJWT(secretKey string, payload Payload) (string, error) {
	header := Header{
		ALG: "HS256",
		TYP: "JWT",
	}

	now := time.Now().Unix()
	
	if payload.IAT == 0 {
		payload.IAT = now
	}

	if payload.EXP == 0 {
		payload.EXP = now + 86400 // 24 hour of exp time.
	}

	headerBytes, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	headerB64 := base64URLencoding(headerBytes)
	paylaodB64 := base64URLencoding(payloadBytes)

	massage := headerB64 + "." + paylaodB64

	secretKeyBytes := []byte(secretKey)
	massageBytes := []byte(massage)

	hash := hmac.New(sha256.New, secretKeyBytes)
	hash.Write(massageBytes)

	signature := hash.Sum(nil)
	signatureB64 := base64URLencoding(signature)

	jwt := headerB64 + "." + paylaodB64 + "." + signatureB64

	return jwt, nil
}

// encode a byte slice into a base64 string
func base64URLencoding(data []byte) string {
	/*
		URLencoding why?
		Standard Base64 encoding uses the characters + and / as part of its 64-character alphabet. However, those characters have special meanings in URLs (e.g., + means space, / separates path segments).
		.URLEncoding replaces those tricky characters with web-safe alternatives:
		+ becomes - (hyphen)
		/ becomes _ (underscore)
	*/
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
