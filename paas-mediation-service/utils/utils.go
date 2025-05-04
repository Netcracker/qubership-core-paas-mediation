package utils

import (
	"github.com/netcracker/qubership-core-lib-go/v3/configloader"
	"k8s.io/apimachinery/pkg/api/resource"
	"time"
)

func GetWatchClientTimeout() time.Duration {
	timeout := configloader.GetKoanf().String("watch.client.timeout") //WATCH_CLIENT_TIMEOUT
	if timeout == "" {
		return 10 * time.Second
	}
	t, err := time.ParseDuration(timeout)
	if err != nil {
		panic("failed to parse 'watch.client.timeout' property: " + err.Error())
	}
	return t
}

func GetMemoryLimit() *resource.Quantity {
	memoryLimitStr := configloader.GetKoanf().MustString("memory.limit")
	memoryLimit, err := resource.ParseQuantity(memoryLimitStr)
	if err != nil {
		panic("failed to parse 'memory.limit' property: " + err.Error())
	}
	return &memoryLimit
}

func GetCacheSettings() (numItems int64, maxSizeInBytes int64, maxItemSizeInBytes int64, ttl time.Duration) {
	maxItemsAmount := configloader.GetKoanf().Int64("cache.max.items.amount")
	maxCacheSizeStr := configloader.GetKoanf().String("cache.max.size")     //CACHE_MAX_SIZE
	maxItemSizeStr := configloader.GetKoanf().String("cache.max.item.size") //CACHE_MAX_ITEM_SIZE
	cacheTTLStr := configloader.GetKoanf().String("cache.ttl")
	mustParse := func(val string) *resource.Quantity {
		quantity := resource.MustParse(val)
		return &quantity
	}
	if maxItemSizeStr == "" || mustParse(maxItemSizeStr).Value() == int64(0) {
		maxItemSizeStr = "8Ki"
	}
	memoryLimit := GetMemoryLimit()
	if maxCacheSizeStr == "" || mustParse(maxCacheSizeStr).Value() == int64(0) {
		maxCacheSizeStr = resource.NewQuantity(memoryLimit.Value()*4/10, memoryLimit.Format).String()
	}
	maxItemSize, err := resource.ParseQuantity(maxItemSizeStr)
	if err != nil {
		panic("failed to parse 'cache.max.item.size' property: " + err.Error())
	}
	maxCacheSize, err := resource.ParseQuantity(maxCacheSizeStr)
	if err != nil {
		panic("failed to parse 'cache.max.size' property: " + err.Error())
	}
	maxItemSizeBytes := maxItemSize.Value()
	maxCacheSizeBytes := maxCacheSize.Value()
	if maxItemsAmount == 0 {
		maxItemsAmount = maxCacheSizeBytes / maxItemSizeBytes
	}
	if cacheTTLStr == "" {
		cacheTTLStr = "0"
	}
	ttl, err = time.ParseDuration(cacheTTLStr)
	if err != nil {
		panic("failed to parse 'cache.ttl' property: " + err.Error())
	}
	return maxItemsAmount, maxCacheSizeBytes, maxItemSizeBytes, ttl
}
