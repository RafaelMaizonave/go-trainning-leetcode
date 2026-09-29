package track

import (
	"fmt"
	"strings"
)

const readmeHeader = `# Treino LeetCode em Go

Repositório pessoal para resolver problemas e acompanhar o progresso. O catálogo (título, dificuldade e tags) segue o índice de [keep-practicing/leetcode-go](https://github.com/keep-practicing/leetcode-go). As soluções ficam nas pastas criadas por você.

Este arquivo é gerado pelo comando ` + "`go run ./cmd/leetcode readme`" + `. Os comandos de progresso também o atualizam.

## Comandos

` + "```bash" + `
go run ./cmd/leetcode list
go run ./cmd/leetcode list --difficulty easy --status todo
go run ./cmd/leetcode stats
go run ./cmd/leetcode start 1
go run ./cmd/leetcode done 1
go run ./cmd/leetcode reset 1
go run ./cmd/leetcode review
go run ./cmd/leetcode review 1
` + "```" + `

Pelo Makefile: ` + "`make list`" + `, ` + "`make start ID=1`" + `, ` + "`make done ID=1`" + `, ` + "`make reset ID=1`" + `, ` + "`make stats`" + ` e ` + "`make review`" + `.

- **A fazer**: o problema ainda não foi iniciado.
- **Tentando**: ` + "`start`" + ` criou a pasta da solução.
- **Resolvido**: ` + "`done`" + ` passou em ` + "`go test`" + ` e o README da solução tem abordagem e complexidade.
- ` + "`reset`" + ` volta o status para a fazer e mantém os arquivos da solução.
- ` + "`review`" + ` lista o que já esfriou. ` + "`review <id>`" + ` grava a data da revisão.

`

const barWidth = 10

// FormatStats formata o painel de progresso em texto.
func FormatStats(st Stats) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Resolvidos: %d/%d\n\n", st.Solved, st.Total)
	for _, item := range st.ByDifficulty {
		frac := fmt.Sprintf("%d/%d", item.Solved, item.Total)
		fmt.Fprintf(&b, "%-6s  %-7s  %s\n", item.Difficulty, frac, bar(item.Solved, item.Total))
	}
	fmt.Fprintf(&b, "\nTentando: %d\nA fazer: %d\n", st.Attempting, st.Todo)
	return b.String()
}

func bar(solved, total int) string {
	filled := 0
	if total > 0 {
		filled = solved * barWidth / total
		if solved > 0 && filled == 0 {
			filled = 1
		}
		if filled > barWidth {
			filled = barWidth
		}
	}
	return strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)
}

// RenderReadme monta o README completo a partir do catálogo e do progresso.
func (s *Service) RenderReadme() (string, error) {
	rows, err := s.List(Filter{})
	if err != nil {
		return "", err
	}
	stats, err := s.Stats()
	if err != nil {
		return "", err
	}
	due, err := s.ReviewQueue()
	if err != nil {
		return "", err
	}

	var b strings.Builder
	b.WriteString(readmeHeader)
	b.WriteString("## Progresso\n\n```\n")
	b.WriteString(FormatStats(stats))
	b.WriteString("```\n\n")
	b.WriteString(renderReviewSection(due))
	b.WriteString("## Problemas\n\n")
	b.WriteString("| ID | Título | Dificuldade | Tags | Status | Solução |\n")
	b.WriteString("| --- | --- | --- | --- | --- | --- |\n")
	for _, row := range rows {
		tags := escapeCell(strings.Join(row.Problem.Tags, ", "))
		if tags == "" {
			tags = "—"
		}
		solution := "—"
		if row.SolutionRel != "" {
			solution = fmt.Sprintf("[solução](%s)", row.SolutionRel)
		}
		title := fmt.Sprintf("[%s](%s)", escapeCell(row.Problem.Title), row.Problem.URL)
		fmt.Fprintf(&b, "| %04d | %s | %s | %s | %s | %s |\n",
			row.Problem.ID, title, row.Problem.Difficulty, tags, StatusLabel(row.Status), solution)
	}
	b.WriteString("\n")
	return b.String(), nil
}

// WriteReadme reescreve o README.md da raiz.
func (s *Service) WriteReadme() error {
	text, err := s.RenderReadme()
	if err != nil {
		return err
	}
	return osWriteFile(s.Root, "README.md", []byte(text))
}

func renderReviewSection(items []ReviewItem) string {
	var b strings.Builder
	b.WriteString("## Para revisar\n\n")
	if len(items) == 0 {
		b.WriteString("Nenhum problema para revisar.\n\n")
		return b.String()
	}
	b.WriteString("| ID | Título | Resolvido em | Desde | Solução |\n")
	b.WriteString("| --- | --- | --- | --- | --- |\n")
	for _, item := range items {
		solution := "—"
		if item.SolutionRel != "" {
			solution = fmt.Sprintf("[solução](%s)", item.SolutionRel)
		}
		title := fmt.Sprintf("[%s](%s)", escapeCell(item.Problem.Title), item.Problem.URL)
		fmt.Fprintf(&b, "| %04d | %s | %s | %s | %s |\n",
			item.Problem.ID, title, item.SolvedAt, item.DueSince, solution)
	}
	b.WriteString("\n")
	return b.String()
}

func escapeCell(s string) string {
	return strings.ReplaceAll(s, "|", "\\|")
}
