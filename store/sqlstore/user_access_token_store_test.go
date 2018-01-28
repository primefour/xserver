package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestUserAccessTokenStore(t *testing.T) {
	StoreTest(t, storetest.TestUserAccessTokenStore)
}
