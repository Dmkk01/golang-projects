package main

import "fmt"

type Job struct {
	State string
	done  chan struct{}
}

func (j *Job) Wait() {
	<-j.done
}

func (j *Job) Done() {
	j.State = "done"
	close(j.done)
}

func main() {
	ch := make(chan Job)
	go func() {
		j := <-ch
		j.Done()
	}()

	j := Job{State: "running", done: make(chan struct{})}
	ch <- j
	j.Wait()
	fmt.Println(j.State)
}
