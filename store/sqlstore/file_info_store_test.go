package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestFileInfoStore(t *testing.T) {
	StoreTest(t, storetest.TestFileInfoStore)
}
