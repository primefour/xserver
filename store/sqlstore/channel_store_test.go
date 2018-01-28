package sqlstore

import (
	"testing"

	"github.com/mattermost/mattermost-server/store/storetest"
)

func TestChannelStore(t *testing.T) {
	StoreTest(t, storetest.TestChannelStore)
}
