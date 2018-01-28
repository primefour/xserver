package sqlstore

import (
	"testing"

	"github.com/primefour/xserver/store/storetest"
)

func TestCommandWebhookStore(t *testing.T) {
	StoreTest(t, storetest.TestCommandWebhookStore)
}
