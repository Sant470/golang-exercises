// ----------------
// Context Value Bag
// ----------------

// data is stored inside the context with a combination of a key and then the value.
// But the key isn't just some value, it's also gonna be based on a type.
// Ideally useful for storing secondary datas (metrics and analytics based data)

package main

import (
	"context"
	"fmt"
)

//TraceID represents the trace id
type TranceID string

// TraceIDKey is the type of value to use for the key. The Key is type specific and only value of the same type will match.

type TranceIDKey int

func main() {
	tranceID := TranceID("23e4567-e89b-12d3-a456-426614174000")
	const tranceIDKey TranceIDKey = 0

	// store the tranceID value inside the context with a value of zero for the TranceIDKey type
	ctx := context.WithValue(context.Background(), tranceIDKey, tranceID)

	// retrieve the tranceID value from context value bag
	if uuid, ok := ctx.Value(tranceIDKey).(TranceID); ok {
		fmt.Println("tranceID: ", uuid)
	}

	// retrieve the tranceID value from the context value bag not using the proper key type.
	if _, ok := ctx.Value(0).(TranceID); !ok {
		fmt.Println("tranceID Not Found")
	}
}
