package utils

import (
	"github.com/ajaka-the-wizard/dnsgo/internal/domain"
	"github.com/miekg/dns"
	"go.uber.org/zap"
)

func GetLogger(w dns.ResponseWriter) *zap.Logger {
	if logger, ok := w.(*domain.ContextWriters); ok {
		return logger.Logger
	}
	return zap.L()
}
