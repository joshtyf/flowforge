package taskpool

import "sync"

type Task struct {
	params  interface{}
	handler func(interface{})
}

type TaskPool struct {
	mutex sync.Mutex
	queue []Task
	stop  chan struct{}
}

func NewTaskPool() *TaskPool {
	return &TaskPool{}
}

func (t *TaskPool) Start() {
	go func() {
		for {
			select {
			case <-t.stop:
				// TODO: add graceful shutdown
				return
			default:
				t.mutex.Lock()
				if len(t.queue) > 0 {
					task := t.queue[0]
					t.queue = t.queue[1:]
					t.mutex.Unlock()
					task.handler(task.params)
				} else {
					t.mutex.Unlock()
				}
			}
		}
	}()
}

func (t *TaskPool) Stop() {
	t.stop <- struct{}{}
}

func (t *TaskPool) AddTask(task Task) {
	defer t.mutex.Unlock()
	t.mutex.Lock()
	t.queue = append(t.queue, task)
}
