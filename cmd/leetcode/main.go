package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"go-trainning-leetcode/internal/track"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "erro: %s\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 || args[0] == "help" || args[0] == "-h" || args[0] == "--help" {
		usage()
		if len(args) == 0 {
			return fmt.Errorf("informe um comando")
		}
		return nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	root, err := track.FindRoot(wd)
	if err != nil {
		return err
	}
	svc := track.New(root)

	switch args[0] {
	case "list":
		return cmdList(svc, args[1:])
	case "stats":
		return cmdStats(svc, args[1:])
	case "start":
		return cmdID("tentando", svc.Start, args[1:])
	case "done":
		return cmdID("resolvido", svc.Done, args[1:])
	case "reset":
		return cmdID("a fazer", svc.Reset, args[1:])
	case "review":
		return cmdReview(svc, args[1:])
	case "readme":
		if len(args) != 1 {
			return fmt.Errorf("readme não recebe argumentos")
		}
		if err := svc.WriteReadme(); err != nil {
			return err
		}
		fmt.Println("README.md atualizado")
		return nil
	default:
		usage()
		return fmt.Errorf("comando desconhecido %q", args[0])
	}
}

func cmdList(svc *track.Service, args []string) error {
	fs := flag.NewFlagSet("list", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	difficulty := fs.String("difficulty", "", "easy, medium ou hard")
	status := fs.String("status", "", "todo, attempting ou solved")
	tag := fs.String("tag", "", "trecho de uma tag")
	fs.Usage = func() {
		fmt.Fprintln(os.Stderr, "uso: leetcode list [--difficulty easy|medium|hard] [--status todo|attempting|solved] [--tag texto]")
		fs.PrintDefaults()
	}
	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return nil
		}
		return err
	}
	if fs.NArg() != 0 {
		return fmt.Errorf("argumento inesperado %q", fs.Arg(0))
	}

	rows, err := svc.List(track.Filter{
		Difficulty: *difficulty,
		Status:     *status,
		Tag:        *tag,
	})
	if err != nil {
		return err
	}
	printRows(rows)
	return nil
}

func cmdStats(svc *track.Service, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("stats não recebe argumentos")
	}
	stats, err := svc.Stats()
	if err != nil {
		return err
	}
	fmt.Print(track.FormatStats(stats))
	return nil
}

func cmdReview(svc *track.Service, args []string) error {
	if len(args) == 0 {
		items, err := svc.ReviewQueue()
		if err != nil {
			return err
		}
		printReview(items)
		return nil
	}
	if len(args) != 1 {
		return fmt.Errorf("informe no máximo um id")
	}
	id, err := strconv.Atoi(strings.TrimSpace(args[0]))
	if err != nil || id <= 0 {
		return fmt.Errorf("id inválido %q", args[0])
	}
	if err := svc.MarkReviewed(id); err != nil {
		return err
	}
	fmt.Printf("problema %d revisado\n", id)
	return nil
}

func printReview(items []track.ReviewItem) {
	if len(items) == 0 {
		fmt.Println("nenhum problema para revisar")
		return
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tTítulo\tResolvido em\tDesde")
	for _, item := range items {
		fmt.Fprintf(w, "%04d\t%s\t%s\t%s\n",
			item.Problem.ID,
			item.Problem.Title,
			item.SolvedAt,
			item.DueSince,
		)
	}
	_ = w.Flush()
}

func cmdID(label string, fn func(int) error, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("informe o id do problema")
	}
	id, err := strconv.Atoi(strings.TrimSpace(args[0]))
	if err != nil || id <= 0 {
		return fmt.Errorf("id inválido %q", args[0])
	}
	if err := fn(id); err != nil {
		return err
	}
	fmt.Printf("problema %d marcado como %s\n", id, label)
	return nil
}

func printRows(rows []track.Row) {
	if len(rows) == 0 {
		fmt.Println("nenhum problema encontrado")
		return
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tTítulo\tDificuldade\tTags\tStatus")
	for _, row := range rows {
		tags := strings.Join(row.Problem.Tags, ", ")
		if tags == "" {
			tags = "-"
		}
		fmt.Fprintf(w, "%04d\t%s\t%s\t%s\t%s\n",
			row.Problem.ID,
			row.Problem.Title,
			row.Problem.Difficulty,
			tags,
			track.StatusLabel(row.Status),
		)
	}
	_ = w.Flush()
}

func usage() {
	fmt.Fprintf(os.Stderr, `uso:
  leetcode list [--difficulty easy|medium|hard] [--status todo|attempting|solved] [--tag texto]
  leetcode stats
  leetcode start <id>
  leetcode done <id>
  leetcode reset <id>
  leetcode review [id]
  leetcode readme
`)
}
