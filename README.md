# LRU-Cache
A custom cache built using go

* Maintains a specific length of items in cache.
* Most recently used items is stored.
* If the capacity full, items goes out.

### For LRU cache:
* If already present, remove item and add it to the beginning.
* An order of item is maintained
* FIFO

### Use hard coded array

```go 
items := [8]string{"parrot", "tree", "lion", "forest", "rain", "tomato", "potato", "mushroom"}
cache := NewCache()

for _, word := range items {
	cache.Check(word) // Check if the value already exists in the queue or not
	cache.Display()
}
```
