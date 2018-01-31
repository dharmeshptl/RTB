package worker

type MessageProcessor interface {
	Process(message []byte) error
}
