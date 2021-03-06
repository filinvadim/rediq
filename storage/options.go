package storage

type cacheOpt func(o *cacheOptions)

type cacheOptions struct {
	ItemsNum   uint
	BucketsNum uint
	DumpPath   string
	GCCap      int
}

func ShardsNum(i uint) cacheOpt {
	return func(o *cacheOptions) {
		o.ItemsNum = i
	}
}

func ItemsPerShard(i uint) cacheOpt {
	return func(o *cacheOptions) {
		o.BucketsNum = i
	}
}

func DumpPath(path string) cacheOpt {
	return func(o *cacheOptions) {
		o.DumpPath = path
	}
}

func GCCap(i int) cacheOpt {
	return func(o *cacheOptions) {
		o.GCCap = i
	}
}
