package qqlt

type Params map[string]interface{}

// NewParams 获取一个post请求参数map
func NewParams() Params {
	return make(Params)
}

// Add params add a item
func (p Params) Add(key string, val interface{}) {
	p[key] = val
}

// Set params set up item, it will be create if it is not exists
func (p Params) Set(key string, val interface{}) {
	if p[key] == nil {
		p.Add(key, val)
	}

	p[key] = val
}

// Del params delete a item
func (p Params) Del(key string) {
	if p[key] == nil {
		return
	}
	delete(p, key)
}

// Get get a item of params
func (p Params) Get(key string, def ...interface{}) interface{} {
	res := p[key]

	if len(def) == 0 {
		return res
	}

	if res == nil {
		p.Add(key, def[0])
		return def[0]
	}
	return res
}

// Clear create a new params, the efficiency is better than iterate to delete every one
func (p Params) Clear() {
	p = NewParams()
}
