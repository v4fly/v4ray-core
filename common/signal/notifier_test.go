package signal_test

import (
	"testing"

	. "github.com/v4fly/v4ray-core/v0/common/signal"
)

func TestNotifierSignal(t *testing.T) {
	n := NewNotifier()

	w := n.Wait()
	n.Signal()

	select {
	case <-w:
	default:
		t.Fail()
	}
}
