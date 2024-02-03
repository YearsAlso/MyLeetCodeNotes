package LCR031

import "container/list"

type LRUCache struct {
	catch     list.List
	indexMap  map[int]*list.Element
	valueMap  []int
	maxLength int
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		maxLength: capacity,
		catch:     list.List{},
		valueMap:  make([]int, 10000),
		indexMap:  make(map[int]*list.Element),
	}

	for i := range lruCache.valueMap {
		lruCache.valueMap[i] = -1
	}

	return lruCache
}

func (this *LRUCache) Get(key int) int {
	if this.catch.Len() == 0 {
		return -1
	}

	if this.valueMap[key] == -1 {
		return -1
	}

	element := this.indexMap[key]
	if element == nil {
		return -1
	}

	this.catch.MoveToFront(element)

	return this.valueMap[key]
}

func (this *LRUCache) Put(key int, value int) {
	if this.catch.Len() == 0 {
		this.catch.PushFront(key)
		this.valueMap[key] = value
		this.indexMap[key] = this.catch.Front()
		return
	}

	if element, ok := this.indexMap[key]; ok {
		this.catch.MoveToFront(element)
		this.valueMap[key] = value
		return
	}

	if this.catch.Len() == this.maxLength {
		// 删除最后一个元素
		if this.catch.Len() == this.maxLength {
			back := this.catch.Back()
			delete(this.indexMap, back.Value.(int))
			this.valueMap[back.Value.(int)] = -1
			this.catch.Remove(back)
		}
	}

	this.catch.PushFront(key)
	this.indexMap[key] = this.catch.Front()
	this.valueMap[key] = value
}

func LCROptimize(funcNames []string, paramsList [][]int) []any {
	inst := Constructor(paramsList[0][0])
	var results []any
	for i := 0; i < len(funcNames); i++ {
		funName := funcNames[i]
		params := paramsList[i]

		switch funName {
		case "LRUCache":
			inst = Constructor(params[0])
			results = append(results, nil)
			break
		case "get":
			result := inst.Get(params[0])
			results = append(results, result)
			break
		case "put":
			inst.Put(params[0], params[1])
			results = append(results, nil)
			break
		}
	}

	return results
}
