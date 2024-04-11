package generic

import (
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
