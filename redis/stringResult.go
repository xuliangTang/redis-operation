package redis

type StringResult struct {
	Result string
	Err error
}

func NewStringResult(result string, err error) *StringResult {
	return &StringResult{Result: result, Err: err}
}

func(this *StringResult) Unwrap() string {
	if this.Err != nil {
		panic(this.Err)
	}
	return this.Result
}

func(this *StringResult) UnwrapOr(val string) string {
	if this.Err != nil {
		return val
	}
	return this.Result
}

func(this *StringResult) UnwrapOrElse(f func() string) string {
	if this.Err != nil {
		return f()
	}
	return this.Result
}
