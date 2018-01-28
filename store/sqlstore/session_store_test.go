package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestSessionStore(t *testing.T) {
	StoreTest(t, storetest.TestSessionStore)
}
