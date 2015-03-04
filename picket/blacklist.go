package picket

import (
	"crypto/md5"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"
)

type checksum [md5.Size / 2]byte
type hosts map[checksum]struct{}
type hostsRX map[checksum]*regexp.Regexp

type Blacklist struct {
	source string
	poll   time.Duration

	mutex   sync.RWMutex
	hosts   hosts
	hostsRX hostsRX
}

func NewBlacklist() *Blacklist {
	return &Blacklist{
		hosts:   hosts{},
		hostsRX: hostsRX{},
	}
}

func NewRemoteBlacklist(src string) *Blacklist {
	b := NewBlacklist()
	b.source = src
	b.Load()
	return b
}

func (b *Blacklist) IsRemote() bool {
	return b.source != ""
}

// IsAllowed returns whether we are allowed access the host
//
// NOTE: "host" must end with a dot.
func (b *Blacklist) IsAllowed(host string) bool {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	_, ok := b.hosts[hash(host)]
	if ok {
		return false
	}

	for _, rx := range b.hostsRX {
		if rx.MatchString(host) {
			return false
		}
	}

	return true
}

func (b *Blacklist) Add(host string) {
	if host == "" {
		return
	}

	// Force all domains to be absolute
	if host[len(host)-1] != '.' {
		host += "."
	}

	if !strings.ContainsRune(host, '*') {
		// Plain host string:
		b.mutex.Lock()
		b.hosts[hash(host)] = struct{}{}
		b.mutex.Unlock()
	} else if rx := compilePattern(host); rx != nil {
		// Host pattern (regex):
		b.mutex.Lock()
		b.hostsRX[hash(rx.String())] = rx
		b.mutex.Unlock()
	}
}

func (b *Blacklist) Remove(host string) {
	if host == "" {
		return
	}
	if host[len(host)-1] != '.' {
		host += "."
	}

	if !strings.ContainsRune(host, '*') {
		// Plain host string:
		b.mutex.Lock()
		delete(b.hosts, hash(host))
		b.mutex.Unlock()
	} else if rx := compilePattern(host); rx != nil {
		// Host pattern (regex):
		b.mutex.Lock()
		delete(b.hostsRX, hash(rx.String()))
		b.mutex.Unlock()
	}
}

// Count how many hosts are in the blacklist
func (b *Blacklist) Count() int {
	b.mutex.Lock()
	n := len(b.hosts) + len(b.hostsRX)
	b.mutex.Unlock()
	return n
}

// Load retrieves the contents of a remote blacklist.
func (b *Blacklist) Load() error {
	// TODO: implement this
	return nil
}

// Monitor updates the remote blacklist every "poll".
// You probably want to run this in a go routine.
func (b *Blacklist) Monitor(poll time.Duration) {
	// TODO: implement this
}

func compilePattern(pat string) *regexp.Regexp {
	pat = strings.Replace(pat, ".", `\.`, -1)
	pat = strings.Replace(pat, "*", ".*", -1)
	pat = "^" + pat + `$`
	rx, err := regexp.Compile(pat)
	if err != nil {
		log.Printf("picket: could not compile %q: %s", pat, err)
		return nil
	}
	return rx
}

func hash(s string) checksum {
	sum := md5.Sum([]byte(s))
	return checksum{
		sum[0] ^ sum[1],
		sum[2] ^ sum[3],
		sum[4] ^ sum[5],
		sum[6] ^ sum[7],
		sum[8] ^ sum[9],
		sum[10] ^ sum[11],
		sum[12] ^ sum[13],
		sum[14] ^ sum[15],
	}
}
