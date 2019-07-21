package memo

/*
Extend the Func type and the (*Memo).Get method so that callers may provide an optional done channel through which they
can cancel the operation (§8.9). The results of a cancelled Func call should not be cached.
*/

import "sync"

// Func is the type of the function to memoize.
type Func func(string, chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

func (memo *Memo) Get(key string, cancel chan struct{}) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		func_result := make(chan result)
		go func() {
			value, err := memo.f(key, cancel)
			func_result <- result{value, err}
		}()

		select {
		case <- cancel:
			//if cancel, don't cache result
			memo.mu.Lock()
			delete(memo.cache, key)
			memo.mu.Unlock()
		case e.res = <- func_result:
		}

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
}
