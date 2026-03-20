// T.U.I — To Unify Imagination
// internal/server/server.go — Local HTTP event bus
// Agents post events here; the TUI reads them via SSE or WebSocket.
package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/Phoenixai36/To-You-I/internal/model"
)

// Server is the local HTTP event bus that bridges agents and the TUI.
type Server struct {
	host string
	port int

	mu       sync.RWMutex
	listeners []chan model.AgentEvent
}

// New creates a new Server bound to host:port.
func New(host string, port int) *Server {
	return &Server{host: host, port: port}
}

// Start registers HTTP routes and starts listening.
func (s *Server) Start() error {
	mux := http.NewServeMux()

	// POST /events — agents push events here
	mux.HandleFunc("/events", s.handlePostEvent)

	// GET /events/stream — TUI subscribes to SSE stream
	mux.HandleFunc("/events/stream", s.handleSSEStream)

	// GET /health — simple healthcheck
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok","project":"T.U.I"}`))
	})

	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	fmt.Printf("T.U.I event server listening on http://%s\n", addr)
	return http.ListenAndServe(addr, mux)
}

// broadcast sends an event to all connected SSE listeners.
func (s *Server) broadcast(ev model.AgentEvent) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, ch := range s.listeners {
		select {
		case ch <- ev:
		default:
			// drop if listener is too slow
		}
	}
}

// handlePostEvent receives an AgentEvent JSON payload from an agent adapter.
func (s *Server) handlePostEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var ev model.AgentEvent
	if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
		http.Error(w, "bad request: "+err.Error(), http.StatusBadRequest)
		return
	}
	s.broadcast(ev)
	w.WriteHeader(http.StatusAccepted)
}

// handleSSEStream streams events to a TUI subscriber via Server-Sent Events.
func (s *Server) handleSSEStream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ch := make(chan model.AgentEvent, 64)
	s.mu.Lock()
	s.listeners = append(s.listeners, ch)
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		for i, l := range s.listeners {
			if l == ch {
				s.listeners = append(s.listeners[:i], s.listeners[i+1:]...)
				break
			}
		}
		s.mu.Unlock()
		close(ch)
	}()

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming not supported", http.StatusInternalServerError)
		return
	}

	for {
		select {
		case ev := <-ch:
			data, _ := json.Marshal(ev)
			fmt.Fprintf(w, "data: %s\n\n", data)
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}
