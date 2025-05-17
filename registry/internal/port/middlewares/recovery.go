package middlewares

import (
	"log/slog"
	"net/http"
	"runtime/debug"
	"strings"
)

type Recovery struct {
	nextHandler http.Handler
}

func NewRecovery(nextHandler http.Handler) *Recovery {
	return &Recovery{
		nextHandler: nextHandler,
	}
}

func (r *Recovery) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		r := recover()
		if r != nil {
			var err string
			switch t := r.(type) {
			case string:
				err = t
			case error:
				err = t.Error()
			default:
				err = "unknown error"
			}

			slog.
				With(slog.String("err", err)).
				With(slog.Any("stack", getStackTrace())).
				ErrorContext(req.Context(), "recovered from panic")

			http.Error(w, "internal error", http.StatusInternalServerError)
		}
	}()

	r.nextHandler.ServeHTTP(w, req)
}

func getStackTrace() []string {
	stack := strings.ReplaceAll(string(debug.Stack()), "\t", "")
	stackRows := strings.Split(stack, "\n")
	return stackRows
}
