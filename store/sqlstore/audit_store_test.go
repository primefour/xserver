package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestAuditStore(t *testing.T) {
	StoreTest(t, storetest.TestAuditStore)
}
