package endpoint

import (
    "context"
    "fmt"
    "net"
    "net/http"
    "strconv"
    "sync"
    
    "github.com/gin-gonic/gin"
    "github.com/go4s/configuration"
    "github.com/go4s/handler"
    "golang.org/x/sync/errgroup"
)

func StartServeHTTP(ctx context.Context, workers *errgroup.Group) {
    var (
        err      error
        router   = gin.Default()
        listener net.Listener
    )
    router.Use(middlewares...)
    handler.Hook(router.Group(prefix))
    
    if listener, err = net.Listen("tcp", fmt.Sprintf(":%d", port)); err == nil {
        workers.Go(startListenAndServe(listener, router))
        workers.Go(closeListenerOnCancel(ctx, listener))
        return
    }
    workers.Go(func() error { return err })
    return
}

func startListenAndServe(listener net.Listener, router *gin.Engine) func() error {
    return func() error { return http.Serve(listener, router) }
}

func closeListenerOnCancel(ctx context.Context, listener net.Listener) func() error {
    return func() error {
        for {
            select {
            case <-ctx.Done():
                return listener.Close()
            }
        }
    }
}

var (
    initializer = sync.Once{}
    middlewares []gin.HandlerFunc
    prefix      = ""
    port        = 80
    noRouters   []gin.HandlerFunc
)

func CustomizeNoRoute(handlers ...gin.HandlerFunc) { noRouters = append(noRouters, handlers...) }
func Intercepts(interceptors ...gin.HandlerFunc)   { middlewares = append(middlewares, interceptors...) }

func init() { initializer.Do(initialize) }
func initialize() {
    middlewares = []gin.HandlerFunc{}
    env := configuration.FromEnv()
    prefix = stringOrDefault(env, ServiceHTTPRouterPrefix, prefix)
    port = intOrDefault(env, ServiceHTTPPort, port)
}

func intOrDefault(env configuration.Configuration, s string, p int) int {
    e, found := env[s]
    if !found {
        return p
    }
    if v, err := strconv.Atoi(e.(string)); err == nil {
        return v
    }
    return p
}
func stringOrDefault(env configuration.Configuration, s string, p string) string {
    if e, found := env[s]; found {
        return e.(string)
    }
    return p
}

const (
    ServiceHTTPRouterPrefix = "service.http.prefix"
    ServiceHTTPPort         = "service.http.port"
)
