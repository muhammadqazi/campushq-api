package logs

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
)

var (
	debugColor   = color.New(color.FgMagenta).SprintFunc()
	infoColor    = color.New(color.FgCyan).SprintFunc()
	warningColor = color.New(color.FgYellow).SprintFunc()
	errorColor   = color.New(color.FgRed).SprintFunc()
	whiteColor   = color.New(color.FgWhite).SprintFunc()
	postColor    = color.New(color.FgGreen).SprintFunc()
	getColor     = color.New(color.FgBlue).SprintFunc()
	putColor     = color.New(color.FgYellow).SprintFunc()
	deleteColor  = color.New(color.FgRed).SprintFunc()
)

var (
	DebugLevel   = "debug"
	InfoLevel    = "info"
	WarningLevel = "warning"
	ErrorLevel   = "error"
)

type Logger struct {
	httpRequest *http.Request
}

func NewLogger(req *http.Request) *Logger {
	return &Logger{
		httpRequest: req,
	}
}

func (l *Logger) Debug(message string, path string) {
	log.Println(debugColor("[DEBUG] " + message + " " + color.BlueString(path)))
}

func (l *Logger) Info(message string, path string) {
	log.Println(infoColor("[INFO] " + message + " " + color.BlueString(path)))
}

func (l *Logger) Warning(message string, path string) {
	log.Println(warningColor("[WARNING] " + message + " " + color.BlueString(path)))
}

func (l *Logger) Error(message string, path string) {
	log.Println(errorColor("[ERROR] " + message + " " + color.BlueString(path)))
}

func (l *Logger) PrintHTTPRequest() {
	if l.httpRequest != nil {
		userAgent := l.httpRequest.Header.Get("User-Agent")

		switch l.httpRequest.Method {
		case http.MethodPost:
			log.Println(postColor("[POST - Request] " + color.BlueString(l.httpRequest.URL.Path) + " " + userAgent))
		case http.MethodGet:
			log.Println(getColor("[GET - Request] " + color.BlueString(l.httpRequest.URL.Path) + " " + userAgent))
		case http.MethodPatch:
			log.Println(putColor("[PATCH - Request] " + color.BlueString(l.httpRequest.URL.Path) + " " + userAgent))
		case http.MethodDelete:
			log.Println(deleteColor("[DELETE - Request] " + color.BlueString(l.httpRequest.URL.Path) + " " + userAgent))
		default:
			log.Println(whiteColor("[UNKNOWN - Request] " + color.BlueString(l.httpRequest.URL.Path) + " " + userAgent))
		}
	}
}

func (l *Logger) PrintHTTPResponse(r *http.Request, status int, message string) {
	if r != nil {
		switch r.Method {
		case http.MethodPost:
			log.Println(postColor("[POST - Response] " + color.BlueString(r.URL.Path) + " " + debugColor(fmt.Sprint(status)) + " " + message))
		case http.MethodGet:
			log.Println(getColor("[GET - Response] " + color.BlueString(r.URL.Path) + " " + debugColor(fmt.Sprint(status)) + " " + message))
		case http.MethodPatch:
			log.Println(putColor("[PATCH - Response] " + color.BlueString(r.URL.Path) + " " + debugColor(fmt.Sprint(status)) + " " + message))
		case http.MethodDelete:
			log.Println(deleteColor("[DELETE - Response] " + color.BlueString(r.URL.Path) + " " + debugColor(fmt.Sprint(status)) + " " + message))
		default:
			log.Println(whiteColor("[UNKNOWN - Response] " + color.BlueString(r.URL.Path) + " " + debugColor(fmt.Sprint(status)) + " " + message))
		}
	}
}
