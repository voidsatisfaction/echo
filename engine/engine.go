package engine

import (
	"io"
	"time"

	"github.com/labstack/echo/logger"
)

type (
	HandlerFunc func(Request, Response)

	Engine interface {
		SetHandler(HandlerFunc)
		SetLogger(logger.Logger)
		Start()
	}

	Request interface {
		TLS() bool
		Scheme() string
		Host() string
		URI() string
		URL() URL
		Header() Header
		// Proto() string
		// ProtoMajor() int
		// ProtoMinor() int
		RemoteAddress() string
		Method() string
		Body() io.ReadCloser
		FormValue(string) string
		Object() interface{}
	}

	Response interface {
		Header() Header
		WriteHeader(int)
		Write(b []byte) (int, error)
		Status() int
		Size() int64
		Committed() bool
		SetWriter(io.Writer)
		Writer() io.Writer
		Object() interface{}
	}

	Header interface {
		Add(string, string)
		Del(string)
		Get(string) string
		Set(string, string)
		Object() interface{}
	}

	URL interface {
		SetPath(string)
		Path() string
		QueryValue(string) string
		Object() interface{}
	}

	Config struct {
		Address      string
		TLSCertfile  string
		TLSKeyfile   string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}
)
