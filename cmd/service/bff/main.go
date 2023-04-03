package main

import (
    "context"
    "log"
    "os"
    
    "golang.org/x/sync/errgroup"
    _ "station/cmd/service/bff/internal/endpoints"
    "station/pkg/endpoint"
    "station/pkg/workers/signals"
)

func main() {
    workers, root := errgroup.WithContext(context.Background())
    signals.StartListenSignals(root, workers, os.Interrupt)
    endpoint.StartServeHTTP(root, workers)
    log.Printf("exit with err : %v", workers.Wait())
}
