package main

import (
	"testing"

	"github.com/nprimo/aoc_2024/internal/assert"
)

const input string = `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn
`

func TestSolve1(t *testing.T) {
	want := 7
	got := Solve1(input)
	assert.Equal(t, want, got)
}
