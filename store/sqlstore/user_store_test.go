package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestUserStore(t *testing.T) {
	StoreTest(t, storetest.TestUserStore)
}
