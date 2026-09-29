package track

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func (s *Service) progressPath() string {
	return filepath.Join(s.Root, "progress", "progress.json")
}

func (s *Service) loadProgress() (map[string]Entry, error) {
	data, err := os.ReadFile(s.progressPath())
	if os.IsNotExist(err) {
		return map[string]Entry{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("não consegui ler o progresso: %w", err)
	}
	if len(bytesTrimSpace(data)) == 0 {
		return map[string]Entry{}, nil
	}
	prog := map[string]Entry{}
	if err := json.Unmarshal(data, &prog); err != nil {
		return nil, fmt.Errorf("progress/progress.json inválido: %w", err)
	}
	if prog == nil {
		prog = map[string]Entry{}
	}
	for key, entry := range prog {
		switch entry.Status {
		case StatusTodo, StatusAttempting, StatusSolved:
		default:
			return nil, fmt.Errorf("status inválido %q no problema %s", entry.Status, key)
		}
	}
	return prog, nil
}

func (s *Service) saveProgress(prog map[string]Entry) error {
	if prog == nil {
		prog = map[string]Entry{}
	}
	if err := os.MkdirAll(filepath.Dir(s.progressPath()), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(prog, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(s.progressPath(), data, 0o644)
}

func effectiveStatus(id int, prog map[string]Entry) string {
	entry, ok := prog[strconv.Itoa(id)]
	if !ok || entry.Status == "" || entry.Status == StatusTodo {
		return StatusTodo
	}
	return entry.Status
}

func bytesTrimSpace(data []byte) []byte {
	start, end := 0, len(data)
	for start < end && isSpace(data[start]) {
		start++
	}
	for end > start && isSpace(data[end-1]) {
		end--
	}
	return data[start:end]
}

func isSpace(b byte) bool {
	return b == ' ' || b == '\n' || b == '\r' || b == '\t'
}
