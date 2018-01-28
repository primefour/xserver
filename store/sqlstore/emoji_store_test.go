package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestEmojiStore(t *testing.T) {
	StoreTest(t, storetest.TestEmojiStore)
}
