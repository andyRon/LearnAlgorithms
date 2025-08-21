package main

import "fmt"

// 哈希算法的设计

// 加法哈希
func addHash(key string) int {
	var hash int64
	var modulus int64 = 1000000007

	for _, b := range []byte(key) {
		hash = (hash + int64(b)) % modulus
	}
	return int(hash)
}

// 乘法哈希
func mulHash(key string) int {
	var hash int64
	var modulus int64 = 1000000007

	for _, b := range []byte(key) {
		hash = (hash*31 + int64(b)) % modulus
	}
	return int(hash)
}

// 异或哈希
func xorHash(key string) int {
	hash := 0
	modulus := 1000000007
	for _, b := range []byte(key) {
		fmt.Println(int(b))
		hash ^= int(b)
		hash = (31*hash + int(b)) % modulus
	}
	return hash & modulus
}

// 旋转哈希
func rotHash(key string) int {
	var hash int64
	var modulus int64 = 1000000007

	for _, b := range []byte(key) {
		hash = ((hash << 4) ^ (hash >> 28) ^ int64(b)) % modulus
	}
	return int(hash)
}
