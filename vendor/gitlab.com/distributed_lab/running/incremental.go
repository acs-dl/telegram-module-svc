package running

import (
	"context"
	"time"

	"fmt"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Logger interface {
	Log(level uint32, fields map[string]interface{}, err error, withStack bool, args ...interface{})
}

// WithBackOff calls the runner with the normalPeriod, while it returns nil error.
// Once the runner returned error, it will be called with the minAbnormalPeriod,
// multiplying the period on 2 each retry, but not bigger than maxRetryPeriod.
// Once the runner returns nil(no error) in abnormal execution,
// it's execution comes back to the normal one and the runner
// is called with the normalPeriod again.
// If after that the runner returned error again -
// execution will be again switched to abnormal.
//
// Runner function must do some job only once(not in a loop), iteration of job execution in loop is
// responsibility of WithBackOff func.
//
// You are generally not supposed to log error inside the runner,
// you should return error instead - errors returned from runner function
// will be logged with stack and the runnerName.
//
// If runner panics, the panic value will be converted to error and logged with stack,
// the retry logic in this case works as if the runner returned error.
//
// WithBackOff is a blocking function, it returns only if ctx is canceled.
func WithBackOff(
	ctx context.Context,
	log Logger,
	runnerName string,
	runner func(context.Context) error,
	normalPeriod time.Duration,
	minAbnormalPeriod,
	maxAbnormalPeriod time.Duration) {

	if normalPeriod == 0 {
		normalPeriod = 1
	}

	fields := logan.F{
		"runner": runnerName,
	}
	normalTicker := time.NewTicker(normalPeriod)

	for ; ; waitForCtxOrTicker(ctx, normalTicker) {
		// To guarantee *no* work is being done after ctx was cancelled.
		if IsCancelled(ctx) {
			log.Log(uint32(logan.InfoLevel), fields, nil, false, fmt.Sprintf("Context is canceled - stopping '%s' runner.", runnerName))
			return
		}

		err := RunSafely(ctx, runnerName, runner)

		if err != nil {
			level := uint32(logan.ErrorLevel)
			if errors.Cause(err) == context.DeadlineExceeded || errors.Cause(err) == context.Canceled {
				level = uint32(logan.InfoLevel)
			}
			log.Log(level, fields, err, true, fmt.Sprintf("Runner '%s' returned error.", runnerName))

			runAbnormalExecution(ctx, log, runnerName, runner, minAbnormalPeriod, maxAbnormalPeriod)
			if IsCancelled(ctx) {
				log.Log(uint32(logan.InfoLevel), fields, nil, false, fmt.Sprintf("Context is canceled - stopping '%s' runner.", runnerName))
				return
			}
		}
	}
}

// UntilSuccess calls the runner again and again while the runner returns false success or a non-nil error.
// The time pause before the retry the runner is at first equal to minRetryPeriod
// and becomes twice bigger each retry, but not bigger than maxRetryPeriod.
//
// You are generally not supposed to log error inside the runner,
// you should return error instead - non-nil errors returned from runner function
// will be logged with stack and the runnerName.
//
// If runner panics, the panic value will be converted to error and logged with stack,
// the retry logic in this case works as if the runner returned false with error.
//
// UntilSuccess returns only if the runner returns success without error, or if ctx was canceled.
func UntilSuccess(
	ctx context.Context,
	log Logger,
	runnerName string,
	runner func(context.Context) (bool, error),
	minRetryPeriod,
	maxRetryPeriod time.Duration) {

	success, err := runSafelyWithSuccess(ctx, runnerName, runner)
	if success && err == nil {
		// Brief success!
		return
	}

	incrementalTimer := newIncrementalTimer(minRetryPeriod, maxRetryPeriod, 2)

	for !success || err != nil {
		fields := logan.F{
			"runner":           runnerName,
			"retry_number":     incrementalTimer.iteration,
			"next_retry_after": incrementalTimer.currentPeriod,
		}

		if err != nil {
			log.Log(uint32(logan.ErrorLevel), fields, err, true, fmt.Sprintf("Runner '%s' returned error.", runnerName))
		} else {
			// Just not success, but without any errors.
			log.Log(uint32(logan.InfoLevel), fields, nil, false, fmt.Sprintf("Runner '%s' didn't meet success.", runnerName))
		}

		select {
		case <-ctx.Done():
			return
		case <-incrementalTimer.next():
			// To guarantee *no* work is being done after ctx was cancelled.
			if IsCancelled(ctx) {
				return
			}

			success, err = runSafelyWithSuccess(ctx, runnerName, runner)
		}
	}
}

// WithThreshold calls the runner again and again while the runner returns false success or a non-nil error
// but takes not more attempts then allowed by `allowedAttempts` param.
// The time pause before the retry the runner is at first equal to minRetryPeriod
// and becomes twice bigger each retry, but not bigger than maxRetryPeriod.
//
// You are generally not supposed to log error inside the runner,
// you should return error instead - non-nil errors returned from runner function
// will be logged with stack and the runnerName.
//
// If runner panics, the panic value will be converted to error and logged with stack,
// the retry logic in this case works as if the runner returned false with error.
//
// WithThreshold logs the result of its work before finishing the execution. The short conclusion
// about the results is logged (runner succeeded, runner returned with error, runner did not succeeded,
// context was canceled). You are allowed to log states during runner execution in purposes of info/debugging.
//
// WithThreshold returns only if the runner returns success without error, or if ctx was canceled, or number of taken
// attempts reached the threshold.
func WithThreshold(
	ctx context.Context,
	log Logger,
	runnerName string,
	runner func(ctx context.Context) (bool, error),
	minRetryPeriod,
	maxRetryPeriod time.Duration,
	allowedAttempts uint64) {

	if allowedAttempts <= 0 {
		allowedAttempts = 1
	}

	success, err := runSafelyWithSuccess(ctx, runnerName, runner)
	if success && err == nil {
		// Brief success!
		return
	}

	incrementalTimer := newIncrementalTimer(minRetryPeriod, maxRetryPeriod, 2)

	canceled := false
	for attempt := uint64(1); attempt < allowedAttempts && (!success || err != nil); attempt++ {
		fields := logan.F{
			"runner":           runnerName,
			"retry_number":     incrementalTimer.iteration,
			"next_retry_after": incrementalTimer.currentPeriod,
		}

		if err != nil {
			log.Log(uint32(logan.ErrorLevel), fields, err, true, fmt.Sprintf("Runner '%s' returned error.", runnerName))
		} else {
			// Just not success, but without any errors.
			log.Log(uint32(logan.InfoLevel), fields, nil, false, fmt.Sprintf("Runner '%s' didn't meet success.", runnerName))
		}

		select {
		case <-ctx.Done():
			canceled = true
			break
		case <-incrementalTimer.next():
			// To guarantee *no* work is being done after ctx was cancelled.
			if IsCancelled(ctx) {
				canceled = true
				break
			}

			success, err = runSafelyWithSuccess(ctx, runnerName, runner)
		}
	}

	fields := logan.F{
		"total_attempts": allowedAttempts,
	}

	switch {
	case canceled:
		log.Log(uint32(logan.ErrorLevel), fields, nil, false, fmt.Sprintf("Runner '%s' stopped from outside with context cancellation", runnerName))
	case err != nil:
		log.Log(uint32(logan.ErrorLevel), fields, err, true, fmt.Sprintf("Runner '%s' returned error.", runnerName))
	case !success:
		log.Log(uint32(logan.InfoLevel), fields, nil, false, fmt.Sprintf("Runner '%s' didn't meet success.", runnerName))
	default:
		log.Log(uint32(logan.InfoLevel), fields, nil, false, fmt.Sprintf("Runner '%s' succeeded", runnerName))
	}
}

func waitForCtxOrTicker(ctx context.Context, ticker *time.Ticker) {
	select {
	case <-ctx.Done():
		return
	case <-ticker.C:
		return
	}
}

// RunAbnormalExecution Only returns if runner returned nil error or ctx was canceled.
func runAbnormalExecution(
	ctx context.Context,
	log Logger,
	runnerName string,
	runner func(context.Context) error,
	minRetryPeriod,
	maxRetryPeriod time.Duration) {

	incrementalTimer := newIncrementalTimer(minRetryPeriod, maxRetryPeriod, 2)

	for {
		select {
		case <-ctx.Done():
			return
		case <-incrementalTimer.next():
			// To guarantee *no* work is being done after ctx was cancelled.
			if IsCancelled(ctx) {
				return
			}

			err := RunSafely(ctx, runnerName, runner)
			if err == nil {
				log.Log(uint32(logan.InfoLevel), logan.F{
					"runner": runnerName,
				}, nil, false, fmt.Sprintf("Runner '%s' is returning to normal execution.", runnerName))
				return
			}

			log.Log(uint32(logan.ErrorLevel), logan.F{
				"runner":           runnerName,
				"retry_number":     incrementalTimer.iteration,
				"next_retry_after": incrementalTimer.currentPeriod,
			}, err, true, fmt.Sprintf("Runner '%s' returned error.", runnerName))
		}
	}
}

// RunSafely handles panic using defer.
//
// If no panic happens - RunSafely does nothing except
// calling the provided runner with the provided context.
func RunSafely(ctx context.Context, runnerName string, runner func(context.Context) error) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = errors.Wrap(errors.WithStack(errors.FromPanic(rec)), fmt.Sprintf("Runner '%s' panicked", runnerName))
		}
	}()

	return runner(ctx)
}

