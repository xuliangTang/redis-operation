package redis

type MyIterator struct {
	data []interface{}
	index int
}

func NewMyIterator(data []interface{}) *MyIterator {
	return &MyIterator{data: data}
}

func(this *MyIterator) HasNext() bool {
	if len(this.data) == 0 || this.data == nil{
		return false
	}
	return this.index < len(this.data)
}

func(this *MyIterator) Next() (ret interface{}) {
	ret = this.data[this.index]
	this.index += 1
	return
}