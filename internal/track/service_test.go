package track

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func fixedNow() time.Time {
	return time.Date(2026, 9, 29, 15, 0, 0, 0, time.UTC)
}

func sampleProblems() []Problem {
	return []Problem{
		{ID: 1, Title: "Two Sum", Slug: "two_sum", Difficulty: DifficultyEasy, Tags: []string{"array", "lookup table"}, URL: "https://leetcode.com/problems/two-sum/"},
		{ID: 2, Title: "Add Two Numbers", Slug: "add_two_numbers", Difficulty: DifficultyMedium, Tags: []string{"linked list"}, URL: "https://leetcode.com/problems/add-two-numbers/"},
		{ID: 4, Title: "Median of Two Sorted Arrays", Slug: "median_of_two_sorted_arrays", Difficulty: DifficultyHard, Tags: []string{"binary search"}, URL: "https://leetcode.com/problems/median-of-two-sorted-arrays/"},
	}
}

func newFixture(t *testing.T) *Service {
	t.Helper()
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "catalog"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(root, "progress"), 0o755); err != nil {
		t.Fatal(err)
	}
	data, err := json.MarshalIndent(sampleProblems(), "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "catalog", "problems.json"), append(data, '\n'), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "progress", "progress.json"), []byte("{}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	svc := New(root)
	svc.Now = fixedNow
	svc.TestDir = func(string) error { return nil }
	return svc
}

func fillNotes(t *testing.T, svc *Service, folder string) {
	t.Helper()
	path := filepath.Join(svc.Root, "solutions", folder, "README.md")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	text := strings.Replace(string(data), "- Abordagem:", "- Abordagem: mapa", 1)
	text = strings.Replace(text, "- Complexidade:", "- Complexidade: O(n)", 1)
	if err := os.WriteFile(path, []byte(text), 0o644); err != nil {
		t.Fatal(err)
	}
}

func readProgress(t *testing.T, svc *Service) map[string]Entry {
	t.Helper()
	prog, err := svc.loadProgress()
	if err != nil {
		t.Fatal(err)
	}
	return prog
}

func TestPackageName(t *testing.T) {
	cases := map[string]string{
		"two_sum":                    "twosum",
		"3sum":                       "p3sum",
		"4sum2":                      "p4sum2",
		"1_bit_and_2_bit_characters": "p1bitand2bitcharacters",
	}
	for slug, want := range cases {
		if got := PackageName(slug); got != want {
			t.Errorf("PackageName(%q) = %q, want %q", slug, got, want)
		}
	}
}

func TestFormatStatsBar(t *testing.T) {
	text := FormatStats(Stats{
		Solved: 1,
		Total:  8,
		Todo:   7,
		ByDifficulty: []DifficultyStat{
			{Difficulty: DifficultyEasy, Solved: 1, Total: 8},
			{Difficulty: DifficultyMedium, Solved: 0, Total: 0},
			{Difficulty: DifficultyHard, Solved: 0, Total: 0},
		},
	})
	if !strings.Contains(text, "█░░░░░░░░░") {
		t.Fatalf("barra parcial ausente:\n%s", text)
	}
	if !strings.Contains(text, "░░░░░░░░░░") {
		t.Fatalf("barra vazia ausente:\n%s", text)
	}
}

