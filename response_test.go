/**
    @author: potten
    @since: 2022/12/15
    @desc: //TODO
**/
package httpd

import (
	"fmt"
	"testing"
)

func TestHandler(t *testing.T) {
	r := newResponse()
	fmt.Println(r)
}
