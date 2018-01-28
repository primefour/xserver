package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestStatusStore(t *testing.T) {
	StoreTest(t, storetest.TestStatusStore)
}