func TestStartDoneReset(t *testing.T) {
	svc := newFixture(t)

	if err := svc.Start(99); err == nil {
		t.Fatal("id fora do catálogo deveria falhar")
	}
	if err := svc.Done(1); err == nil {
		t.Fatal("done sem pasta deveria falhar")
	}
	if prog := readProgress(t, svc); len(prog) != 0 {
		t.Fatalf("progresso mudou sem start: %+v", prog)
	}

	if err := svc.Start(1); err != nil {
		t.Fatal(err)
	}
	if err := svc.Start(1); err == nil {
		t.Fatal("start não deveria sobrescrever a pasta")
	}

	prog := readProgress(t, svc)
	if prog["1"].Status != StatusAttempting || prog["1"].StartedAt != "2026-09-29" || prog["1"].SolvedAt != "" {
		t.Fatalf("progresso após start: %+v", prog["1"])
	}
	solution := filepath.Join(svc.Root, "solutions", "0001_two_sum", "solution.go")
	if _, err := os.Stat(solution); err != nil {
		t.Fatal(err)
	}
	body, err := os.ReadFile(filepath.Join(svc.Root, "solutions", "0001_two_sum", "solution.go"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), "package twosum") {
		t.Fatalf("pacote inesperado:\n%s", body)
	}
	readme, err := os.ReadFile(filepath.Join(svc.Root, "README.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(readme), "Tentando") || !strings.Contains(string(readme), "solutions/0001_two_sum/solution.go") {
		t.Fatalf("readme após start:\n%s", readme)
	}
	if !strings.Contains(string(readme), "Easy") || !strings.Contains(string(readme), "Medium") || !strings.Contains(string(readme), "Hard") {
		t.Fatal("readme sem as três dificuldades")
	}
	if !strings.Contains(string(readme), "## Para revisar") || !strings.Contains(string(readme), "Nenhum problema para revisar.") {
		t.Fatalf("fila de revisão ausente:\n%s", readme)
	}

	if err := svc.Done(1); err == nil || !strings.Contains(err.Error(), "Abordagem") || !strings.Contains(err.Error(), "Complexidade") {
		t.Fatalf("done sem caderno: %v", err)
	}
	if readProgress(t, svc)["1"].Status != StatusAttempting {
		t.Fatal("caderno vazio não pode marcar como resolvido")
	}
	fillNotes(t, svc, "0001_two_sum")

	svc.TestDir = func(string) error { return fmt.Errorf("os testes falharam") }
	if err := svc.Done(1); err == nil {
		t.Fatal("done deveria recusar teste que falha")
	}
	if readProgress(t, svc)["1"].Status != StatusAttempting {
		t.Fatal("falha de teste não pode marcar como resolvido")
	}

	svc.TestDir = func(string) error { return nil }
	if err := svc.Done(1); err != nil {
		t.Fatal(err)
	}
	prog = readProgress(t, svc)
	if prog["1"].Status != StatusSolved || prog["1"].StartedAt != "2026-09-29" || prog["1"].SolvedAt != "2026-09-29" || prog["1"].ReviewedAt != "" {
		t.Fatalf("progresso após done: %+v", prog["1"])
	}
	problemReadme, err := os.ReadFile(filepath.Join(svc.Root, "solutions", "0001_two_sum", "README.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(problemReadme), "- Status: Resolvido") {
		t.Fatalf("readme do problema:\n%s", problemReadme)
	}
	rootReadme, err := os.ReadFile(filepath.Join(svc.Root, "README.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(rootReadme), "Resolvidos: 1/3") || !strings.Contains(string(rootReadme), "██████████") {
		t.Fatalf("painel após done:\n%s", rootReadme)
	}

	if err := svc.Reset(1); err != nil {
		t.Fatal(err)
	}
	if _, ok := readProgress(t, svc)["1"]; ok {
		t.Fatal("reset deveria remover o progresso")
	}
	if _, err := os.Stat(solution); err != nil {
		t.Fatal("reset apagou a solução")
	}
	problemReadme, err = os.ReadFile(filepath.Join(svc.Root, "solutions", "0001_two_sum", "README.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(problemReadme), "- Status: A fazer") {
		t.Fatalf("status após reset:\n%s", problemReadme)
	}
}

func TestListFilters(t *testing.T) {
	svc := newFixture(t)
	if err := svc.Start(1); err != nil {
		t.Fatal(err)
	}

	rows, err := svc.List(Filter{Difficulty: "EASY", Status: "attempting", Tag: "Array"})
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 || rows[0].Problem.ID != 1 || rows[0].Status != StatusAttempting {
		t.Fatalf("filtro easy/attempting/array: %+v", rows)
	}

	rows, err = svc.List(Filter{Status: "todo", Difficulty: "hard"})
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 || rows[0].Problem.ID != 4 {
		t.Fatalf("filtro hard/todo: %+v", rows)
	}

	rows, err = svc.List(Filter{Tag: "linked"})
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 || rows[0].Problem.ID != 2 {
		t.Fatalf("filtro tag: %+v", rows)
	}

	rows, err = svc.List(Filter{Status: "solved"})
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 0 {
		t.Fatalf("não deveria haver resolvidos: %+v", rows)
	}

	if _, err := svc.List(Filter{Difficulty: "nope"}); err == nil {
		t.Fatal("dificuldade inválida deveria falhar")
	}
	if _, err := svc.List(Filter{Status: "done"}); err == nil {
		t.Fatal("status inválido deveria falhar")
	}
}

func TestStats(t *testing.T) {
	svc := newFixture(t)
	stats, err := svc.Stats()
	if err != nil {
		t.Fatal(err)
	}
	if stats.Total != 3 || stats.Todo != 3 || stats.Solved != 0 || stats.Attempting != 0 {
		t.Fatalf("stats iniciais: %+v", stats)
	}
	if len(stats.ByDifficulty) != 3 {
		t.Fatalf("dificuldades: %+v", stats.ByDifficulty)
	}
	if err := svc.Start(2); err != nil {
		t.Fatal(err)
	}
	stats, err = svc.Stats()
	if err != nil {
		t.Fatal(err)
	}
	if stats.Attempting != 1 || stats.Todo != 2 {
		t.Fatalf("stats após start: %+v", stats)
	}
	if stats.ByDifficulty[1].Difficulty != DifficultyMedium || stats.ByDifficulty[1].Total != 1 || stats.ByDifficulty[1].Solved != 0 {
		t.Fatalf("medium: %+v", stats.ByDifficulty[1])
	}
}

func TestDoneUsesGoTest(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "catalog"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(root, "progress"), 0o755); err != nil {
		t.Fatal(err)
	}
	data, err := json.MarshalIndent(sampleProblems()[:1], "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "catalog", "problems.json"), append(data, '\n'), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "progress", "progress.json"), []byte("{}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "go.mod"), []byte("module example.com/tmp\n\ngo 1.22\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	svc := New(root)
	svc.Now = fixedNow
	if err := svc.Start(1); err != nil {
		t.Fatal(err)
	}
	fillNotes(t, svc, "0001_two_sum")
	if err := svc.Done(1); err == nil {
		t.Fatal("o teste de exemplo deveria falhar")
	}
	if readProgress(t, svc)["1"].Status != StatusAttempting {
		t.Fatal("teste falho marcou o problema como resolvido")
	}

	passing := "package twosum\n\nimport \"testing\"\n\nfunc TestExample(t *testing.T) {}\n"
	testPath := filepath.Join(root, "solutions", "0001_two_sum", "solution_test.go")
	if err := os.WriteFile(testPath, []byte(passing), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := svc.Done(1); err != nil {
		t.Fatal(err)
	}
	entry := readProgress(t, svc)["1"]
	if entry.Status != StatusSolved || entry.SolvedAt != "2026-09-29" {
		t.Fatalf("done com teste real: %+v", entry)
	}
}

func TestReviewQueue(t *testing.T) {
	svc := newFixture(t)
	if err := svc.saveProgress(map[string]Entry{
		"1": {Status: StatusSolved, StartedAt: "2026-09-22", SolvedAt: "2026-09-22"},
		"2": {Status: StatusSolved, StartedAt: "2026-09-28", SolvedAt: "2026-09-28"},
		"4": {Status: StatusAttempting, StartedAt: "2026-09-01"},
	}); err != nil {
		t.Fatal(err)
	}

	items, err := svc.ReviewQueue()
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 1 || items[0].Problem.ID != 1 || items[0].DueSince != "2026-09-29" {
		t.Fatalf("fila no dia 7: %+v", items)
	}

	readme, err := svc.RenderReadme()
	if err != nil {
		t.Fatal(err)
	}
	section, _, _ := strings.Cut(readme, "## Problemas")
	if !strings.Contains(section, "## Para revisar") || !strings.Contains(section, "0001") || strings.Contains(section, "0002") {
		t.Fatalf("seção para revisar:\n%s", section)
	}

	if err := svc.MarkReviewed(4); err == nil {
		t.Fatal("não deveria revisar um problema tentando")
	}
	if err := svc.MarkReviewed(2); err != nil {
		t.Fatal(err)
	}
	if readProgress(t, svc)["2"].ReviewedAt != "2026-09-29" {
		t.Fatalf("revisão antecipada: %+v", readProgress(t, svc)["2"])
	}
	items, err = svc.ReviewQueue()
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 1 || items[0].Problem.ID != 1 {
		t.Fatalf("revisão do problema 2 não tira o 1 da fila: %+v", items)
	}
	if err := svc.MarkReviewed(1); err != nil {
		t.Fatal(err)
	}
	if readProgress(t, svc)["1"].ReviewedAt != "2026-09-29" {
		t.Fatalf("reviewed_at: %+v", readProgress(t, svc)["1"])
	}
	items, err = svc.ReviewQueue()
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 0 {
		t.Fatalf("revisão do marco de 7 dias deveria esvaziar a fila: %+v", items)
	}
}

func TestReviewMilestones(t *testing.T) {
	today := time.Date(2026, 9, 29, 0, 0, 0, 0, time.UTC)

	if _, due := reviewDueSince("2026-09-28", "", today); due {
		t.Fatal("um dia depois de resolver não entra na fila")
	}
	if since, due := reviewDueSince("2026-09-22", "", today); !due || since != "2026-09-29" {
		t.Fatalf("marco de 7 dias: due=%v since=%s", due, since)
	}
	if _, due := reviewDueSince("2026-09-22", "2026-09-29", today); due {
		t.Fatal("revisão no marco de 7 dias não pode continuar vencida")
	}
	if since, due := reviewDueSince("2026-09-08", "2026-09-15", today); !due || since != "2026-09-29" {
		t.Fatalf("marco de 21 dias: due=%v since=%s", due, since)
	}
	if _, due := reviewDueSince("2026-09-08", "2026-09-29", today); due {
		t.Fatal("revisão no marco de 21 dias não pode continuar vencida")
	}
	if since, due := reviewDueSince("2026-07-31", "2026-08-21", today); !due || since != "2026-09-29" {
		t.Fatalf("marco de 60 dias: due=%v since=%s", due, since)
	}
	if _, due := reviewDueSince("2026-07-01", "2026-09-28", today); due {
		t.Fatal("ciclo seguinte ainda não completou 7 dias")
	}
	if since, due := reviewDueSince("2026-07-01", "2026-09-22", today); !due || since != "2026-09-29" {
		t.Fatalf("ciclo depois do marco de 60 dias: due=%v since=%s", due, since)
	}
}
