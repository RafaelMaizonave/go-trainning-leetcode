package track

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	StatusTodo       = "todo"
	StatusAttempting = "attempting"
	StatusSolved     = "solved"

	DifficultyEasy   = "Easy"
	DifficultyMedium = "Medium"
	DifficultyHard   = "Hard"
)

// Problem é uma entrada do catálogo, sem a solução.
type Problem struct {
	ID         int      `json:"id"`
	Title      string   `json:"title"`
	Slug       string   `json:"slug"`
	Difficulty string   `json:"difficulty"`
	Tags       []string `json:"tags"`
	URL        string   `json:"url"`
}

// Entry é o progresso de um problema já iniciado.
type Entry struct {
	Status     string `json:"status"`
	StartedAt  string `json:"started_at,omitempty"`
	SolvedAt   string `json:"solved_at,omitempty"`
	ReviewedAt string `json:"reviewed_at,omitempty"`
}

// ReviewItem é um problema resolvido que já passou de um marco de revisão.
type ReviewItem struct {
	Problem     Problem
	SolvedAt    string
	DueSince    string
	SolutionRel string
}

// Filter restringe a listagem. Campos vazios não filtram.
type Filter struct {
	Difficulty string
	Status     string
	Tag        string
}

// Row é um problema com o status efetivo e o caminho da solução, se existir.
type Row struct {
	Problem     Problem
	Status      string
	SolutionRel string
}

// DifficultyStat conta resolvidos e o total de uma dificuldade.
type DifficultyStat struct {
	Difficulty string
	Solved     int
	Total      int
}

// Stats resume o progresso no formato do painel.
type Stats struct {
	Solved       int
	Attempting   int
	Todo         int
	Total        int
	ByDifficulty []DifficultyStat
}

var difficultyOrder = []string{DifficultyEasy, DifficultyMedium, DifficultyHard}

var slugPattern = regexp.MustCompile(`^[a-z0-9_]+$`)

// StatusLabel devolve o rótulo em português exibido no README.
func StatusLabel(status string) string {
	switch status {
	case StatusAttempting:
		return "Tentando"
	case StatusSolved:
		return "Resolvido"
	default:
		return "A fazer"
	}
}

// FolderName é o diretório da solução, com id de 4 dígitos.
func FolderName(p Problem) string {
	return fmt.Sprintf("%04d_%s", p.ID, p.Slug)
}

// PackageName deriva um nome de pacote Go válido a partir do slug.
func PackageName(slug string) string {
	var b strings.Builder
	for _, r := range slug {
		if r == '_' || r == '-' {
			continue
		}
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
		}
	}
	name := b.String()
	if name == "" || (name[0] >= '0' && name[0] <= '9') {
		name = "p" + name
	}
	return name
}

func canonicalDifficulty(s string) (string, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "":
		return "", nil
	case "easy":
		return DifficultyEasy, nil
	case "medium":
		return DifficultyMedium, nil
	case "hard":
		return DifficultyHard, nil
	default:
		return "", fmt.Errorf("dificuldade inválida %q (use easy, medium ou hard)", s)
	}
}

func canonicalStatus(s string) (string, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "":
		return "", nil
	case StatusTodo, StatusAttempting, StatusSolved:
		return strings.ToLower(strings.TrimSpace(s)), nil
	default:
		return "", fmt.Errorf("status inválido %q (use todo, attempting ou solved)", s)
	}
}
