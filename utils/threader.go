package utils

const maxThreads = 10000

type Threader struct {
	threads int
	n       int
}

func MakeThreader(n int) *Threader {
	threader := new(Threader)
	threader.n = n
	if n <= maxThreads {
		threader.threads = n
	} else {
		threader.threads = maxThreads
	}
	return threader
}

func (t *Threader) GetThreadsCount() int {
	return t.threads
}

func (t *Threader) GetInterval(batchNumber int) (int, int) {
	to := t.getThreadLoading() * batchNumber
	from := to - t.getThreadLoading() + 1
	if batchNumber == t.threads {
		to += t.getThreadLoadingRemainder()
	}
	return from, to
}

func (t *Threader) getThreadLoading() int {
	return t.n / t.threads
}

func (t *Threader) getThreadLoadingRemainder() int {
	return t.n % t.threads
}
