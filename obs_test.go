package observe_test

import (
	"errors"
	"github.com/audrenbdb/observe"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestObservable(t *testing.T) {
	numbers := make([]int, 0)
	done := false

	observable := observe.NewObservable[int](func(subscriber observe.Subscriber[int]) error {
		subscriber.Next(1)
		subscriber.Next(2)
		subscriber.Next(3)

		subscriber.Complete()
		return nil
	})

	observable.Subscribe(observe.Subscription[int]{
		Next: func(n int) {
			numbers = append(numbers, n)
		},
		Complete: func() {
			done = true
		},
	})

	diff := cmp.Diff(numbers, []int{1, 2, 3})
	if diff != "" {
		t.Errorf("got %#v, want [1, 2, 3]", numbers)
	}
	if done != true {
		t.Errorf("subscription completion should set done to true, got false")
	}
}

func TestObservableError(t *testing.T) {
	errMsg := ""

	observable := observe.NewObservable[int](func(subscriber observe.Subscriber[int]) error {
		return errors.New("fatal crash")
	})

	observable.Subscribe(observe.Subscription[int]{
		Error: func(err error) {
			errMsg = err.Error()
		},
	})

	if errMsg != "fatal crash" {
		t.Errorf("want fatal crash error")
	}
}
