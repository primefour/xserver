package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestChannelMemberHistoryStore(t *testing.T) {
	StoreTest(t, storetest.TestChannelMemberHistoryStore)
}
