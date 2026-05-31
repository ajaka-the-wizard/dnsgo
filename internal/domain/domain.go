package domain

import (
	"github.com/miekg/dns"
	"go.uber.org/zap"
)

type ContextWriters struct {
	dns.ResponseWriter
	Logger *zap.Logger
}
