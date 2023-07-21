package sql

import "time"

func RedSget(key string) string {
	get := Red.Get(Ctx, key)
	return get.Val()
}

func RedSadd(key string, v interface{}, t uint) string {
	get := Red.Set(Ctx, key, v, time.Duration(t)*time.Second)
	return get.Val()
}
