package session

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

const userIdKey = "userId"
const msgKey = "msg"

type Manager struct {
	store      *sessions.FilesystemStore
	sessionKey string
}

func NewSessionManager(path, secretKey, sessionKey string) *Manager {
	store := sessions.NewFilesystemStore(path, []byte(secretKey))
	store.MaxLength(0)
	store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   3600 * 12,
	}
	return &Manager{store: store, sessionKey: sessionKey}
}

func (sm *Manager) SetUser(w http.ResponseWriter, r *http.Request, value string) error {
	err := sm.setValue(w, r, userIdKey, value)
	if err != nil {
		return err
	}

	return nil
}

func (sm *Manager) ClearUser(w http.ResponseWriter, r *http.Request) error {
	err := sm.deleteValue(w, r, userIdKey)
	if err != nil {
		return err
	}

	return nil
}

func (sm *Manager) GetUser(r *http.Request) (string, error) {
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

func (sm *Manager) setValue(w http.ResponseWriter, r *http.Request, key string, value interface{}) error {
	session, err := sm.store.Get(r, sm.sessionKey)
	if err != nil {
		return err
	}

	session.Values[key] = value

	return session.Save(r, w)
}

func (sm *Manager) getValue(r *http.Request, key string) (interface{}, error) {
	session, err := sm.store.Get(r, sm.sessionKey)
	if err != nil {
		return nil, err
	}

	return session.Values[key], nil
}

func (sm *Manager) deleteValue(w http.ResponseWriter, r *http.Request, key string) error {
	session, err := sm.store.Get(r, sm.sessionKey)
	if err != nil {
		return err
	}

	delete(session.Values, key)
	return session.Save(r, w)
}

func (sm *Manager) AddInfoMessage(w http.ResponseWriter, r *http.Request, msg string) error {
	sm.addMessage(w, r, "success", msg)

	return nil
}

func (sm *Manager) AddWarningMessage(w http.ResponseWriter, r *http.Request, msg string) error {
	sm.addMessage(w, r, "success", msg)

	return nil
}

func (sm *Manager) retrieveMessages(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return sm.getValue(r, msgKey)
}

func (sm *Manager) addMessage(w http.ResponseWriter, r *http.Request, msgType string, msg string) error {
	value, err := sm.getValue(r, msgKey)
	if err != nil {
		log.Printf("Error retrieving value: %v", err)
		return nil
	}

	prevMsgs, ok := value.([]map[string]string)
	if !ok {
		log.Println("Type assertion to []map[string]string failed")
		return nil
	}

	newMessage := map[string]string{
		"type": msgType,
		"msg":  msg,
	}

	messages := append(prevMsgs, newMessage)
	sm.setValue(w, r, msgKey, messages)

	return nil
}
