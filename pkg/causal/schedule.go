package causal

import (
	"github.com/matzhouse/causal/pkg/alerters"
	"github.com/matzhouse/causal/pkg/watchers"
	"log"
	"time"
)

type interval int // interval between runs in seconds

type Scheduler struct {
	jobs map[interval][]*Job
}

// NewScheduler returns a Scheduler ready for use
func NewScheduler() *Scheduler {
	return &Scheduler{
		jobs: map[interval][]*Job{},
	}
}

// Register adds a new Job based on it's interval value
func (s *Scheduler) Register(j *Job) (err error) {

	if _, ok := s.jobs[interval(j.Interval)]; !ok {
		s.jobs[interval(j.Interval)] = []*Job{}
	}

	s.jobs[interval(j.Interval)] = append(s.jobs[interval(j.Interval)], j)

	return nil

}

func (s *Scheduler) Run() (err error) {

	for intVal, job := range s.jobs {

		go func(i interval, j *Job) {

			for {

				select {
				case <-time.After(10 * time.Second):
					r, err := j.W.Run()

					if err != nil {
						log.Println(err)
						continue
					}

					// if the run worked, then fire the alerter
					if r.State == watchers.FAIL {

						a := &alerters.Alert{
							Name:    "disk watch /dev/disk3",
							Type:    "not sure what type is for",
							Message: r.Message,
							Time:    time.Now(),
							Result:  r,
						}

						err := j.A.Alert(a)

						if err != nil {
							log.Print(err)
						}

					}

				}

			}

		}(intVal, job[0])

	}

	return nil
}
