package LeetCodeHot100

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*CacheNode
	head, tail *CacheNode
}

type CacheNode struct {
	key, value int
	prev, next *CacheNode
}

func newCacheNode(key, value int) *CacheNode {
	return &CacheNode{
		key:   key,
		value: value,
		prev:  nil,
		next:  nil,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*CacheNode),
		head:     newCacheNode(0, 0),
		tail:     newCacheNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) removeNode(node *CacheNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *CacheNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.cache[key]; ok {
		this.moveToHead(n)
		return n.value
	} else {
		return -1
	}
}

func (this *LRUCache) addToHead(node *CacheNode) {
	this.head.next.prev = node
	node.next = this.head.next
	node.prev = this.head
	this.head.next = node
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.cache[key]; ok {
		n.value = value
		this.moveToHead(n)
	} else {
		if this.size == this.capacity {
			delete(this.cache, this.tail.prev.key)
			this.removeNode(this.tail.prev)

			n = newCacheNode(key, value)
			this.cache[key] = n
			this.addToHead(n)
		} else {
			this.size++
			n = newCacheNode(key, value)
			this.cache[key] = n
			this.addToHead(n)
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
