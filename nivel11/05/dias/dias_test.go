package dias

import "testing"

type testpair struct{
	values string
	dia string
}

var diasSemana = []testpair{
	{"lunes", "lunes"},
	{"martes", "martes"},
	{"miercoles", "miercoles"},
	{"jueves", "jueves"},
	{"viernes", "viernes"},
	{"sabado", "sabado"},
	{"domingo", "domingo"},
}

func TestDia(t *testing.T) {
	for _, pair := range diasSemana{
		v := Dia(pair.values)
		if v != pair.dia {
			t.Error(
				"For", pair.values,
				"expected", pair.dia,
				"got", v,
			)
		}
	}
		
}
