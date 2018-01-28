package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestSystemStore(t *testing.T) {
	StoreTest(t, storetest.TestSystemStore)
}
