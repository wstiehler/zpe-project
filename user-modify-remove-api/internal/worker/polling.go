package worker

func createPolling(input Input, consumer Consumer) {
	go ListenToNats(input, consumer)

}
