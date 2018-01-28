package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestLicenseStore(t *testing.T) {
	StoreTest(t, storetest.TestLicenseStore)
}
