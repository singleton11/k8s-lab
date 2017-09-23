package handlers

import (
	"net/http"
	"testing"

	"github.com/singleton11/k8s-lab/pkg/config"
	"github.com/singleton11/k8s-lab/pkg/logger"
	"github.com/singleton11/k8s-lab/pkg/logger/standard"
	"github.com/singleton11/k8s-lab/pkg/router/bitroute"
)

func TestReady(t *testing.T) {
	h := New(standard.New(&logger.Config{}), new(config.Config))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Base(h.Ready)(bitroute.NewControl(w, r))
	})

	testHandler(t, handler, http.StatusOK, http.StatusText(http.StatusOK))
}
