package eventconsumer

type EventConsumer interface {
	StartConsuming() error
}
