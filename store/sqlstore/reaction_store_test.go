package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestReactionStore(t *testing.T) {
	StoreTest(t, storetest.TestReactionStore)
}
