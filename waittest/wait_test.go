package waittest

import (
	"k8s.io/apimachinery/pkg/util/wait"
	"testing"
	"time"
)

func TestWaitRetry(t *testing.T) {
	err := wait.PollImmediate(200*time.Millisecond, time.Second, func() (done bool, err error) {
		return false, nil
	})
	t.Logf("Error: %s", err.Error())

	for i := 0; i < 3; i++ {
		t.Logf("%d", i)
	}
}
