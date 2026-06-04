package fuzzing

import "charm.land/glamour/v2"

func Fuzz(data []byte) int {
	_, err := glamour.RenderBytes(data, "dark")
	if err != nil {
		return 0
	}
	return 1
}
