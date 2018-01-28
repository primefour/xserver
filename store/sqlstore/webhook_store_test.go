package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestWebhookStore(t *testing.T) {
	StoreTest(t, storetest.TestWebhookStore)
}
