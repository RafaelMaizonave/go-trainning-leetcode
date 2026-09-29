package track

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	statusLine     = regexp.MustCompile(`(?m)^- Status: .*$`)
	approachLine   = regexp.MustCompile(`(?m)^- Abordagem:[ \t]*(.*)$`)
	complexityLine = regexp.MustCompile(`(?m)^- Complexidade:[ \t]*(.*)$`)
)

func (s *Service) solutionDir(p Problem) string {
	return filepath.Join(s.Root, "solutions", FolderName(p))
}

func writeScaffold(dir string, p Problem) error {
	pkg := PackageName(p.Slug)
	solution := fmt.Sprintf(`package %s

// Solve é o ponto de partida da sua solução.
// Troque a assinatura pela função pedida no LeetCode.
func Solve() {
	panic("implemente a solução")
}
`, pkg)
	testSrc := fmt.Sprintf(`package %s

import "testing"

func TestExample(t *testing.T) {
	t.Fatal("substitua este teste pelo exemplo oficial do LeetCode")
}
`, pkg)
	files := map[string]string{
		"solution.go":      solution,
		"solution_test.go": testSrc,
		"README.md":        renderProblemReadme(p, StatusAttempting),
	}
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
			return err
		}
	}
	return nil
}

func renderProblemReadme(p Problem, status string) string {
	tags := strings.Join(p.Tags, ", ")
	if tags == "" {
		tags = "—"
	}
	return fmt.Sprintf(`# %d. %s

- Dificuldade: %s
- Tags: %s
- Link: %s
- Status: %s

## Notas

## Resultado

- Abordagem:
- Complexidade:
`, p.ID, p.Title, p.Difficulty, tags, p.URL, StatusLabel(status))
}

// notesReady recusa o README enquanto abordagem ou complexidade seguem vazias.
func notesReady(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("preencha Abordagem e Complexidade no README da solução antes de marcar como resolvido")
		}
		return err
	}
	var missing []string
	if !noteFilled(approachLine, string(data)) {
		missing = append(missing, "Abordagem")
	}
	if !noteFilled(complexityLine, string(data)) {
		missing = append(missing, "Complexidade")
	}
	if len(missing) == 0 {
		return nil
	}
	return fmt.Errorf("preencha %s no README da solução antes de marcar como resolvido", strings.Join(missing, " e "))
}

func noteFilled(re *regexp.Regexp, text string) bool {
	match := re.FindStringSubmatch(text)
	return match != nil && strings.TrimSpace(match[1]) != ""
}

func updateStatusLine(path, status string) error {
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	if !statusLine.Match(data) {
		return nil
	}
	line := "- Status: " + StatusLabel(status)
	updated := statusLine.ReplaceAll(data, []byte(line))
	return os.WriteFile(path, updated, 0o644)
}
