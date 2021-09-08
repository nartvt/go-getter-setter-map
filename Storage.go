package main

type Storage interface {
	Delete(key string)
	Get(key string) string
	Put(key, value string)
}

type DataMap map[string]string

func (d DataMap) Delete(key string) {
	delete(d, key)
}
func (d DataMap) Get(key string) string {
	return d[key]
}

func (d DataMap) Put(key, value string) {
	d[key] = value
}

