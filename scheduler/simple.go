package scheduler

import "remfath.com/crawler/engine"

type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workChan <- r
	}()
}
