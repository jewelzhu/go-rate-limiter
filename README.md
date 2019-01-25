I cannot find a proper Golang rate limiter lib so I wrote one which is referenced from a lua implementation https://github.com/openresty/lua-resty-limit-traffic/blob/master/lib/resty/limit/req.lua
I want requests to be accepted at once if not exceeds limit, delay time returned if it's in the burst, rejected at once if exceeds burst

    l := NewRateLimiter(10, 5)
	c := make(chan int, 1)
	go func() {
		for {
			i := <- c
            // ok=true and delay=0 : the request can be accepted at once
            // ok=true and delay=0 : the request can be executed after the delay
            // ok=false : reject the request at once
			delay, ok := l.incoming()
			if !ok {
				fmt.Println("-")
			} else if delay > 0 {
				log.Printf("accept %d but need sleep %f seconds %s\n", i, delay, time.Now().Format("2016-01-02 15:04:05.000"))
			} else {
				log.Printf("accept %d at once %s", i, time.Now().Format("2016-01-02 15:04:05.000"))
			}
		}
	}()

	for i:=0;i<20;i++{
		time.Sleep(time.Millisecond * 100)
		c <- i
	}

	for i:=0;i<20;i++{
		time.Sleep(time.Millisecond * 50)
		c <- i
	}

	for i:=0;i<20;i++{
		time.Sleep(time.Millisecond * 20)
		c <- i
	}