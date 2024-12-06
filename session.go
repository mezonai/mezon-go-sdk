package mezonsdk

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
)

// ISession represents the session interface.
type ISession interface {
	IsExpired(currentTime int64) bool
	IsRefreshExpired(currentTime int64) bool
	Update(token string, refreshToken string) error
}

// Session represents a user session.
type Session struct {
	Token            string                 `json:"token"`
	Created          bool                   `json:"created"`
	CreatedAt        int64                  `json:"created_at"`
	ExpiresAt        *int64                 `json:"expires_at,omitempty"`
	RefreshExpiresAt *int64                 `json:"refresh_expires_at,omitempty"`
	RefreshToken     string                 `json:"refresh_token"`
	Username         *string                `json:"username,omitempty"`
	UserID           *string                `json:"user_id,omitempty"`
	Vars             map[string]interface{} `json:"vars,omitempty"`
}

// NewSession creates a new session instance.
func NewSession(token string, refreshToken string, created bool) (*Session, error) {
	session := &Session{
		Token:        token,
		RefreshToken: refreshToken,
		Created:      created,
		CreatedAt:    time.Now().Unix(),
	}

	if err := session.Update(token, refreshToken); err != nil {
		return nil, err
	}

	return session, nil
}

// IsExpired checks if the session token has expired.
func (s *Session) IsExpired(currentTime int64) bool {
	if s.ExpiresAt == nil {
		return true
	}
	return (*s.ExpiresAt - currentTime) < 0
}

// IsRefreshExpired checks if the refresh token has expired.
func (s *Session) IsRefreshExpired(currentTime int64) bool {
	if s.RefreshExpiresAt == nil {
		return true
	}
	return (*s.RefreshExpiresAt - currentTime) < 0
}

// Update updates the session with new tokens.
func (s *Session) Update(token string, refreshToken string) error {
	tokenParts := strings.Split(token, ".")
	if len(tokenParts) != 3 {
		return errors.New("JWT is not valid")
	}

	tokenDecoded, err := decodeJWTPart(tokenParts[1])
	if err != nil {
		return err
	}

	expStr, ok := tokenDecoded["exp"].(string)
	if !ok {
		return errors.New("invalid token expiry format")
	}

	exp, err := strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		return err
	}

	s.ExpiresAt = &exp

	if username, ok := tokenDecoded["usn"].(string); ok {
		s.Username = &username
	}

	if userID, ok := tokenDecoded["uid"].(string); ok {
		s.UserID = &userID
	}

	if vars, ok := tokenDecoded["vrs"].(map[string]interface{}); ok {
		s.Vars = vars
	}

	if refreshToken != "" {
		refreshTokenParts := strings.Split(refreshToken, ".")
		if len(refreshTokenParts) != 3 {
			return errors.New("refresh JWT is not valid")
		}

		refreshTokenDecoded, err := decodeJWTPart(refreshTokenParts[1])
		if err != nil {
			return err
		}

		refreshExpStr, ok := refreshTokenDecoded["exp"].(string)
		if !ok {
			return errors.New("invalid refresh token expiry format")
		}

		refreshExp, err := strconv.ParseInt(refreshExpStr, 10, 64)
		if err != nil {
			return err
		}

		s.RefreshExpiresAt = &refreshExp
		s.RefreshToken = refreshToken
	}

	s.Token = token
	return nil
}

// Restore restores a session from tokens.
func Restore(token string, refreshToken string) (*Session, error) {
	return NewSession(token, refreshToken, false)
}

// Helper function to decode a JWT part.
func decodeJWTPart(part string) (map[string]interface{}, error) {
	decoded, err := base64.RawURLEncoding.DecodeString(part)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(decoded, &result); err != nil {
		return nil, err
	}

	return result, nil
}
