package httpechoes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Config gathers available configuration values.
type Config struct {
	Port              int
	MaxBytes          int64
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	ReadHeaderTimeout time.Duration
}

func (config Config) Addr() string {
	return fmt.Sprintf(":%d", config.Port)
}

// NewDefaultConfig creates a default configuration.
func NewDefaultConfig() Config {
	return Config{
		Port:              8080,
		MaxBytes:          1024,
		ReadTimeout:       time.Second * 5,
		WriteTimeout:      time.Second * 5,
		IdleTimeout:       time.Second * 30,
		ReadHeaderTimeout: time.Second * 2,
	}
}

// New returns an HTTP server that echoes status codes.
func New(config Config) *http.Server {
	statuses := []int{
		http.StatusContinue,
		http.StatusSwitchingProtocols,
		http.StatusProcessing,
		http.StatusEarlyHints,
		http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
		http.StatusNonAuthoritativeInfo,
		http.StatusNoContent,
		http.StatusResetContent,
		http.StatusPartialContent,
		http.StatusMultiStatus,
		http.StatusAlreadyReported,
		http.StatusIMUsed,
		http.StatusMultipleChoices,
		http.StatusMovedPermanently,
		http.StatusFound,
		http.StatusSeeOther,
		http.StatusNotModified,
		http.StatusUseProxy,
		http.StatusTemporaryRedirect,
		http.StatusPermanentRedirect,
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusPaymentRequired,
		http.StatusForbidden,
		http.StatusNotFound,
		http.StatusMethodNotAllowed,
		http.StatusNotAcceptable,
		http.StatusProxyAuthRequired,
		http.StatusRequestTimeout,
		http.StatusConflict,
		http.StatusGone,
		http.StatusLengthRequired,
		http.StatusPreconditionFailed,
		http.StatusRequestEntityTooLarge,
		http.StatusRequestURITooLong,
		http.StatusUnsupportedMediaType,
		http.StatusRequestedRangeNotSatisfiable,
		http.StatusExpectationFailed,
		http.StatusTeapot,
		http.StatusMisdirectedRequest,
		http.StatusUnprocessableEntity,
		http.StatusLocked,
		http.StatusFailedDependency,
		http.StatusTooEarly,
		http.StatusUpgradeRequired,
		http.StatusPreconditionRequired,
		http.StatusTooManyRequests,
		http.StatusRequestHeaderFieldsTooLarge,
		http.StatusUnavailableForLegalReasons,
		http.StatusInternalServerError,
		http.StatusNotImplemented,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
		http.StatusHTTPVersionNotSupported,
		http.StatusVariantAlsoNegotiates,
		http.StatusInsufficientStorage,
		http.StatusLoopDetected,
		http.StatusNotExtended,
		http.StatusNetworkAuthenticationRequired,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		for _, code := range statuses {
			_, err := fmt.Fprintf(w, "/%d → %s\n", code, http.StatusText(code))
			if err != nil {
				fmt.Printf("[ERROR] %s\n", err)
			}
		}
	})

	for _, code := range statuses {
		code := code
		mux.HandleFunc("/"+strconv.Itoa(code), func(w http.ResponseWriter, r *http.Request) {
			r.Body = http.MaxBytesReader(w, r.Body, config.MaxBytes)
			w.WriteHeader(code)
			_, err := fmt.Fprintf(w, "%d %s\n", code, http.StatusText(code))
			if err != nil {
				fmt.Printf("[ERROR] %s\n", err)
			}
		})
	}

	return &http.Server{
		Addr:              config.Addr(),
		Handler:           mux,
		ReadTimeout:       config.ReadTimeout,
		WriteTimeout:      config.WriteTimeout,
		IdleTimeout:       config.IdleTimeout,
		ReadHeaderTimeout: config.ReadHeaderTimeout,
	}
}
