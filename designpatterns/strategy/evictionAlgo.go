package designpatterns

type evictionAlgo interface {
	evict(c *cache)
}
