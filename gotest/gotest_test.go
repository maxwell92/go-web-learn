package gotest

import (
    "testing"
)


func Test_Division_1(t *testing.T) {
    if i, e := Division(6, 2); i != 3 || e != nil {
        t.Error("Division test didn't pass!")
    } else {
        t.Log("test 1 passed!")
    }
}

/*func Test_Division_2(t *testing.T) {
 *    t.Error("Can't Pass!")
 *}
 */

func Test_Division_2(t *testing.T) {
    if _ ,e := Division(6, 0); e == nil {
        t.Error("Division did not work as expected!")
    } else {
        t.Log("one test passwd.", e)
    }
}



