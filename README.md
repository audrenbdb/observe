Provides an observable that can be subscribed to.
# Usage

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