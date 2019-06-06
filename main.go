package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"simulator/engine"
)

func main() {

	engine.Simulate("resources/data.csv")



	pool := newPool()

	conn := pool.Get()

	defer conn.Close()


	err := ping(conn)
	if err != nil {
		fmt.Println("Failed to connect to DB")
		fmt.Println(err)
	}


	err = get(conn,"549646412US")
	if err != nil {
		fmt.Println(err)
	}

	err = get(conn,"286855778US")
	if err != nil {
		fmt.Println(err)
	}

	err = get(conn,"703626483US")
	if err != nil {
		fmt.Println(err)
	}

	err = get(conn,"451032526US")
	if err != nil {
		fmt.Println(err)
	}

}

func newPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// ping tests connectivity for redis
func ping(c redis.Conn) error {
	s, err := redis.String(c.Do("PING"))
	if err != nil {
		return err
	}

	fmt.Printf( s)
	return nil
}


// get executes the redis GET command
func get(c redis.Conn,crid string) error {

	s, err := redis.String(c.Do("GET", crid))
	if err != nil {
		return (err)
	}
	fmt.Printf("%s = %s\n", crid, s)

	return nil
}



