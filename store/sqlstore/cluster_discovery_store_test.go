package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestClusterDiscoveryStore(t *testing.T) {
	StoreTest(t, storetest.TestClusterDiscoveryStore)
}
