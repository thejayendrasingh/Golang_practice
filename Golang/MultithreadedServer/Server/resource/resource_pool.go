package resource

/*
	Resource package to simulate 2 types of resources, A or B
	We can have any number of resource to match any system.
	We can also introduce extra delay to simuate resource usage additiona to execution time
*/

import (
	"log"
	"sync"
)

var resourceACount int
var resourceAPool = &sync.Pool{
	New: func() interface{} {
		log.Println("resource 1 created")
		resourceACount++
		mem := make([]byte, 1024)
		return &mem
	},
}

var resourceBCount int
var resourceBPool = &sync.Pool{
	New: func() interface{} {
		log.Println("resource 2 created")
		resourceBCount++
		mem := make([]byte, 1024)
		return &mem
	},
}

func GetTotalResourceACount() int {
	return resourceACount
}

func GetTotalResourceBCount() int {
	return resourceBCount
}

func FetchFromResourceA() any {
	return resourceAPool.New()
}

func ReturnToResourceA(mem *any) {
	log.Println("resource 1 returned back")
	resourceAPool.Put(mem)
}

func FetchFromResourceB() any {
	return resourceBPool.New()
}

func ReturnToResourceB(mem *any) {
	log.Println("resource 2 returned back")
	resourceBPool.Put(mem)
}
