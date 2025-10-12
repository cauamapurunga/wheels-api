package config

import "os"

func GetJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return []byte("your_secret_key")
	}
	return []byte(secret)
}
