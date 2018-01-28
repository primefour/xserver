package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestPluginStore(t *testing.T) {
	StoreTest(t, storetest.TestPluginStore)
}
