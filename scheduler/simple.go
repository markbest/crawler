package scheduler

import "github.com/markbest/crawler/engine"

type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workChan <- r
	}()
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workChan
}

func (s *SimpleScheduler) WorkerReady(w chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workChan = make(chan engine.Request)
}
