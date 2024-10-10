package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Operation func(ctx context.Context) error

func GracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]Operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		ch:=<-s
		log.Println("shutdown")
		log.Println(ch)

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Println(fmt.Sprintf("timeout %d ms has been elapsed, force exit app", timeout.Milliseconds()))
			os.Exit(0)
		})
		defer timeoutFunc.Stop()

		var wg sync.WaitGroup
		for key, op := range ops {
			wg.Add(1)
			go func(name string, f Operation) {
				defer wg.Done()
				log.Println(fmt.Sprintf("shutting down: %s", name))
				if err := f(ctx); err != nil {
					log.Println(fmt.Sprintf("%s: shutting down failed: %s", name, err.Error()))
					return
				}
				log.Println(fmt.Sprintf("%s was shutted down gracefully", name))
			}(key, op)
		}
		wg.Wait()
		close(wait)
	}()
	return wait
}
