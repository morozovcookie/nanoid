package nanoid_test

import (
	"crypto/rand"
	"fmt"

	"github.com/aidarkhanov/nanoid"
)

func ExampleFormat() {
	alphabet := "-0123456789ABCDEFGHIJKLNQRTUVWXYZ_cfgijkpqtvxz"
	size := 21

	// randomBytesGenerator returns random bytes buffer
	// to pass it in Format function
	randomBytesGenerator := func(step int) ([]byte, error) {
		buffer := make([]byte, step)
		if _, err := rand.Read(buffer); err != nil {
			return nil, err
		}
		return buffer, nil
	}

	id, err := nanoid.Format(randomBytesGenerator, alphabet, size)
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
}

func ExampleGenerate() {
	alphabet := "-0123456789ABCDEFGHIJKLNQRTUVWXYZ_cfgijkpqtvxz"
	size := 21

	id, err := nanoid.Generate(alphabet, size)
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
}

func ExampleMustFormat() {
	alphabet := "-0123456789ABCDEFGHIJKLNQRTUVWXYZ_cfgijkpqtvxz"
	size := 21

	// randomBytesGenerator returns random bytes buffer
	// to pass it in Format function
	randomBytesGenerator := func(step int) ([]byte, error) {
		buffer := make([]byte, step)
		if _, err := rand.Read(buffer); err != nil {
			return nil, err
		}
		return buffer, nil
	}

	id := nanoid.MustFormat(randomBytesGenerator, alphabet, size)

	fmt.Println(id)
}

func ExampleMustGenerate() {
	alphabet := "-0123456789ABCDEFGHIJKLNQRTUVWXYZ_cfgijkpqtvxz"
	size := 21

	id := nanoid.MustGenerate(alphabet, size)

	fmt.Println(id)
}

func ExampleNew() {
	id := nanoid.New()

	fmt.Println(id)
}