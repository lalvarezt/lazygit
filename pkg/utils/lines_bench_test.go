package utils

import "testing"

func BenchmarkEscapeSpecialChars(b *testing.B) {
	for _, fixture := range []struct {
		name, path string
	}{
		{"ordinary", "many/dir000/file000.txt"},
		{"controls", "many/dir000/file\twith\ncontrols.txt"},
	} {
		b.Run(fixture.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				EscapeSpecialChars(fixture.path)
			}
		})
	}
}
