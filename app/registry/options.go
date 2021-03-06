package registry

import (
	"context"
	"crypto/tls"
	"time"
)

type Options struct {
	Addrs     []string
	Timeout   time.Duration
	Secure    bool
	TLSConfig *tls.Config
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

type AddOptions struct {
	TTL time.Duration
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
	// Domain to register the service in
	Domain string
}

type WatchOptions struct {
	// Specify a service to watch
	// If blank, the watch is for all services
	App string
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
	// Domain to watch
	Domain string
}

type RemoveOptions struct {
	Context context.Context
	// Domain the service was registered in
	Domain string
}

type GetOptions struct {
	Context context.Context
	// Domain to scope the request to
	Domain string
}

type ListOptions struct {
	Context context.Context
	// Domain to scope the request to
	Domain string
}

// Addrs is the registry addresses to use
func Addrs(addrs ...string) Option {
	return func(o *Options) {
		o.Addrs = addrs
	}
}

func Timeout(t time.Duration) Option {
	return func(o *Options) {
		o.Timeout = t
	}
}

// Secure communication with the registry
func Secure(b bool) Option {
	return func(o *Options) {
		o.Secure = b
	}
}

// Specify TLS Config
func TLSConfig(t *tls.Config) Option {
	return func(o *Options) {
		o.TLSConfig = t
	}
}

func AddTTL(t time.Duration) AddOption {
	return func(o *AddOptions) {
		o.TTL = t
	}
}

func AddContext(ctx context.Context) AddOption {
	return func(o *AddOptions) {
		o.Context = ctx
	}
}

func AddDomain(d string) AddOption {
	return func(o *AddOptions) {
		o.Domain = d
	}
}

// Watch a service
func WatchApp(name string) WatchOption {
	return func(o *WatchOptions) {
		o.App = name
	}
}

func WatchContext(ctx context.Context) WatchOption {
	return func(o *WatchOptions) {
		o.Context = ctx
	}
}

func WatchDomain(d string) WatchOption {
	return func(o *WatchOptions) {
		o.Domain = d
	}
}

func RemoveContext(ctx context.Context) RemoveOption {
	return func(o *RemoveOptions) {
		o.Context = ctx
	}
}

func RemoveDomain(d string) RemoveOption {
	return func(o *RemoveOptions) {
		o.Domain = d
	}
}

func GetContext(ctx context.Context) GetOption {
	return func(o *GetOptions) {
		o.Context = ctx
	}
}

func GetDomain(d string) GetOption {
	return func(o *GetOptions) {
		o.Domain = d
	}
}

func ListContext(ctx context.Context) ListOption {
	return func(o *ListOptions) {
		o.Context = ctx
	}
}

func ListDomain(d string) ListOption {
	return func(o *ListOptions) {
		o.Domain = d
	}
}
