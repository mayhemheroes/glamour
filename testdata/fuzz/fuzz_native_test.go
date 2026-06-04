package fuzzing

import (
	"testing"

	"charm.land/glamour/v2"
)

func FuzzRenderBytes(f *testing.F) {
	f.Add([]byte("# Hello\n"))
	f.Fuzz(func(t *testing.T, data []byte) {
		_, _ = glamour.RenderBytes(data, "dark")
	})
}
