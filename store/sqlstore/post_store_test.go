package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestPostStore(t *testing.T) {
	StoreTest(t, storetest.TestPostStore)
}
