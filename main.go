package main

import (
	"fmt"
	"go-queue/qu"
	"sync"
)

func main(){
	q := qu.Queue{}
	q.Init()
	w := sync.WaitGroup{} 
	w.Add(1)

	go func (){
		for i:=0; i<1000000; i++ {
			q.Enqueue(i)
		}
		fmt.Println("done")
		w.Done()
	}()

	w.Wait()
	fmt.Println("Enqueue done")
	for i:=0; i<1000; i++ {
		w.Add(1)
		go func (){
			for {
				e := q.Dequeue()
				if e == nil{
					break
				}else{
					fmt.Println(e)
				}
			}
			fmt.Println("done")
			w.Done()
		}()
	}

	

	w.Wait()
}

