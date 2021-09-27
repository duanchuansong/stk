package retry

import (
	"errors"
	"github.com/duanchuansong/stk/xlog"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	err := Retry(func() error {
		xlog.Debug("retry func")
		return errors.New("please retry")
	}, 3, 2*time.Second)
	if err != nil {
		xlog.Error(err)
	}
}