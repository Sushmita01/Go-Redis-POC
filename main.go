package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {

	// newPool returns a pointer to a redis.Pool
	pool := newPool()
	// get a connection from the pool (redis.Conn)
	conn := pool.Get()
	// use defer to close the connection when the function completes
	defer conn.Close()

	// call Redis PING command to test connectivity
	err := ping(conn)
	if err != nil {
		fmt.Println(err)
	}

	// set demonstrates the redis SET command using a simple
	// string key:value pair
	err = setConfig(conn,"549646412US",`{
		"crid": "549646412US",
			"config": {
			"SoftCaps": [100.0],
				"SoftCapMultipliers": [0.2],
				"HardCap": 100.0,
				"FloorPrice": 0.00,
				"Multiplier": 1.5,
				"UrlL2ACascadingThreshold": 500,
				"L2ABiasCascadingThreshold": 2000,
				"UrlRPCCascadingThreshold": 4,
				"RPCBiasCascadingThreshold": 50,
				"BiasRPCWeight": 0,
				"URLRPCWeight": 4,
				"BiasL2AWeight": 0,
				"URLL2AWeight": 0,
				"ConsiderVisitor": 1.0,
				"L2AModel": 1.0,
				"RPCModel": 1.0,
				"L2AModelName": "549646412US",
				"RPCModelName": "549646412US",
				"ExplorationMultiplier": 1.5
		},
		"experiments": [
		],
			"rpcFile": "rpc549646412US",
			"l2aFile": "l2a549646412US"
	}`)
	if err != nil {
		fmt.Println(err)
	}

	err = setConfig(conn,"703626483US",`"crid": "703626483US",
    "config": {
      "SoftCaps": [3, 9],
      "SoftCapMultipliers": [0.7, 0.2],
      "HardCap": 17.0,
      "FloorPrice": 0.00,
      "Multiplier": 1.0,
      "UrlL2ACascadingThreshold": 200,
      "L2ABiasCascadingThreshold": 2000,
      "UrlRPCCascadingThreshold": 3,
      "RPCBiasCascadingThreshold": 6,
      "BiasRPCWeight": 0,
      "URLRPCWeight": 0.0,
      "BiasL2AWeight": 0,
      "URLL2AWeight": 0.0,
      "ConsiderVisitor": 1.0,
      "L2AModel": 1.0,
      "RPCModel": 1.0,
      "L2AModelName": "703626483US_CONTROL",
      "RPCModelName": "703626483US_CONTROL",
      "ExplorationMultiplier": 0.6
    },
    "experiments": [
      {
        "key": "RPCIPBias",
        "buckets": [
          {
            "value": "0.0",
            "traffic": 50,
            "name": ""
          },
          {
            "value": "0.0",
            "traffic": 50,
            "name": ""
          }
        ]
      }
    ],
    "rpcFile": "rpc703626483US",
    "l2aFile": "l2a703626483US"
  }`)

	err = setConfig(conn,"286855778US",`"crid": "286855778US",
    "config": {
      "SoftCaps": [15.0],
      "SoftCapMultipliers": [0.2],
      "HardCap": 20.0,
      "FloorPrice": 0.00,
      "Multiplier": 0.9,
      "UrlL2ACascadingThreshold": 200,
      "L2ABiasCascadingThreshold": 2000,
      "UrlRPCCascadingThreshold": 3,
      "RPCBiasCascadingThreshold": 50,
      "BiasRPCWeight": 0,
      "URLRPCWeight": 2,
      "BiasL2AWeight": 0,
      "URLL2AWeight": 0,
      "ConsiderVisitor": 1.0,
      "L2AModel": 1.0,
      "RPCModel": 1.0,
      "L2AModelName": "286855778US",
      "RPCModelName": "286855778US",
      "ExplorationMultiplier": 0.5
    },
    "experiments": [
    ],
    "rpcFile": "rpc286855778US",
    "l2aFile": "l2a286855778US"
  }`)


	err = setConfig(conn,"451032526US",`{
    "crid": "451032526US",
    "config": {
      "SoftCaps": [1.0],
      "SoftCapMultipliers": [0.2],
      "HardCap": 1.0,
      "FloorPrice": 0.00,
      "Multiplier": 0.0,
      "UrlL2ACascadingThreshold": 500,
      "L2ABiasCascadingThreshold": 2000,
      "UrlRPCCascadingThreshold": 6,
      "RPCBiasCascadingThreshold": 50,
      "BiasRPCWeight": 0,
      "URLRPCWeight": 0.2,
      "BiasL2AWeight": 0,
      "URLL2AWeight": 0.33,
      "ConsiderVisitor": 1.0,
      "L2AModel": 1.0,
      "RPCModel": 1.0,
      "L2AModelName": "451032526US",
      "RPCModelName": "451032526US",
      "ExplorationMultiplier": 0.6
    },
    "experiments": [
    ],
    "rpcFile": "rpc451032526US",
    "l2aFile": "l2a451032526US"
  }`)
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

// ping tests connectivity for redis (PONG should be returned)
func ping(c redis.Conn) error {
	// Send PING command to Redis
	// PING command returns a Redis "Simple String"
	// Use redis.String to convert the interface type to string
	s, err := redis.String(c.Do("PING"))
	if err != nil {
		return err
	}

	fmt.Printf("PING Response = %s\n", s)
	// Output: PONG

	return nil
}

func setConfig(c redis.Conn,crid string,config string) error {

	_, error := c.Do("SET", crid, config)
	if error != nil {
		return err
	}


return nil
}



