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
	// 使用指针对象，结果就是 done
	//ch := make(chan Job)
	ch := make(chan *Job)
	go func() {
		j := <-ch
		j.Done()
	}()

	job := Job{"ready", make(chan struct{})}
	ch <- &job
	job.Wait()
	fmt.Println(job.State)
}
