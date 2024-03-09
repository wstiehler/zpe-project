package worker

type Consumer interface {
	QueueSubject() string
	PollingIntervalSeconds() int64
	Handler(Input, UserEntity) error
}
