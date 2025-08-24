package hello_service

import context "context"

// PerRPCCredentials defines the common interface for the credentials which need to
// attach security information to every RPC (e.g., oauth2).
// type PerRPCCredentials interface {
// 	// GetRequestMetadata gets the current request metadata, refreshing tokens
// 	// if required. This should be called by the transport layer on each
// 	// request, and the data should be populated in headers or other
// 	// context. If a status code is returned, it will be used as the status for
// 	// the RPC (restricted to an allowable set of codes as defined by gRFC
// 	// A54). uri is the URI of the entry point for the request.  When supported
// 	// by the underlying implementation, ctx can be used for timeout and
// 	// cancellation. Additionally, RequestInfo data will be available via ctx
// 	// to this call.  TODO(zhaoq): Define the set of the qualified keys instead
// 	// of leaving it as an arbitrary string.
// 	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
// 	// RequireTransportSecurity indicates whether the credentials requires
// 	// transport security.
// 	RequireTransportSecurity() bool
// }

func NewClientAuthentication(clientId, clientSecret string) *Authentication {
	return &Authentication{
		clientID:     clientId,
		clientSecret: clientSecret,
	}
}

// Authentication todo
type Authentication struct {
	clientID     string
	clientSecret string
}

// WithClientCredentials todo
func (a *Authentication) WithClientCredentials(clientID, clientSecret string) {
	a.clientID = clientID
	a.clientSecret = clientSecret
}

// GetRequestMetadata todo
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (
	map[string]string, error,
) {
	return map[string]string{
		ClientHeaderKey: a.clientID,
		ClientSecretKey: a.clientSecret,
	}, nil
}

// RequireTransportSecurity todo
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}
