package util

import "time"

type callback func() error

func DoWithInterval(duration time.Duration, f callback) chan error {
	err := make(chan error)
	go func() {
		t := time.NewTicker(duration)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				e := f()
				if e != nil {
					err <- e
					return
				}
			}
		}
	}()
	return err
}
