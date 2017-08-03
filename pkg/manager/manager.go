package manager

type Job func()

type Worker struct {
	done chan struct{}
}

func (w Worker) Do(job Job) {
	job()
	w.done <- struct{}{}
}

type Manager struct {
	q          chan Job
	maxWorkers int
	workers    int
	done       chan struct{}
}

func New(maxWorkers int) *Manager {

	return &Manager{
		maxWorkers: maxWorkers,
		q:          make(chan Job, 1000),
		done:       make(chan struct{}, maxWorkers),
	}

}

func (m Manager) Wake() {
	go m.run()
}

func (m *Manager) run() {
	for job := range m.q {
		worker := m.worker()
		m.workers++
		go worker.Do(job)
	}

}

func (m *Manager) worker() Worker {

	if m.workers < m.maxWorkers {
		return Worker{done: m.done}
	} else {
		<-m.done
		m.workers--
		return Worker{done: m.done}
	}

}

func (m Manager) Add(job Job) {
	m.q <- job
}
