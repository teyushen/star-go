![step1](https://github.com/teyushen/star-go/blob/master/logo.png)


[![Build Status](https://travis-ci.org/teyushen/star-go.svg?branch=master)](https://travis-ci.org/teyushen/star-go)   ![MIT](https://img.shields.io/packagist/l/doctrine/orm.svg)

## Features

Sometimes we want to trace Github Projects we are interesed. And we always want to know how many numbers of ⭐️ they get now.

This project can easily help us focus on projects we are interested, and can also help us order these projects by numbers of ⭐️.


## Demo

-  ![step1](https://github.com/teyushen/star-go/blob/master/images/star-go-ls.gif)

-  ![step1](https://github.com/teyushen/star-go/blob/master/images/star-go-c.gif)


## Prerequisites

* Create a github token ([how to generate a github token](https://github.com/teyushen/star-go/tree/master/images))

	> You will get something like `3061ba66c81c7590e3b2a3bd3055fece429fb531`

## Determine how to use star-go

* [Golang](https://golang.org/doc/install) 
* [Docker](https://docs.docker.com/install/) 

## Use with Golang

1. Install star-go

	```
	$ go get github.com/teyushen/star-go 
	$ go install github.com/teyushen/star-go 
	```


2. Initial the star-go

	```
	$ star-go init <token>
	```
	 
	> e.g.
	> `$ star-go init 3061ba66c81c7590e3b2a3bd3055fece429fb531`
	

3. Add you interesting github repositories

	```
	$ star-go focus <owner/repository>...
	```
	
	> e.g.
	> `$ star-go focus teyushen/star-go teyushen/dockerfile golang/go`

4. Order the numbers of size of repositories you are interested

	```
	$ star-go compare
	```
	
## Use with Docker

1. Create a directory

	- Linux or Mac
	
		```
		$ mkdir -p ~/.star-go
		```

2. Initial the star-go
	
	- Linux or Mac
	
		```
		$ docker run -it -v ~/.star-go:/root/.star-go --rm sldennis/star-go init <token>
		```
	
		> e.g. 
		> `docker run -it -v ~/.star-go:/root/.star-go --rm sldennis/star-go init 3061ba66c81c7590e3b2a3bd3055fece429fb531
`

3. Add you interesting github repositories

	- Linux or Mac
	
		```
		$ docker run -it -v ~/.star-go:/root/.star-go --rm sldennis/star-go focus <owner/repository>...
		```
	
		> e.g.
		> `docker run -it -v ~/.star-go:/root/.star-go --rm sldennis/star-go focus teyushen/star-go teyushen/dockerfile golang/go`
	
4. Order the numbers of size of repositories you are interested

	- Linux or Mac
	
		```
		$ docker run -it -v ~/.star-go:/root/.star-go --rm sldennis/star-go compare
		```
		
## Authors

-	[Blair Lee](https://github.com/blairlee227) - *Logo Designer* - [dribbble](https://dribbble.com/blairlee)
		
## License

MIT
