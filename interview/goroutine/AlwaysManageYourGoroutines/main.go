package main

import (
	"context"

	pb "etcd-3.3.12/gopath/src/gopkg.in/cheggaaa/pb.v1"
)

func Invoke(ctx context.Context, job *api.Job) (*pb.Empty, error) {
	go func() {
		if err := ma.invoke(job); err != nil {
			logger.Errorf(ctx, "failed to invoke")
		}
	}()
	return &pb.Empty{}, nil
}

func Invoke(ctx context.Context, job *api.Job) (*pb.Empty, error{
	ma.workerChannel <- &worker.Work{
	   Task:           job,
	   WorkerFunction: ma.invoke,
	}
	return &protobuf.Empty{}, nil
 }

 func (worker *Worker) Start(){
	go func(){
	   for {
		  select{
		  case job :=<- worker.tasks:{
			 if job != nil{
				err := job.WorkerFunction(job.Task)
				if err != nil {
				   logger.Errorf("worker function failed")
				}
			 }
		  }
		  case <- worker.stop:
			 break
		  }
	   }
	}()
 }

func main() {

}
