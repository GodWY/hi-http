package logging

import "testing"

func TestLogger(t *testing.T) {
	tl := Default()
	tl.MustAppend().Int("w", 1).String("nam", "w")
	tl.Flush()
}
