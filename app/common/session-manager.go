package common

import (
	"net/http"

	"github.com/gorilla/sessions"
)

const userIdKey = "userId"

type SessionManager struct {
	store      *sessions.FilesystemStore
	sessionKey string
}

func NewSessionManager(path, secretKey, sessionKey string) *SessionManager {
	store := sessions.NewFilesystemStore(path, []byte(secretKey))
	store.MaxLength(0)
	return &SessionManager{store: store, sessionKey: sessionKey}
}

func (sm *SessionManager) SetUser(w http.ResponseWriter, r *http.Request, value string) error {
	err := sm.setValue(w, r, userIdKey, value)
	if err != nil {
		return err
	}

	return nil
}

func (sm *SessionManager) ClearUser(w http.ResponseWriter, r *http.Request) error {
	err := sm.deleteValue(w, r, userIdKey)
	if err != nil {
		return err
	}

	return nil
}

func (sm *SessionManager) GetUser(r *http.Request) (string, error) {
	value, err := sm.getValue(r, userIdKey)
	if err != nil {
		return "", err
	}

	userId, ok := value.(string)
	if !ok {
		return "", nil // Якщо тип не відповідає, повертаємо порожній рядок
	}

	return userId, nil
}

func (sm *SessionManager) setValue(w http.ResponseWriter, r *http.Request, key string, value interface{}) error {
	session, err := sm.store.Get(r, sm.sessionKey)
	if err != nil {
		return err
	}

	session.Values[key] = value

	return session.Save(r, w)
}

func (sm *SessionManager) getValue(r *http.Request, key string) (interface{}, error) {
	session, err := sm.store.Get(r, sm.sessionKey)
	if err != nil {
		return nil, err
	}

	return session.Values[key], nil
}

func (sm *SessionManager) deleteValue(w http.ResponseWriter, r *http.Request, key string) error {
	session, err := sm.store.Get(r, sm.sessionKey)
	if err != nil {
		return err
	}

	delete(session.Values, key)
	return session.Save(r, w)
}