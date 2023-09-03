package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/slack-viewer/pkg/dtos"
)

type MemoryCache struct {
	slackHistoryCache *bigcache.BigCache
}

var memoryCache *MemoryCache

func CurrentCache() *MemoryCache {
	if memoryCache == nil {
		memoryCache, _ = NewBigCache()
	}

	return memoryCache
}

func NewBigCache() (*MemoryCache, error) {
	fmt.Println("Initializing memory cache")
	bCache, err := bigcache.NewBigCache(bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 128,

		// time after which entry can be evicted
		LifeWindow: 168 * time.Hour, //7 days

		// Interval between removing expired entries (clean up).
		// If set to <= 0 then no action is performed.
		// Setting to < 1 second is counterproductive — bigcache has a one second resolution.
		CleanWindow: 24 * time.Hour, //everyday

		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,

		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 1000000 * 10, //100MB

		// prints information about additional memory allocation
		Verbose: true,

		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 4096,

		// callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A bitmask representing the reason will be returned.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		OnRemove: nil,

		// OnRemoveWithReason is a callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A constant representing the reason will be passed through.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		// Ignored if OnRemove is specified.
		OnRemoveWithReason: nil,
	})

	if err != nil {
		return nil, fmt.Errorf("new big cache: %w", err)
	}

	return &MemoryCache{
		slackHistoryCache: bCache,
	}, nil
}

func (bc *MemoryCache) Set(key string, slackHistory *dtos.SlackHistory) error {
	sh, err := json.Marshal(slackHistory)
	if err != nil {
		return fmt.Errorf("Marshal error: %w", err)
	}

	err = bc.slackHistoryCache.Set(key, sh)

	if err != nil {
		return fmt.Errorf("Error caching data in memory: %w", err)
	}

	return err
}

func (bc *MemoryCache) GetStats() {

	fmt.Printf("Cache Collisions: %v \n", bc.slackHistoryCache.Stats().Collisions)
	fmt.Printf("Cache Del Hits: %v \n", bc.slackHistoryCache.Stats().DelHits)
	fmt.Printf("Cache DelMisses: %v \n", bc.slackHistoryCache.Stats().DelMisses)
	fmt.Printf("Cache Hits: %v \n", bc.slackHistoryCache.Stats().Hits)
	fmt.Printf("Cache Misses: %v \n", bc.slackHistoryCache.Stats().Misses)
	fmt.Printf("Cache Len: %v \n", bc.slackHistoryCache.Len())
	fmt.Printf("Cache Capacity: %v \n", bc.slackHistoryCache.Capacity())

}

func (bc *MemoryCache) Get(key string) (result *dtos.SlackHistory, err error) {
	fmt.Println("Searching into memory cache for", key)

	slackHistory := dtos.SlackHistory{
		User:     dtos.SlackUser{},
		DMs:      []dtos.SlackHistoryGroup{},
		Mpims:    []dtos.SlackHistoryGroup{},
		Groups:   []dtos.SlackHistoryGroup{},
		Channels: []dtos.SlackHistoryGroup{},
	}

	sh, err := bc.slackHistoryCache.Get(key)

	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			fmt.Println("Entry not found")
			return &slackHistory, nil
		}

		return result, fmt.Errorf("Get Error: %w", err)
	}

	err = json.Unmarshal(sh, &slackHistory)
	if err != nil {
		return result, fmt.Errorf("Unmarshal error: %w", err)
	}

	fmt.Println("Getting data from cache for ", key)
	return &slackHistory, nil
}

func (bc *MemoryCache) Delete(username string) {
	bc.slackHistoryCache.Delete(username)
}
