package lrucache

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

type LRUCache[K comparable, V any] struct {
	length        uint
	capacity      uint
	head          *node[V]
	tail          *node[V]
	lookup        map[K]*node[V]
	reverseLookup map[*node[V]]K
}

func (c *LRUCache[K, V]) Upsert(key K, value V) {
	if node, ok := c.lookup[key]; ok {
		node.value = value
		c.detach(node)
		c.prepend(node)
		return
	}

	node := &node[V]{value: value}
	c.lookup[key] = node
	c.reverseLookup[node] = key
	c.length++
	c.prepend(node)

	if c.length > c.capacity {
		c.evict()
	}
}

func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	node, ok := c.lookup[key]
	if !ok {
		return *new(V), false
	}

	c.detach(node)
	c.prepend(node)

	return node.value, true
}

func (c *LRUCache[K, V]) evict() {
	evicted := c.tail
	c.detach(evicted)

	key := c.reverseLookup[evicted]
	delete(c.lookup, key)
	delete(c.reverseLookup, evicted)

	c.length--
}

func (c *LRUCache[K, V]) detach(node *node[V]) {
	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if c.head == node {
		c.head = node.next
	}

	if c.tail == node {
		c.tail = node.prev
	}

	node.next = nil
	node.prev = nil
}

func (c *LRUCache[K, V]) prepend(node *node[V]) {
	if c.head == nil {
		c.head = node
		c.tail = node
		return
	}

	node.next = c.head
	c.head.prev = node
	c.head = node
}
