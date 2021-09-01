package magnetinfo

import (
	"testing"
)

func Test_Parse(t *testing.T) {
	magnets := []string{
		"magnet:?xt=urn:btih:OVFYDTCS55I7T5JLQFZA2NH6LKZI5XDM",

		"magnet:?xt=urn:btih:83918ea4bb488cefd3d8b8b8762597d32aebb4fa&tr=http%3A%2F%2Facademictorrents.com%2Fannounce.php&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969",

		"magnet:?xt=urn:btih:e8b1f9c5bf555fe58bc73addb83457dd6da69630&tr=http%3A%2F%2Facademictorrents.com%2Fannounce.php&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969",
	}

	for _, magnet := range magnets {
		_, err := Parse(magnet)
		if err != nil {
			t.Errorf("Parse failed: %s", err)
		}
	}
}
