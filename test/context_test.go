package test

import (
	"context"
	"testing"
	"time"
)

func TestContext_timeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	select {
	case <-ctx.Done():
		t.Logf("context done: %v", ctx.Err())
	}
}

func TestContext_cancel(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

	cancel()
	select {
	case <-ctx.Done():
		t.Logf("context done %v", ctx.Err())
	}
}

func sleep(timeout time.Duration, ch chan struct{}) {
	time.Sleep(timeout)
	ch <- struct{}{}
}

func TestContext_multiTask(t *testing.T) {

}
