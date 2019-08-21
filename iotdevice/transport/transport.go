package transport

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/vbogretsov/iothub/logger"
	"github.com/vbogretsov/iothub/common"
)

// Transport interface.
type Transport interface {
	SetLogger(logger logger.Logger)
	Connect(ctx context.Context, creds Credentials) error
	Send(ctx context.Context, msg *common.Message) error
	RegisterDirectMethods(ctx context.Context, mux MethodDispatcher) error
	SubscribeEvents(ctx context.Context, mux MessageDispatcher) error
	SubscribeTwinUpdates(ctx context.Context, mux TwinStateDispatcher) error
	RetrieveTwinProperties(ctx context.Context) (payload []byte, err error)
	UpdateTwinProperties(ctx context.Context, payload []byte) (version int, err error)
	Close() error
}

type Credentials interface {
	GetDeviceID() string
	GetHostName() string
	GetCertificate() *tls.Certificate
	Token(resource string, lifetime time.Duration) (*common.SharedAccessSignature, error)
}

// MessageDispatcher handles incoming messages.
type MessageDispatcher interface {
	Dispatch(msg *common.Message)
}

// TwinStateDispatcher handles twin state updates.
type TwinStateDispatcher interface {
	Dispatch(b []byte)
}

// MethodDispatcher handles direct method calls.
type MethodDispatcher interface {
	Dispatch(methodName string, b []byte) (rc int, data []byte, err error)
}
