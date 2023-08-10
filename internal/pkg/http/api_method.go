package http

import (
	"bytes"
	"context"
	"fmt"
	"github.com/avast/retry-go"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// APIError ...
type APIError struct {
	StatusCode int
	Error      string
}

var (
	responseHeadersError = fmt.Errorf("net/http: timeout awaiting response headers")
	timeout503Error      = fmt.Errorf("api call has timed out")
	unmarshalError       = fmt.Errorf("Could not unmarshal the request")
)

// Do returns response body in bytes with status code or error with status code
func (dao *Service) Do(ctx context.Context, request *http.Request) ([]byte, int, error) {
	log.Infof("making api call to %s", request.URL.String())
	request = request.WithContext(ctx)

	var err error
	if request.Body != nil {
		requestBodyBytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.WithError(err).Error(unmarshalError)
			return nil, http.StatusInternalServerError, unmarshalError
		}

		// body can only be read once, so create a new reader
		request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBodyBytes))
	}
	var response *http.Response
	retryError := retry.Do(
		func() error {
			start := time.Now()
			response, err = dao.HTTPClient.Do(request)
			log.Infof("received response from api call to %s with duration %d ms", request.URL.String(), time.Now().Sub(start).Milliseconds())
			if err != nil {
				if strings.Contains(err.Error(), responseHeadersError.Error()) {
					log.Warnf("received a context timeout from api %s", request.URL.String())
					return responseHeadersError
				}
				log.Error(err)
				return err
			}
			if response == nil {
				log.Errorf("received a nil response from %s", request.URL.String())
				err = fmt.Errorf("an error occurred trying to call http do request")
				return err
			}
			if response.StatusCode == 503 {
				log.Warnf("received a 503 response code from api %s", request.URL.String())
				return timeout503Error
			}
			return nil
		},
		retry.RetryIf(func(err error) bool {
			if err == timeout503Error || err == responseHeadersError {
				return true
			}
			return false
		}),
	)
	if retryError != nil {
		return nil, http.StatusServiceUnavailable, fmt.Errorf("retry attempts failed for call %s", request.URL.String())
	}
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.WithError(err).Error(unmarshalError)
		return nil, http.StatusInternalServerError, unmarshalError
	}

	return bodyBytes, response.StatusCode, nil
}

// WithTimeout adds timeout to context for http requests
func (dao *Service) WithTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, time.Duration(dao.Config.ClientTimeoutInSeconds)*time.Second)
}
