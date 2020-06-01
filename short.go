package short

import (
	"fmt"
	"github.com/googege/short/bloom"
	"github.com/googege/short/hash"
	qr "github.com/skip2/go-qrcode"
	"path/filepath"
	"sync"
)

type Node struct {
	lock       *sync.Mutex
	bf         *bloom.BloomFilter
	LongValue  string // long string
	ShortValue string // short string
}

func NewShort() *Node {
	return &Node{
		lock: &sync.Mutex{},
		bf:   bloom.NewBloom(),
	}
}

// short the long url to a small one in hash method.
func (n *Node) ShortAdd(longURL string, db DB) (string, error) {
	n.lock.Lock()
	defer n.lock.Unlock()
	murmurStringValue := hash.MurmurHash(longURL)
	n.LongValue = longURL
	n.ShortValue = murmurStringValue
	n.bf.Add([]byte(n.ShortValue))
	return db.Add(n)
}

// delete the shortURL,and it's long url
func (n *Node) ShortDelete(shortURL string, db DB) {
	n.lock.Lock()
	db.Delete(n.LongValue)
	n.lock.Unlock()
}

// find the longurl's shortURL.
func (n *Node) ShortFind(shortURL string, db DB) (string, error) {
	n.lock.Lock()
	defer n.lock.Unlock()
	_, short := filepath.Split(shortURL)
	// bloom filter
	if !n.bf.IsExit([]byte(short)) {
		return "", fmt.Errorf("bloom filter:no long url.")
	}
	return db.Find(short)
}

// change a new short URL.
func (n *Node) ShortChange(newShortURL string, db DB) error {
	n.lock.Lock()
	defer n.lock.Unlock()
	return db.Change(n, newShortURL)
}

// set shortURL to a QR CODE
func (n *Node) SetQR(url string, size int, file string) (err error) {
	n.lock.Lock()
	defer n.lock.Unlock()
	return qr.WriteFile(url+n.ShortValue, qr.Medium, size, file)
}
