package logger

import (
	"os"
	"testing"

	"github.com/singleton11/k8s-lab/pkg/logger"
)

func TestNewXLog(t *testing.T) {
	log1 := newXLog(&logger.Config{
		Level: logger.LevelDebug,
	})
	if log1 == nil {
		t.Error("Got uninitialized XLog logger")
	}
	log2 := newXLog(&logger.Config{
		Level: logger.LevelInfo,
		Out:   os.Stdout,
		Err:   os.Stdout,
	})
	if log2 == nil {
		t.Error("Got uninitialized XLog logger")
	}
}