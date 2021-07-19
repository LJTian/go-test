package main

import (
    "log"
    "runtime"
    "time"
	"net/http"
	_ "net/http/pprof"
)

func readMemStats(){

	var ms runtime.MemStats

	runtime.ReadMemStats(&ms)

	log.Printf("===> Alloc:%d(M) HeapIdle:%d(M) HeapReleased:%d(M)", ms.Alloc/1024/1024, ms.HeapIdle/1024/1024, ms.HeapReleased/1024/1024 )

}

func test() {
    //slice 会动态扩容，用slice来做堆内存申请
    container := make([]int, 8)

    log.Println(" ===> loop begin.")
    for i := 0; i < 32*1000*1000; i++ {
        container = append(container, i)
    }
    log.Println(" ===> loop end.")
}

func main() {

	go func(){
		log.Println(http.ListenAndServe("0.0.0.0:10000", nil))
	}()

    log.Println("Start.")

	readMemStats()
    test()
	readMemStats()

    log.Println("force gc.")
    runtime.GC() //强制调用gc回收

    log.Println("Done.")
	readMemStats()

	go func(){
		for { 
			readMemStats()
			time.Sleep(10 * time.Second)
		}
	
	}()

    time.Sleep(3600 * time.Second) //睡眠，保持程序不退出
}