// RunSafely handles panic using defer.
func runSafelyWithSuccess(ctx context.Context, runnerName string, runner func(context.Context) (bool, error)) (success bool, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			success = false
			err = errors.Wrap(errors.WithStack(errors.FromPanic(rec)), fmt.Sprintf("Runner '%s' panicked", runnerName))
		}
	}()

	return runner(ctx)
}

//RunEachMonth - runs `runner` right after the call and then at the begging of the next month at utc.
// In case of error or panic uses same logic as `UntilSuccess`.
func RunEachMonth(
	ctx context.Context,
	log Logger,
	runnerName string,
	runner func(context.Context) error,
	minRetryPeriod,
	maxRetryPeriod time.Duration) {

	UntilSuccess(ctx, log, runnerName, func(ctx context.Context) (bool, error) {
		for {
			if ctx.Err() != nil {
				return false, ctx.Err()
			}
			err := runner(ctx)
			if err != nil {
				return false, err
			}
			waitUntilNextMonthFirstDayMidnightUTC(ctx)
		}
	}, minRetryPeriod, maxRetryPeriod)
}


func waitUntilNextMonthFirstDayMidnightUTC(ctx context.Context) {
	now := time.Now().UTC()
	// time.Date automatically normalizes months
	firstDayOfNextMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 10, 0, time.UTC)
	timer := time.NewTimer(time.Until(firstDayOfNextMonth))
	select {
	case <-timer.C:
	case <-ctx.Done():
	}
	timer.Stop()
}
