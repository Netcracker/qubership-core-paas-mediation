package utils

import (
	"github.com/knadh/koanf/providers/confmap"
	"github.com/netcracker/qubership-core-lib-go/v3/configloader"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/resource"
	"testing"
	"time"
)

func Test_CacheSettings_BinarySI_Format(t *testing.T) {
	assertions := require.New(t)
	configloader.Init(&configloader.PropertySource{Provider: configloader.AsPropertyProvider(confmap.Provider(
		map[string]any{"memory.limit": "80Mi"}, "."))})
	memoryLimit := 80 * 1024 * 1024
	numItems, maxSizeInBytes, itemSizeInBytes, ttl := GetCacheSettings()
	numItemsExpected := maxSizeInBytes / itemSizeInBytes
	assertions.Equal(numItemsExpected, numItems)
	assertions.Equal(int64(memoryLimit*4/10), maxSizeInBytes)
	assertions.Equal(int64(8*1024), itemSizeInBytes)
	assertions.Equal(0*time.Second, ttl)
}

func Test_CacheSettings_DecimalSI_Format(t *testing.T) {
	assertions := require.New(t)
	configloader.Init(&configloader.PropertySource{configloader.AsPropertyProvider(confmap.Provider(
		map[string]any{"memory.limit": "80M"}, ".")), nil})
	memoryLimit := 80 * 1000 * 1000
	numItems, maxSizeInBytes, itemSizeInBytes, ttl := GetCacheSettings()
	numItemsExpected := maxSizeInBytes / itemSizeInBytes
	assertions.Equal(numItemsExpected, numItems)
	assertions.Equal(int64(memoryLimit*4/10), maxSizeInBytes)
	assertions.Equal(int64(8*1024), itemSizeInBytes)
	assertions.Equal(0*time.Second, ttl)
}

func Test_CacheSettings_BinarySI_FormatCustom(t *testing.T) {
	assertions := require.New(t)
	expectedMaxTimeSize := 10 * 1024
	quantity := resource.NewQuantity(int64(expectedMaxTimeSize), resource.BinarySI)
	configloader.Init(&configloader.PropertySource{configloader.AsPropertyProvider(confmap.Provider(
		map[string]any{"memory.limit": "80Mi", "cache.max.item.size": quantity.String(), "cache.ttl": "2h"}, ".")), nil})
	memoryLimit := 80 * 1024 * 1024
	numItems, maxSizeInBytes, itemSizeInBytes, ttl := GetCacheSettings()
	numItemsExpected := maxSizeInBytes / itemSizeInBytes
	assertions.Equal(numItemsExpected, numItems)
	assertions.Equal(int64(memoryLimit*4/10), maxSizeInBytes)
	assertions.Equal(int64(expectedMaxTimeSize), itemSizeInBytes)
	assertions.Equal(2*time.Hour, ttl)
}
