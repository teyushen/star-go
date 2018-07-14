# star-go


## Prerequisites

* Create a github token ([how to generate a github token]())

	> You will get something like `3061ba66c81c7590e3b2a3bd3055fece429fb531`

## Determine how to use star-go

* Golang
* Docker

## Use with Golang

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
	
## Use with Docker

1. Create a directory

	```
	$ mkdir -p /.star-go && cd /.star-gow
	```

2. Initial the star-go
	```
	$ docker run -it --name star-go -v "$(pwd)":/root --rm sldennis/star-go init 3061ba66c81c7590e3b2a3bd3055fece429fb531
	```

3. Add you interesting github repositories

	```
	$ docker run -it --name star-go -v "$(pwd)":/root --rm sldennis/star-go focus teyushen/star-go teyushen/dockerfile golang/go
	```

4. Order the numbers of size of repositories you are interested

	```
	$ docker run -it --name star-go -v "$(pwd)":/root --rm sldennis/star-go focus teyushen/star-go teyushen/dockerfile golang/go
	```