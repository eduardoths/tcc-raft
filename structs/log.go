package structs

type LogEntry struct {
	Term    int
	Index   int
	Command LogCommand
}

type LogCommand struct {
	Operation string
	Key       string
	Value     []byte
}
