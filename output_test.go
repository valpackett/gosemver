package gosemver

import (
	"testing"
)

type outTestCase struct {
	Input  Version
	Output string
}

var outTests = []outTestCase{
	{Version{"", 0, 1, 0, "", ""}, "0.1.0"},
	{Version{"v", 0, 1, 0, "", ""}, "v0.1.0"},
	{Version{"v", 0, 1, 0, "alpha", ""}, "v0.1.0-alpha"},
	{Version{"v", 0, 1, 0, "alpha", "myBuild"}, "v0.1.0-alpha+myBuild"},
	{Version{"", 64, 128, 256, "", "build.1.2.3"}, "64.128.256+build.1.2.3"},
}

func TestOutput(t *testing.T) {
	for _, tcase := range outTests {
		result := tcase.Input.String()
		if result != tcase.Output {
			t.Error(
				"For", tcase.Input,
				"expected", tcase.Output,
				"got", result,
			)
		}
	}
}
