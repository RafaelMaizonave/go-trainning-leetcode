package track

import (
	"path/filepath"
	"runtime"
	"testing"
)

func moduleRoot(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller falhou")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(file), "..", ".."))
}

func TestRealCatalog(t *testing.T) {
	svc := New(moduleRoot(t))
	problems, err := svc.loadCatalog()
	if err != nil {
		t.Fatal(err)
	}

	want := []int{
		1, 2, 3, 4, 7, 9, 11, 13, 14, 15, 17, 19, 20, 21, 23, 24, 25, 26, 27, 28,
		33, 34, 35, 48, 53, 58, 61, 62, 63, 64, 66, 67, 69, 70, 75, 76, 77, 79, 80, 82,
		83, 86, 88, 92, 94, 100, 101, 102, 104, 107, 111, 112, 120, 121, 122, 125, 136, 144, 148, 150,
		153, 155, 165, 167, 179, 198, 200, 203, 206, 208, 209, 211, 215, 217, 219, 226, 235, 236, 237, 257,
		258, 283, 300, 303, 304, 307, 328, 343, 344, 345, 347, 349, 350, 376, 392, 404, 416, 435, 437, 438,
		447, 454, 455, 557, 674, 677, 704, 713, 717, 728, 735, 747, 872, 1021,
	}
	if len(problems) != len(want) {
		t.Fatalf("catálogo tem %d problemas, esperado %d", len(problems), len(want))
	}
	for i, id := range want {
		if problems[i].ID != id {
			t.Fatalf("posição %d: id %d, esperado %d", i, problems[i].ID, id)
		}
	}

	hard := map[int]struct{}{4: {}, 23: {}, 25: {}, 76: {}}
	for _, p := range problems {
		if _, ok := hard[p.ID]; ok && p.Difficulty != DifficultyHard {
			t.Errorf("problema %d deveria ser Hard, veio %s", p.ID, p.Difficulty)
		}
		if p.Difficulty == DifficultyHard {
			if _, ok := hard[p.ID]; !ok {
				t.Errorf("problema %d marcado como Hard fora da lista de referência", p.ID)
			}
		}
		if PackageName(p.Slug) == "" {
			t.Errorf("slug %q não gera pacote", p.Slug)
		}
	}
}
