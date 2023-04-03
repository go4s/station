package signals

import (
    "context"
    "log"
    "os"
    "os/signal"
    "sync"
    
    "golang.org/x/sync/errgroup"
)

type CancelCallback func() error

var (
    initializer = sync.Once{}
    registry    []CancelCallback
)

func init() { initializer.Do(initialize) }

func initialize() { registry = []CancelCallback{} }

func AddCallbacks(fns ...CancelCallback) { registry = append(registry, fns...) }

func StartListenSignals(ctx context.Context, workers *errgroup.Group, signals ...os.Signal) {
    workers.Go(listenSignals(ctx, signals...))
}

func listenSignals(ctx context.Context, signals ...os.Signal) func() error {
    return func() error {
        signalCtx, cancel := signal.NotifyContext(ctx, signals...)
        defer cancel()
        var fn CancelCallback
        for {
            select {
            case <-signalCtx.Done():
                for i := range registry {
                    if fn = registry[i]; fn == nil {
                        continue
                    }
                    if err := fn(); err != nil {
                        log.Println("exit callback err : %v", err)
                    }
                }
                return ctx.Err()
            }
        }
    }
}
