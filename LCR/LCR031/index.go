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

	for e := this.catch.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			// 将当前元素移动到最前面
			this.catch.MoveToFront(e)
			this.indexMap[key] = this.catch.Front()
			this.valueMap[key] = value
			return
		}
	}

	if this.catch.Len() == this.maxLength {
		// 删除最后一个元素
		back := this.catch.Back()
		delete(this.indexMap, back.Value.(int))
		this.catch.Remove(back)
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
