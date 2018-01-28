package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestComplianceStore(t *testing.T) {
	StoreTest(t, storetest.TestComplianceStore)
}
