package main

import (
	"fmt"
	// "./matchstring"
	// "./random"
	// "./systemtime"
	// "./crypto"
	// "./readenv"
)

func main() {
	fmt.Println("Golang Utils")
	fmt.Println("==================================================")

	// result, _ := matchstring.Match09azAZ("123ahZXCshs567aaaa") // true
	// fmt.Println("123ahZXCshs567aaaa,", result)
	// result, _ = matchstring.Match09azAZ("123ahZXCshs567aa!a") // false
	// fmt.Println("123ahZXCshs567aa!a,", result)

	// rand := random.Random(10, 100)
	// fmt.Println(rand)

	// time := systemtime.GetSysTimeInSec()
	// fmt.Println("second:    ", time)
	// time = systemtime.GetSysTimeInNanoSec()
	// fmt.Println("nanosecond:", time)

	// sha3 := crypto.Sha3Encrypt("test sha3 encrypt", 256)
	// fmt.Println("sha3-256: " + sha3)
	// sha3 = crypto.Sha3Encrypt("test sha3 encrypt", 384)
	// fmt.Println("sha3-384: " + sha3)

	// readEnv, _ := readenv.LoadFile(".env")
	// fmt.Println("key1: " + readEnv["key1"])
	// fmt.Println("key2: " + readEnv["key2"])
	// fmt.Println("key3: " + readEnv["key3"])
}
