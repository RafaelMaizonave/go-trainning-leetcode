package track

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Service lê o catálogo e grava o progresso no diretório do repositório.
type Service struct {
	Root    string
	Now     func() time.Time
	TestDir func(dir string) error
}

// New devolve um serviço que roda go test de verdade ao marcar um problema.
func New(root string) *Service {
	return &Service{
		Root:    root,
		Now:     time.Now,
		TestDir: runGoTest,
	}
}

// FindRoot sobe a partir de start até achar catalog/problems.json.
func FindRoot(start string) (string, error) {
	dir, err := filepath.Abs(start)
	if err != nil {
		return "", err
	}
	for {
		info, statErr := os.Stat(filepath.Join(dir, "catalog", "problems.json"))
		if statErr == nil && !info.IsDir() {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("não encontrei catalog/problems.json a partir de %s", start)
		}
		dir = parent
	}
}

func (s *Service) clock() time.Time {
	if s.Now != nil {
		return s.Now()
	}
	return time.Now()
}

func (s *Service) testDir(dir string) error {
	if s.TestDir != nil {
		return s.TestDir(dir)
	}
	return runGoTest(dir)
}

// Start cria a pasta da solução e marca o problema como tentando.
func (s *Service) Start(id int) error {
	p, err := s.mustProblem(id)
	if err != nil {
		return err
	}
	dir := s.solutionDir(p)
	if _, statErr := os.Stat(dir); statErr == nil {
		return fmt.Errorf("a pasta da solução já existe: solutions/%s", FolderName(p))
	} else if !os.IsNotExist(statErr) {
		return statErr
	}
	if err := os.MkdirAll(filepath.Join(s.Root, "solutions"), 0o755); err != nil {
		return err
	}
	if err := os.Mkdir(dir, 0o755); err != nil {
		return err
	}
	if err := writeScaffold(dir, p); err != nil {
		_ = os.RemoveAll(dir)
		return err
	}

	prog, err := s.loadProgress()
	if err != nil {
		return err
	}
	prog[strconv.Itoa(id)] = Entry{
		Status:    StatusAttempting,
		StartedAt: s.clock().Format("2006-01-02"),
	}
	if err := s.saveProgress(prog); err != nil {
		return err
	}
	return s.WriteReadme()
}

// Done roda os testes da solução e, se passarem, marca o problema como resolvido.
func (s *Service) Done(id int) error {
	p, err := s.mustProblem(id)
	if err != nil {
		return err
	}
	dir := s.solutionDir(p)
	info, statErr := os.Stat(dir)
	if statErr != nil || !info.IsDir() {
		return fmt.Errorf("crie a solução com start antes de marcar o problema %d como resolvido", id)
	}
	if _, err := os.Stat(filepath.Join(dir, "solution_test.go")); err != nil {
		return fmt.Errorf("a pasta solutions/%s não tem solution_test.go", FolderName(p))
	}
	if err := notesReady(filepath.Join(dir, "README.md")); err != nil {
		return err
	}
	if err := s.testDir(dir); err != nil {
		return err
	}

	prog, err := s.loadProgress()
	if err != nil {
		return err
	}
	key := strconv.Itoa(id)
	entry := prog[key]
	today := s.clock().Format("2006-01-02")
	if entry.StartedAt == "" {
		entry.StartedAt = today
	}
	entry.Status = StatusSolved
	entry.SolvedAt = today
	entry.ReviewedAt = ""
	prog[key] = entry
	if err := s.saveProgress(prog); err != nil {
		return err
	}
	if err := updateStatusLine(filepath.Join(dir, "README.md"), StatusSolved); err != nil {
		return err
	}
	return s.WriteReadme()
}

// Reset devolve o problema para a fazer sem apagar a solução.
func (s *Service) Reset(id int) error {
	p, err := s.mustProblem(id)
	if err != nil {
		return err
	}
	prog, err := s.loadProgress()
	if err != nil {
		return err
	}
	delete(prog, strconv.Itoa(id))
	if err := s.saveProgress(prog); err != nil {
		return err
	}
	if err := updateStatusLine(filepath.Join(s.solutionDir(p), "README.md"), StatusTodo); err != nil {
		return err
	}
	return s.WriteReadme()
}

