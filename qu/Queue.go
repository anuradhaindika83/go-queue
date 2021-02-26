package qu
import (
	"container/list"
	"sync"
)

var qu *list.List
var wlck *sync.Mutex
var rlck *sync.Mutex

type Queue struct {	
}


func (q Queue) Init(){
	qu = list.New()
	wlck = &sync.Mutex{}
	rlck = &sync.Mutex{}
}

func (q Queue) Enqueue(e interface{}){
	wlck.Lock()
	defer wlck.Unlock()
	qu.PushBack(e)
}

func (q Queue) Dequeue() interface{} {
	rlck.Lock()
	defer rlck.Unlock()
	e := qu.Front()
	if e != nil {
		qu.Remove(e)
		return e.Value
	}else{
		return nil
	}
}

