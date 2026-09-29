package track

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func (s *Service) catalogPath() string {
	return filepath.Join(s.Root, "catalog", "problems.json")
}

func (s *Service) loadCatalog() ([]Problem, error) {
	data, err := os.ReadFile(s.catalogPath())
	if err != nil {
		return nil, fmt.Errorf("não consegui ler o catálogo: %w", err)
	}
	var problems []Problem
	if err := json.Unmarshal(data, &problems); err != nil {
		return nil, fmt.Errorf("catalog/problems.json inválido: %w", err)
	}
	if err := validateCatalog(problems); err != nil {
		return nil, err
	}
	sort.Slice(problems, func(i, j int) bool {
		return problems[i].ID < problems[j].ID
	})
	return problems, nil
}

func validateCatalog(problems []Problem) error {
	if len(problems) == 0 {
		return fmt.Errorf("o catálogo está vazio")
	}
	seenID := make(map[int]struct{}, len(problems))
	seenSlug := make(map[string]struct{}, len(problems))
	for _, p := range problems {
		if p.ID <= 0 {
			return fmt.Errorf("id inválido no catálogo: %d", p.ID)
		}
		if _, ok := seenID[p.ID]; ok {
			return fmt.Errorf("id duplicado no catálogo: %d", p.ID)
		}
		seenID[p.ID] = struct{}{}
		if strings.TrimSpace(p.Title) == "" {
			return fmt.Errorf("problema %d sem título", p.ID)
		}
		if !slugPattern.MatchString(p.Slug) {
			return fmt.Errorf("slug inválido no problema %d: %q", p.ID, p.Slug)
		}
		if _, ok := seenSlug[p.Slug]; ok {
			return fmt.Errorf("slug duplicado no catálogo: %s", p.Slug)
		}
		seenSlug[p.Slug] = struct{}{}
		switch p.Difficulty {
		case DifficultyEasy, DifficultyMedium, DifficultyHard:
		default:
			return fmt.Errorf("dificuldade inválida no problema %d: %q", p.ID, p.Difficulty)
		}
		if !strings.HasPrefix(p.URL, "https://leetcode.com/problems/") || !strings.HasSuffix(p.URL, "/") {
			return fmt.Errorf("url inválida no problema %d: %q", p.ID, p.URL)
		}
	}
	return nil
}

func findProblem(problems []Problem, id int) (Problem, bool) {
	for _, p := range problems {
		if p.ID == id {
			return p, true
		}
	}
	return Problem{}, false
}
