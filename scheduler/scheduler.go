// Select the worker on which the task can be run
// Score candidates
// Pick candidates from good to worse
package scheduler

type Scheduler interface {
	SelectCandidateNodes()
	Score()
	Pick()
}
