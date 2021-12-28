package main

import (
	"errors"
	"log"
	"sync"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

func main() {
	syncing()
	broad()
	errgrp()
	sema()
	running()
}
func running() {
	ch := make(chan struct{}, 3)
	doneCh := make(chan struct{})
	go send(ch, doneCh)
	go receive(ch)
	go receive(ch)

	<-doneCh
}
func send(ch, doneCh chan<- struct{}) {
	t := time.NewTimer(3 * time.Second)
	for {
		select {
		case <-t.C:
			close(ch)
			close(doneCh)
			return
		case ch <- struct{}{}:
		}
	}
}
func receive(ch <-chan struct{}) {
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
		}
		log.Println("received")
	}

}

func sema() {
	sem := semaphore.NewWeighted(5)

	go semDo(sem, func() { time.Sleep(1 * time.Second) }, 1)
	go semDo(sem, func() { time.Sleep(1 * time.Second) }, 2)
	go semDo(sem, func() { time.Sleep(1 * time.Second) }, 3)

	time.Sleep(5 * time.Second)
}
func semDo(sem *semaphore.Weighted, f func(), w int64) {
	if err := sem.Acquire(context.Background(), w); err != nil {
		log.Println(err)
		return
	}
	defer sem.Release(w)
	log.Printf("acquired %d", w)
	f()
}
func errgrp() {
	var eg errgroup.Group

	for i := 0; i < 10; i++ {
		n := i
		eg.Go(func() error {
			return even(n)
		})
	}
	if err := eg.Wait(); err != nil {

	}
}
func even(n int) error {
	if n%2 == 0 {
		return errors.New("err")
	}
	time.Sleep(1 * time.Second)
	log.Printf("%d err called", n)
	return nil
}
func syncing() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			doing(n)
		}(i)
	}
	wg.Wait()
}
func doing(n int) {
	time.Sleep(1 * time.Second)
	log.Printf("%d called", n)
}
func broad() {
	doneCh := make(chan struct{})
	for i := 0; i < 10; i++ {
		i := i
		go do(i, doneCh)
	}
	close(doneCh)
	time.Sleep(300 * time.Millisecond)

}
func do(n int, doneCh <-chan struct{}) {
	for {
		select {
		case <-doneCh:
			log.Printf("fineshed %d", n)
			return
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
