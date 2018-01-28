package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestTeamStore(t *testing.T) {
	StoreTest(t, storetest.TestTeamStore)
}
