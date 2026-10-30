package credentials

import (
	"crypto/rand"
	"encoding/hex"
	"goslug/internal/logger"
	"io"
	"sync"
	"time"
)

var AUTH_USERNAME string
var AUTH_PASSWORD string

var activeSessions map[string]time.Time = make(map[string]time.Time, 5)
var mu sync.RWMutex

const sessionDuration = 48 * time.Hour

// Checks user and string against sessions database
func Authenticate(user string, pass string) (string, bool) {
	mu.Lock()
	defer mu.Unlock()

	if user != AUTH_USERNAME || pass != AUTH_PASSWORD {

		// We dont have valid credetials
		logger.Log("AUTH_ATTEMPT", "FAILED Invalid credentials")
		return "", false
	}

	// Valid credentials
	logger.Log("AUTH_ATTEMPT", "Valid credentials")

	return createSession()
}

// If returns true means that the request is authenticated
func ValidateAuthentication(key string) bool {
	mu.RLock()
	defer mu.RUnlock()

	// If entry with the key exist we are authenticated
	authTime, isAuthenticated := activeSessions[key]

	if !isAuthenticated {
		return false
	}

	if time.Since(authTime) > sessionDuration {
		revoqueSession(key)
		return false
	}

	return true
}

// Returns the session auth key and the status of the request
func createSession() (string, bool) {
	tokenBytes := make([]byte, 32)

	_, err := io.ReadFull(rand.Reader, tokenBytes)
	if err != nil {
		logger.Log("AUTH_ERROR", "Failed to generate secure token")
		return "", false
	}

	sessionKey := hex.EncodeToString(tokenBytes)

	//Store key on sessions map
	activeSessions[sessionKey] = time.Now()

	return sessionKey, true
}

func revoqueSession(key string) bool {
	logger.Log("Session expired need to reauthenticate")
	delete(activeSessions, key)
	return true
}
