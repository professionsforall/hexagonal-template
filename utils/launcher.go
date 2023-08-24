package utils

import (
	"context"
	"sync"
	"time"

	"github.com/professionsforall/hexagonal-template/internal/bootstrap/runnable"
	"github.com/professionsforall/hexagonal-template/pkg/log/logger"
)

func Launch(logger logger.AppLogger, wg *sync.WaitGroup, runnables ...runnable.Runnable) {
	defer logger.Info("graceful shutdown completed")
	for _, item := range runnables {
		wg.Add(1)
		go func(i runnable.Runnable) {
			defer wg.Done()
			if err := i.Start(); err != nil {
				logger.Panic(err)
			}
		}(item)
		defer func(i runnable.Runnable) {
			defer wg.Wait()
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			if err := i.ShutDown(ctx); err != nil {
				logger.Error(err)
			}
		}(item)
	}
	Notify()
}
