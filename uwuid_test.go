package uwuid

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	ids := make(map[string]bool)

	for i := 0; i < 100; i++ {
		id := New()
		if len(id) == 0 {
			t.Errorf("New() returned an empty string")
		}

		segments := strings.Split(id, "-")
		if len(segments) != 6 {
			t.Errorf("Expected 6 segments separated by '-', got %d in %q", len(segments), id)
		}

		for _, segment := range segments {
			valid := false
			for _, token := range uwuids {
				if strings.Contains(segment, token) {
					valid = true
					break
				}
			}
			if !valid {
				t.Errorf("Segment %q contains no valid uwuid tokens in %q", segment, id)
			}
		}

		if ids[id] {
			continue
		}

		ids[id] = true
	}

	if len(ids) < 10 {
		t.Errorf("Expected at least 10 unique IDs out of 100, got %d", len(ids))
	}
}