// List devolve os problemas que passam no filtro, em ordem de id.
func (s *Service) List(f Filter) ([]Row, error) {
	difficulty, err := canonicalDifficulty(f.Difficulty)
	if err != nil {
		return nil, err
	}
	status, err := canonicalStatus(f.Status)
	if err != nil {
		return nil, err
	}
	tag := strings.ToLower(strings.TrimSpace(f.Tag))

	problems, err := s.loadCatalog()
	if err != nil {
		return nil, err
	}
	prog, err := s.loadProgress()
	if err != nil {
		return nil, err
	}

	rows := make([]Row, 0, len(problems))
	for _, p := range problems {
		current := effectiveStatus(p.ID, prog)
		if difficulty != "" && p.Difficulty != difficulty {
			continue
		}
		if status != "" && current != status {
			continue
		}
		if tag != "" && !hasTag(p, tag) {
			continue
		}
		rows = append(rows, Row{
			Problem:     p,
			Status:      current,
			SolutionRel: s.solutionRel(p),
		})
	}
	return rows, nil
}

// Stats conta resolvidos, tentando e a fazer, inclusive por dificuldade.
func (s *Service) Stats() (Stats, error) {
	problems, err := s.loadCatalog()
	if err != nil {
		return Stats{}, err
	}
	prog, err := s.loadProgress()
	if err != nil {
		return Stats{}, err
	}

	var stats Stats
	stats.Total = len(problems)
	solvedBy := map[string]int{}
	totalBy := map[string]int{}
	for _, p := range problems {
		totalBy[p.Difficulty]++
		switch effectiveStatus(p.ID, prog) {
		case StatusSolved:
			stats.Solved++
			solvedBy[p.Difficulty]++
		case StatusAttempting:
			stats.Attempting++
		default:
			stats.Todo++
		}
	}
	stats.ByDifficulty = make([]DifficultyStat, 0, len(difficultyOrder))
	for _, difficulty := range difficultyOrder {
		stats.ByDifficulty = append(stats.ByDifficulty, DifficultyStat{
			Difficulty: difficulty,
			Solved:     solvedBy[difficulty],
			Total:      totalBy[difficulty],
		})
	}
	return stats, nil
}

func (s *Service) mustProblem(id int) (Problem, error) {
	if id <= 0 {
		return Problem{}, fmt.Errorf("id inválido %d", id)
	}
	problems, err := s.loadCatalog()
	if err != nil {
		return Problem{}, err
	}
	p, ok := findProblem(problems, id)
	if !ok {
		return Problem{}, fmt.Errorf("problema %d não está no catálogo", id)
	}
	return p, nil
}

func (s *Service) solutionRel(p Problem) string {
	dir := s.solutionDir(p)
	info, err := os.Stat(dir)
	if err != nil || !info.IsDir() {
		return ""
	}
	solution := filepath.Join(dir, "solution.go")
	if fileInfo, fileErr := os.Stat(solution); fileErr == nil && !fileInfo.IsDir() {
		return "solutions/" + FolderName(p) + "/solution.go"
	}
	return "solutions/" + FolderName(p)
}

func hasTag(p Problem, needle string) bool {
	for _, tag := range p.Tags {
		if strings.Contains(strings.ToLower(tag), needle) {
			return true
		}
	}
	return false
}

func runGoTest(dir string) error {
	cmd := exec.Command("go", "test", ".")
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		msg := strings.TrimSpace(string(out))
		if msg == "" {
			msg = err.Error()
		}
		return fmt.Errorf("os testes falharam:\n%s", msg)
	}
	return nil
}

func osWriteFile(root, name string, data []byte) error {
	return os.WriteFile(filepath.Join(root, name), data, 0o644)
}
