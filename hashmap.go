package hashmap

type Node struct {
	Key string
	Val interface{}
}

type HashMap struct {
	Buckets []*Node
	Len int
	Cap int
	// load factor to decide when should we expand buckets size
	LF float32
}

func (hashMap *HashMap) Init(bucketSize int) *HashMap {
	hashMap.Buckets = make([]*Node, bucketSize)
	hashMap.Cap = bucketSize
	hashMap.LF = 0.70
	return hashMap
}

func (hashMap *HashMap) Put(key string, value interface{}) {
	index := Hash(key, hashMap.Cap)

	node := hashMap.Buckets[index]

	if node != nil {
		// is this the needle or look for open address
		if node.Key == key {
			// update it
			node.Val = value
		} else {
			// check for next open address
			for {
				index = (index + 1) % hashMap.Cap
				node = hashMap.Buckets[index]

				if node == nil {
					break
				}
			}
		}
	}

	if node == nil {
		// empty spot, fill it in
		newNode := new(Node)
		newNode.Key = key
		newNode.Val = value
		node = newNode

		hashMap.Len += 1
	}

	hashMap.Buckets[index] = node

	loadFactor := float32(hashMap.Len / hashMap.Cap)
	if loadFactor >= hashMap.LF {
		hashMap.Rehash()
	}
}

func (hashMap *HashMap) Rehash() {
	temp := hashMap.Buckets

	hashMap.Cap = hashMap.Cap * 2

	hashMap.Buckets = make([]*Node, hashMap.Cap)

	for i := 0; i < len(temp); i++ {
		node := temp[i]
		index := Hash(node.Key, hashMap.Cap)
		hashMap.Buckets[index] = node
	}
}

func (hashMap *HashMap) Get(key string) (interface{}, bool) {
	index := Hash(key, hashMap.Cap)
	node := hashMap.Buckets[index]

	if node == nil {
		return "", false
	} else {
		if node.Key == key {
			return node.Val, true
		} else {
			// look in next address
			for {
				index = (index + 1) % hashMap.Cap
				node = hashMap.Buckets[index]

				if node == nil {
					return "", false
				} else  if node.Key == key {
					return node.Val, true
				}
			}
		}
	}
}

func (hashMap *HashMap) Del(key string) bool {
	index := Hash(key, hashMap.Cap)
	node := hashMap.Buckets[index]
	lookup := 1

	if node == nil {
		return false
	} else {
		if node.Key == key {
			hashMap.Buckets[index] = nil
			hashMap.Len -= 1
			return true
		} else {
			// look in next address
			for {
				index = (index + 1) % hashMap.Cap
				node = hashMap.Buckets[index]

				lookup += 1

				if node == nil {
					return false
				} else  if node.Key == key {
					hashMap.Buckets[index] = nil
					hashMap.Len -= 1
					return true
				}
			}
		}
	}
}

func Hash(key string, bucketSize int) int {
	sum := 0
	for i := 0; i < len(key); i++ {
		sum += i * int(key[i])
	}
	return sum % bucketSize
}
