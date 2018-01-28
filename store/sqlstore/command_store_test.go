package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestCommandStore(t *testing.T) {
	StoreTest(t, storetest.TestCommandStore)
}
