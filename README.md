# go-ratelimiter

[![Build Status](https://travis-ci.org/m90/go-ratelimiter.svg?branch=master)](https://travis-ci.org/m90/go-ratelimiter)
[![godoc](https://godoc.org/github.com/m90/go-ratelimiter?status.svg)](http://godoc.org/github.com/m90/go-ratelimiter)

> ratelimit an operation based on id tokens

Package `ratelimiter` enables rate limiting of operations based on id tokens

### Installation using go get

```sh
$ go get github.com/m90/go-ratelimiter
```

### Usage

Create a new Limiter instance using `New(limit time.Duration, cache GetSetter) Throttler`:

```go
limiter := ratelimiter.New(time.Second, cache)
for {
	// this will only run every second
	<-limiter.Throttle("never ending for loop")
	fmt.Println("one second has passed")
}
```

The passed cache needs to implement `GetSetter`:

```go
type GetSetter interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, expiry time.Duration)
}
```

[`go-cache`](https://github.com/patrickmn/go-cache) works out of the box.

Each call to `Throttle(id string) <-chan Result` returns information on the resulting delay or possible errors:

```go
limiter := ratelimiter.New(time.Second, cache)

http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	throttling := <-limiter.Throttle(id)
	if throttling.Error != nil {
		fmt.Println("Error rate limiting")
	}
	if throttling.Delay > 0 {
		fmt.Printf("Delayed next request for %s by %s\n", id, throttling.Delay.String())
	}
})
```

If you want to fulfill `Throttler` without any rate limiting taking place, you can use `NewNoopRateLimiter() Throttler`:

```go
limiter := ratelimiter.NewNoopRateLimiter()
for {
	// this will only run every second
	<-limiter.Throttle("never ending for loop")
	fmt.Println("this will run without any interruption")
}
```


### License
MIT © [Frederik Ring](http://www.frederikring.com)
