/*
Package closer is a one purpose (function) package that helps to call defer close.
*/
package closer

import (
	"io"
	"log"
)

// CloseOrLog is a helper for any defer closer.Close() call.
func CloseOrLog(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Printf("fail to close: %v", err)
	}
}
