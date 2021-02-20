package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

var gt = flag.Float64("gt", 0, "Print strings with entropy greater than float")
var lt = flag.Float64("lt", 0, "Print strings with entropy less than float")

func main() {
	flag.Parse()

	var ents entropies

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		ents = append(ents, entropy{
			str: s.Text(),
			ent: shannon(s.Text()),
		})
	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}

	format := "%f\t%s\n"
	sort.Sort(ents)

	if *gt != 0 {
		for _, e := range ents {
			if e.ent > *gt {
				fmt.Printf(format, e.ent, e.str)
			}
		}
		os.Exit(0)
	}

	if *lt != 0 {
		for _, e := range ents {
			if e.ent < *lt {
				fmt.Printf(format, e.ent, e.str)
			}
		}
		os.Exit(0)
	}

	for _, e := range ents {
		fmt.Printf(format, e.ent, e.str)
	}
}

type entropy struct {
	str string
	ent float64
}

type entropies []entropy

func (e entropies) Len() int           { return len(e) }
func (e entropies) Less(i, j int) bool { return e[i].ent < e[j].ent }
func (e entropies) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

// https://rosettacode.org/wiki/Entropy#Go
func shannon(data string) (entropy float64) {
	if data == "" {
		return 0
	}
	for i := 0; i < 256; i++ {
		px := float64(strings.Count(data, string(byte(i)))) / float64(len(data))
		if px > 0 {
			entropy += -px * math.Log2(px)
		}
	}
	return entropy
}
