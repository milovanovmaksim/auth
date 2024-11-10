package closer

import (
	"log"
	"os"
	"os/signal"
	"sync"
)

var globalCloser = NewCloser()

// Add добавляет функцию в Closer.
func Add(f ...func() error) {
	globalCloser.Add(f...)
}

// Wait ждет выполнение всех функций, добавленных в Closer.
func Wait() {
	globalCloser.Wait()
}

// CloseAll вызывает все функции, добавленные в Closer.
func CloseAll() {
	globalCloser.CloseAll()
}

// Closer используется для освобождения системных ресурсов (БД, сокеты, и.т.д.) при завершении работы приложения.
type Closer struct {
	mu    sync.Mutex
	once  sync.Once
	done  chan struct{}
	funcs []func() error
}

// NewCloser создает новый объект Closer.
func NewCloser(sig ...os.Signal) *Closer {
	c := &Closer{done: make(chan struct{})}
	if len(sig) > 0 {
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, sig...)
			<-ch
			signal.Stop(ch)
			c.CloseAll()
		}()
	}

	return c
}

// Wait ждет выполнение всех функций, добавленных в Closer.
func (c *Closer) Wait() {
	<-c.done
}

// CloseAll вызывает все функции, добавленные в Closer.
func (c *Closer) CloseAll() {
	c.once.Do(func() {
		defer close(c.done)

		c.mu.Lock()
		funcs := c.funcs
		c.funcs = nil
		c.mu.Unlock()

		errs := make(chan error, len(funcs))
		for _, f := range funcs {
			go func(f func() error) {
				errs <- f()
			}(f)
		}

		for i := 0; i < cap(errs); i++ {
			if err := <-errs; err != nil {
				log.Println("error returned from Closer")
			}
		}
	})
}

// Add добавляет функцию в Closer.
func (c *Closer) Add(f ...func() error) {
	c.mu.Lock()
	c.funcs = append(c.funcs, f...)
	c.mu.Unlock()
}
