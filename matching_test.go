package gosemver

import (
	"testing"
)

type findAllTestCase struct {
	Input  []Version
	Constr string
	Output []Version
}

var findAllTests = []findAllTestCase{
	{
		[]Version{Version{"", 1, 1, 1, "", ""}, Version{"", 0, 5, 10, "", ""}, Version{"", 1, 2, 0, "", ""}},
		">= 1.0.0",
		[]Version{Version{"", 1, 1, 1, "", ""}, Version{"", 1, 2, 0, "", ""}},
	},
}

func TestFindAll(t *testing.T) {
	for _, tcase := range findAllTests {
		result, _ := FindAll(tcase.Input, tcase.Constr)
		for i, v := range result {
			if v != tcase.Output[i] {
				t.Error(
					"For", tcase.Input, tcase.Constr,
					"expected", tcase.Output,
					"got", result,
				)
			}
		}
	}
}

type findMaxTestCase struct {
	Input  []Version
	Constr string
	Output Version
}

var findMaxTests = []findMaxTestCase{
	{
		[]Version{Version{"", 1, 2, 0, "", ""}, Version{"", 0, 5, 10, "", ""}, Version{"", 1, 1, 0, "", ""}},
		">= 1.0.0",
		Version{"", 1, 2, 0, "", ""},
	},
	{
		[]Version{},
		">= 1.0.0",
		Version{},
	},
}

func TestFindMax(t *testing.T) {
	for _, tcase := range findMaxTests {
		result, _ := FindMax(tcase.Input, tcase.Constr)
		if result != tcase.Output {
			t.Error(
				"For", tcase.Input, tcase.Constr,
				"expected", tcase.Output,
				"got", result,
			)
		}
	}
}
