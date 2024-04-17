package generic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	s1 := []string{"A", "B", "C", "D", "E", "F"}

	s2 := Filter(s1, func(f string) bool {
		if f == "C" || f == "D" {
			return false
		}
		return true
	})

	assert.Equal(t, 4, len(s2))
	assert.Contains(t, s2, "A")
	assert.Contains(t, s2, "B")
	assert.Contains(t, s2, "E")
	assert.Contains(t, s2, "F")
}

func TestChaining(t *testing.T) {
	input1 := []string{"A", "B", "C", "D", "E", "F"}

	output := Map(
		Filter(input1, func(f string) bool {
			if f == "C" || f == "D" {
				return false
			}
			return true
		}),
		func(f string) int {
			switch f {
			case "A":
				return 1
			case "B":
				return 2
			case "C":
				return 3
			case "D":
				return 4
			case "E":
				return 5
			case "F":
				return 6
			default:
				return 0
			}
		})
	assert.Equal(t, 4, len(output))
	assert.Contains(t, output, 1)
	assert.Contains(t, output, 2)
	assert.Contains(t, output, 5)
	assert.Contains(t, output, 6)
}

func TestMap(t *testing.T) {
	type birne struct {
		farbe   string
		gewicht int
	}

	type gluehbirne struct {
		leistung int
		dimmbar  bool
	}

	b := []birne{{farbe: "gelb", gewicht: 150}, {farbe: "rot", gewicht: 200}}

	g := Map(b, func(b birne) gluehbirne {
		switch b.farbe {
		case "gelb":
			return gluehbirne{leistung: 100, dimmbar: true}
		case "rot":
			return gluehbirne{leistung: 40, dimmbar: false}
		default:
			return gluehbirne{}
		}
	})

	assert.Equal(t, 2, len(g))
	assert.Contains(t, g, gluehbirne{leistung: 100, dimmbar: true})
	assert.Contains(t, g, gluehbirne{leistung: 40, dimmbar: false})
}

func TestReduce_Sum(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5}

	res := Reduce(elements, func(zahl, sum int) int {
		return zahl + sum
	})

	assert.Equal(t, 15, res)
}

func TestReduce_Concat(t *testing.T) {
	elements := []string{"Go", "ist", "cool"}

	res := Reduce(elements, func(toAppend, str string) string {
		trenner := " "
		if str == "" {
			trenner = ""
		}
		return fmt.Sprint(str, trenner, toAppend)
	})

	assert.Equal(t, "Go ist cool", res)
}
