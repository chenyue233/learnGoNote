package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
}
type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	ReadNotifier
	Run()
}
type ReadNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParserResult)
	e.Scheduler.Run()

	for i := 0;i < e.WorkerCount;i++{
		createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)
	}

	for _,r := range seeds{
		e.Scheduler.Submit(r)
	}

	for {
		result := <- out
		for _,item := range result.Items{
			go func() { e.ItemChan <- item }()
		}
		for _,request := range result.Requests{
			if isDuplicate(request.Url){
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request,out chan ParserResult,ready ReadNotifier)  {
	go func() {
		for {
			// tell scheduler i`m ready
			ready.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url]{
		return true
	}
	visitedUrls[url] =true
	return false
}
