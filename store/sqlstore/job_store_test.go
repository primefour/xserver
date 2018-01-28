package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestJobStore(t *testing.T) {
	StoreTest(t, storetest.TestJobStore)
}
