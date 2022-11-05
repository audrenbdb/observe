Provides an observable that can be subscribed to.
# Usage

Given observe should write numbers to subscriber:

```go
    ctx := context.Background()

    observable := observe.NewObservable[int](ctx, func(subscriber observe.Subscriber[int]) error {
        subscriber.Next(1)
        subscriber.Next(2)
        subscriber.Next(3)
        
        return nil
    })

    observable.Subscribe(ctx, observe.Subscription[int]{
        Next: func(n int) {
            fmt.Printf("Received number: %d", n)
        },
    }) 
	
    // output:
    // Received number: 1
    // Received number: 2
    // Received number: 3
```

# Options

Listened events with OnEmit can be filtered with a boolean match.

Only matching events will be emitted.

I.E: 
```go
  match := func(n int) bool { return n > 5 }
  emitter.OnEmit(ctx, func(n int) {
    fmt.Printf("Received number: %d\n", n)
  })

  for _, n := range []int{1, 5, 8, 10} {
    emitter.Emit(n)
  }
  // output:
  // Received number: 8
  // Received number: 10
```


