package usecase

import (
	"sync"
)

// LogManager stores logs per project and allows subscribers to receive updates.
type LogManager struct {
	mu      sync.Mutex
	streams map[uint]*logStream
}

type logStream struct {
	mu          sync.Mutex
	logs        []string
	subscribers []chan string
	closed      bool
}

// NewLogManager creates a new LogManager.
func NewLogManager() *LogManager {
	return &LogManager{streams: make(map[uint]*logStream)}
}

func (m *LogManager) getStream(id uint) *logStream {
	m.mu.Lock()
	defer m.mu.Unlock()
	s, ok := m.streams[id]
	if !ok {
		s = &logStream{}
		m.streams[id] = s
	}
	return s
}

// Append adds a message to the project log and broadcasts to subscribers.
func (m *LogManager) Append(id uint, msg string) {
	s := m.getStream(id)
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return
	}
	s.logs = append(s.logs, msg)
	for _, sub := range s.subscribers {
		select {
		case sub <- msg:
		default:
		}
	}
}

// Subscribe returns a channel that receives future log messages for the project.
// Existing logs are sent immediately in a separate goroutine.
func (m *LogManager) Subscribe(id uint) <-chan string {
	s := m.getStream(id)
	ch := make(chan string, 10)
	s.mu.Lock()
	if s.closed {
		close(ch)
		s.mu.Unlock()
		return ch
	}
	s.subscribers = append(s.subscribers, ch)
	logs := append([]string(nil), s.logs...)
	s.mu.Unlock()
	go func() {
		for _, l := range logs {
			ch <- l
		}
	}()
	return ch
}

// Close marks the project log as closed and closes all subscriber channels.
func (m *LogManager) Close(id uint) {
	s := m.getStream(id)
	s.mu.Lock()
	if s.closed {
		s.mu.Unlock()
		return
	}
	s.closed = true
	subs := s.subscribers
	s.subscribers = nil
	s.mu.Unlock()
	for _, ch := range subs {
		close(ch)
	}
}

// GetLogs returns a copy of all logs for the project.
func (m *LogManager) GetLogs(id uint) []string {
	s := m.getStream(id)
	s.mu.Lock()
	defer s.mu.Unlock()
	logs := append([]string(nil), s.logs...)
	return logs
}
