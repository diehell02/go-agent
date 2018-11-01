package internal

import (
	"net/http"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func BonreeRoundTripper(crossReqheader string) http.RoundTripper {
	return roundTripperFunc(func(request *http.Request) (*http.Response, error) {
		request.Header.Set(CrossRequestHeader, crossReqheader)

		response, err := http.DefaultTransport.RoundTrip(request)

		return response, err
	})
}