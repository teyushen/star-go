# star-go


## Prerequisites

* Create a github token ([how to generate a github token]())

	> You will get something like `3061ba66c81c7590e3b2a3bd3055fece429fb531`

## Determine how to install star-go

* Golang
* Docker

## Install with Golang

2. Install star-go

	```
	$ go get github.com/teyushen/star-go 
	$ go install github.com/teyushen/star-go 
	```


3. Initial the star-go

	```
	$ star-go init <token>
	```
	 
	> e.g.
	> `$ star-go init 3061ba66c81c7590e3b2a3bd3055fece429fb531`
	

4. Add you interesting github repositories

	```
	$ star-go focus teyushen/star-go teyushen/dockerfile golang/go
	```

5. Order the numbers of size of repositories you are interested

	```
	$ star-go compare
	```
	
## Install with Docker

```
$ mkdir -p ~/.star-go && cd ~/.star-go
```

```
$ docker run -it --name star-go -v "$(pwd)":/tmp --rm sldennis/star-go init 3061ba66c81c7590e3b2a3bd3055fece429fb531
```

```
$ docker run -it --name star-go -v "$(pwd)":/tmp --rm sldennis/star-go focus teyushen/star-go teyushen/dockerfile golang/go
```