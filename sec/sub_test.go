
package acc

import "testing"

func TestSub(t *testing.T){

    got := sub(4, 6)
    want := -2

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

