package track

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

// Marcos a partir da data em que o problema foi resolvido.
// Depois do último marco, a revisão volta 7 dias após reviewed_at.
var reviewAfterDays = []int{7, 21, 60}

// ReviewQueue lista problemas resolvidos cujo próximo marco já chegou.
func (s *Service) ReviewQueue() ([]ReviewItem, error) {
	problems, err := s.loadCatalog()
	if err != nil {
		return nil, err
	}
	prog, err := s.loadProgress()
	if err != nil {
		return nil, err
	}
	today := civilDay(s.clock())
	items := make([]ReviewItem, 0)
	for _, p := range problems {
		entry, ok := prog[strconv.Itoa(p.ID)]
		if !ok || effectiveStatus(p.ID, prog) != StatusSolved {
			continue
		}
		dueSince, due := reviewDueSince(entry.SolvedAt, entry.ReviewedAt, today)
		if !due {
			continue
		}
		items = append(items, ReviewItem{
			Problem:     p,
			SolvedAt:    entry.SolvedAt,
			DueSince:    dueSince,
			SolutionRel: s.solutionRel(p),
		})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].DueSince != items[j].DueSince {
			return items[i].DueSince < items[j].DueSince
		}
		return items[i].Problem.ID < items[j].Problem.ID
	})
	return items, nil
}

// MarkReviewed grava a data de hoje num problema já resolvido.
func (s *Service) MarkReviewed(id int) error {
	if _, err := s.mustProblem(id); err != nil {
		return err
	}
	prog, err := s.loadProgress()
	if err != nil {
		return err
	}
	key := strconv.Itoa(id)
	entry, ok := prog[key]
	if !ok || entry.Status != StatusSolved {
		return fmt.Errorf("só é possível revisar um problema resolvido")
	}
	entry.ReviewedAt = civilDay(s.clock()).Format("2006-01-02")
	prog[key] = entry
	if err := s.saveProgress(prog); err != nil {
		return err
	}
	return s.WriteReadme()
}

// reviewDueSince devolve o marco vencido mais antigo que a última revisão não cobriu.
func reviewDueSince(solvedAt, reviewedAt string, today time.Time) (string, bool) {
	solved, err := time.Parse("2006-01-02", solvedAt)
	if err != nil {
		return "", false
	}
	solved = civilDay(solved)
	today = civilDay(today)

	var reviewed time.Time
	hasReview := false
	if reviewedAt != "" {
		parsed, parseErr := time.Parse("2006-01-02", reviewedAt)
		if parseErr != nil {
			return "", false
		}
		reviewed = civilDay(parsed)
		hasReview = true
	}

	for _, days := range reviewAfterDays {
		milestone := solved.AddDate(0, 0, days)
		if today.Before(milestone) {
			break
		}
		if !hasReview || reviewed.Before(milestone) {
			return milestone.Format("2006-01-02"), true
		}
	}

	last := solved.AddDate(0, 0, reviewAfterDays[len(reviewAfterDays)-1])
	if hasReview && !reviewed.Before(last) {
		next := reviewed.AddDate(0, 0, reviewAfterDays[0])
		if !today.Before(next) {
			return next.Format("2006-01-02"), true
		}
	}
	return "", false
}

func civilDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
