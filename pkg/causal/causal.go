package causal

import (
	"log"
	"os"

	"github.com/matzhouse/causal/pkg/alerters/logger"
	"github.com/matzhouse/causal/pkg/watchers/disk"
	"github.com/spf13/viper"
	"time"
)

type Agent struct {
	config    *viper.Viper
	Watchers  map[string]Watcher
	Alerters  map[string]Alerter
	Scheduler *Scheduler
}

func New() (a *Agent, err error) {

	conf, err := SetupConfig()

	if err != nil {
		log.Println("error in config:", err)
		os.Exit(1)
	}

	//watchers, err := processWatchers(conf)
	//alerters, err := processAlerters(conf)

	// #####################################
	// temp hard code

	// create an alerter
	lg := &logger.Logger{}

	// create a watcher
	wch := disk.New("/dev/disk3", 90)

	// create the Scheduler
	s := NewScheduler()

	j := NewJob(wch, lg, 10)

	s.Register(j)

	// end of temp hard code
	// #####################################

	a = &Agent{
		config:    conf,
		Watchers:  map[string]Watcher{"disk": wch},
		Alerters:  map[string]Alerter{"logger": lg},
		Scheduler: s,
	}

	a.Start()

	return a, err

}

func (a *Agent) Start() (err error) {

	// Run should block until finished.
	err = a.Scheduler.Run()

	time.Sleep(20 * time.Second)

	return nil

}

func processWatchers(conf *viper.Viper) (watchers map[string]Watcher, err error) {

	return
}

func processAlerters(conf *viper.Viper) (alerters map[string]Alerter, err error) {

	return
}

// Job wraps a watcher and an alerter in a job flow with a
// run interval in seconds
type Job struct {
	W        Watcher
	A        Alerter
	Interval int
}

// NewJob takes a watcher, an alerter, and an interval and returns a Job
func NewJob(w Watcher, a Alerter, interval int) (j *Job) {
	return &Job{
		W:        w,
		A:        a,
		Interval: interval,
	}
}
