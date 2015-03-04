package picket_test

import (
	"testing"

	"github.com/ryanfaerman/picket/picket"
)

func generateBlacklist() *picket.Blacklist {
	list := picket.NewBlacklist()
	list.Add("*.xxx")
	list.Add("*.sex")
	list.Add("doubleclick.net")
	list.Add("porn.com")

	return list
}

func TestBlacklistIsAllowed(t *testing.T) {
	t.Parallel()

	list := generateBlacklist()

	for host, ok := range map[string]bool{
		"bar.spam.xxx.":    false,
		"doubleclick.net.": false,
		"example.com.":     true,
		"foo.sex.":         false,
		"github.com.":      true,
		"google.com.":      true,
		"porn.com.":        false,
	} {
		if act := list.IsAllowed(host); ok != act {
			t.Errorf("expected s.IsAllowed(%q) to be %v, got %v", host, ok, act)
		}
	}
}

func BenchmarkBlacklistIsAllowed(b *testing.B) {
	list := generateBlacklist()
	hosts := map[string]bool{
		"bar.spam.xxx.":    false,
		"doubleclick.net.": false,
		"example.com.":     true,
		"foo.sex.":         false,
		"github.com.":      true,
		"google.com.":      true,
		"porn.com.":        false,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for host, _ := range hosts {
			list.IsAllowed(host)
		}
	}

}
