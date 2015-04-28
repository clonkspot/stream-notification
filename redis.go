package main

// Caches the given item. Errors are ignored.
func PutCache(key string, duration int, item []byte) {
	r, err := redisPool.Get()
	if err != nil {
		return
	}
	defer redisPool.CarefullyPut(r, &err)

	err = r.Cmd("SETEX", "stream-notification:"+key, duration, item).Err
}

// Retrieves the given item from cache, returning nil if it doesn't exist.
func GetCache(key string) []byte {
	r, err := redisPool.Get()
	if err != nil {
		return nil
	}
	defer redisPool.CarefullyPut(r, &err)

	res, err := r.Cmd("GET", "stream-notification:"+key).Bytes()
	if err != nil {
		return nil
	}
	return res
}
