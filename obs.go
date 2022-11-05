package observe

// Observable refers to an object that one can subscribe to.
type Observable[T any] struct {
	onSub func(subscriber Subscriber[T]) error
}

// Subscribe to the observable with a Subscription object.
func (o *Observable[T]) Subscribe(sub Subscription[T]) {
	s := &subscription[T]{
		next:     sub.Next,
		error:    sub.Error,
		complete: sub.Complete,
	}
	err := o.onSub(s)
	if err != nil {
		s.error(err)
	}
}

// Subscriber is an interface that exposes methods available to an observable.
type Subscriber[T any] interface {
	Next(T)
	Complete()
}

// Subscription is the parameter passed to an observable Subscribe method.
// Next, Error, and Complete are all optional functions to handle observable flow.
type Subscription[T any] struct {
	Next     func(T)
	Error    func(error)
	Complete func()
}

// subscription is a private method implementing the Subscriber interface.
type subscription[T any] struct {
	next     func(T)
	error    func(error)
	complete func()
}

func (s *subscription[T]) Next(t T) {
	if s.next != nil {
		s.next(t)
	}
}

func (s *subscription[T]) Error(err error) {
	if s.error != nil {
		s.error(err)
	}
}

func (s *subscription[T]) Complete() {
	if s.complete != nil {
		s.complete()
	}
}

// NewObservable creates an observable that one can subscribe to.
// All subscribers may provide a Subscription object to manage subscription flow.
func NewObservable[T any](fn func(subscriber Subscriber[T]) error) *Observable[T] {
	return &Observable[T]{
		onSub: fn,
	}
}
