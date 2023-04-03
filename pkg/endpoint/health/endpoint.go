package health

import (
    "github.com/gin-gonic/gin"
    "github.com/go4s/handler"
)

type impl struct{}

func (i impl) Singular() string       { return singular }
func (i impl) Plural() string         { return plural }
func (i impl) Version() string        { return Version }
func (i impl) Group() handler.Grouper { return i }
func (i impl) List(ctx *gin.Context)  { ctx.Status(204) }

func init() { handler.Add(impl{}) }

const (
    singular = "health"
    plural   = "healths"
    Version  = "v1"
)
