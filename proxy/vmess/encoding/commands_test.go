package encoding_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"

	"github.com/v4fly/v4ray-core/v0/common"
	"github.com/v4fly/v4ray-core/v0/common/buf"
	"github.com/v4fly/v4ray-core/v0/common/protocol"
	"github.com/v4fly/v4ray-core/v0/common/uuid"
	. "github.com/v4fly/v4ray-core/v0/proxy/vmess/encoding"
)

func TestSwitchAccount(t *testing.T) {
	sa := &protocol.CommandSwitchAccount{
		Port:     1234,
		ID:       uuid.New(),
		AlterIds: 1024,
		Level:    128,
		ValidMin: 16,
	}

	buffer := buf.New()
	common.Must(MarshalCommand(sa, buffer))

	cmd, err := UnmarshalCommand(1, buffer.BytesFrom(2))
	common.Must(err)

	sa2, ok := cmd.(*protocol.CommandSwitchAccount)
	if !ok {
		t.Fatal("failed to convert command to CommandSwitchAccount")
	}
	if r := cmp.Diff(sa2, sa); r != "" {
		t.Error(r)
	}
}

func TestSwitchAccountBugOffByOne(t *testing.T) {
	sa := &protocol.CommandSwitchAccount{
		Port:     1234,
		ID:       uuid.New(),
		AlterIds: 1024,
		Level:    128,
		ValidMin: 16,
	}

	buffer := buf.New()
	csaf := CommandSwitchAccountFactory{}
	common.Must(csaf.Marshal(sa, buffer))

	Payload := buffer.Bytes()

	cmd, err := csaf.Unmarshal(Payload[:len(Payload)-1])
	assert.Error(t, err)
	assert.Nil(t, cmd)
}
