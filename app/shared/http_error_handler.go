package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var errorLog *log.Logger

func init() {
	setupLogFile()
}

func setupLogFile() {
	today := time.Now().Format("2006-01-02")
	fileName := fmt.Sprintf("./var/log/errors-%s.log", today)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open error log file: %v", err)
	}
	errorLog = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func HandleHttpError(w http.ResponseWriter, r *http.Request, err error, httpStatusCode *int) {
	statusCode := http.StatusInternalServerError
	msg := "Internal server error"

	fmt.Println(w, "Error: %s\n", err.Error())

	if httpStatusCode != nil {
		statusCode = *httpStatusCode
		msg = err.Error()
	} else {
		var invalidInputErr *InvalidInputError
		if errors.As(err, &invalidInputErr) {
			statusCode = http.StatusBadRequest
			msg = invalidInputErr.Error()
		}
	}

	if statusCode == http.StatusNotFound {
		errorLog.Printf("404 Error: %s %s", r.Method, r.URL)
	} else if statusCode >= http.StatusInternalServerError {
		errorLog.Printf("WriteErrorToResponse: %v\nRequest: %s %s\nStack Trace:\n%s", err, r.Method, r.URL, debug.Stack())
	}

	if isJSONRequest(r) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(map[string]string{
			"error": msg,
		})
	} else {
		w.WriteHeader(statusCode)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	}
}

func isJSONRequest(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Accept"), "application/json") ||
		strings.Contains(r.Header.Get("Content-Type"), "application/json")
}
