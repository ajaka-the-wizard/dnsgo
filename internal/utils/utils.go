package utils

import (
	"fmt"

	"github.com/ajaka-the-wizard/dnsgo/internal/domain"
	"github.com/miekg/dns"
	"go.uber.org/zap"
)

func GenerateID() uint64 {
	// Placeholder, will generate it later
	return 1000
}

func GetLogger(w dns.ResponseWriter) *zap.Logger {
	if logger, ok := w.(*domain.ContextWriters); ok {
		return logger.Logger
	}
	fmt.Println("Yeah, nothing")
	return zap.L()
}
