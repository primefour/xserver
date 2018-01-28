package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestOAuthStore(t *testing.T) {
	StoreTest(t, storetest.TestOAuthStore)
}
