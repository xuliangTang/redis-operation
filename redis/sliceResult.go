package redis

type SliceResult struct {
	Result []interface{}
	Err error
}

func NewSliceResult(result []interface{}, err error) *SliceResult {
	return &SliceResult{Result: result, Err: err}
}

func(this *SliceResult) Unwrap() []interface{} {
	if this.Err != nil {
		panic(this.Err)
	}
	return this.Result
}

func(this *SliceResult) UnwrapOr(values []interface{}) []interface{} {
	if this.Err != nil {
		return values
	}
	return this.Result
}

func(this *SliceResult) Iterator() *MyIterator {
	return NewMyIterator(this.Result)
}