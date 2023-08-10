package http

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/avast/retry-go"
)

const (
	//AuthorizationTokenHeader ...
	AuthorizationTokenHeader = "Authorization"
	//AcceptHeader http header for content type
	AcceptHeader = "Accept"
	//ContentTypeHeader ...
	ContentTypeHeader = "Content-Type"
	//JSONContentType ...
	JSONContentType = "application/json"
	//BearerToken ...
	BearerToken = "Bearer"
)

// Service ...
type Service struct {
	HTTPClient *http.Client
	Config     *Config
}

// Config holds the db connection information
type Config struct {
	ClientTimeoutInSeconds                 int `required:"true" split_words:"true" default:"30"`
	DefaultRetryAttemptsOn503              int `required:"true" split_words:"true" default:"3"`
	DefaultRetryBackOffOn503InMilliseconds int `required:"true" split_words:"true" default:"200"`
}

// NewHTTPDAO ....
func NewHTTPDAO(config *Config) *Service {
	s := &Service{
		Config: config,
	}
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			MinVersion:         tls.VersionTLS11,
		},
	}
	if config.DefaultRetryAttemptsOn503 != 0 {
		retry.DefaultAttempts = uint(config.DefaultRetryAttemptsOn503)
	}
	if config.DefaultRetryBackOffOn503InMilliseconds != 0 {
		retry.DefaultDelay = time.Duration(config.DefaultRetryBackOffOn503InMilliseconds) * time.Millisecond
	}
	c := &http.Client{
		Timeout:   time.Second * time.Duration(config.ClientTimeoutInSeconds),
		Transport: transport,
	}
	s.HTTPClient = c
	return s
}
