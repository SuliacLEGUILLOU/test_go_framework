package utils

import "testing"
 
func TestBasic(t *testing.T) {
    one := FizzOrBuzz(1, 3, 5, "fizz", "buzz")
    three := FizzOrBuzz(3, 3, 5, "fizz", "buzz")
    five := FizzOrBuzz(5, 3, 5, "fizz", "buzz")
    fifteen := FizzOrBuzz(15, 3, 5, "fizz", "buzz")

    if (one != "1"){
        t.Fail()
    }
    if (three != "fizz"){
        t.Fail()
    }
    if (five != "buzz"){
        t.Fail()
    }
    if (fifteen != "fizzbuzz"){
        t.Fail()
    }
}

func TestVariation(t *testing.T) {
    one := FizzOrBuzz(1, 4, 5, "fizz", "buzz")
    three := FizzOrBuzz(3, 4, 5, "fizz", "buzz")
    four := FizzOrBuzz(4, 4, 5, "fizz", "buzz")
    fifteen := FizzOrBuzz(15, 4, 5, "fizz", "buzz")
    twenty := FizzOrBuzz(20, 4, 5, "fizz", "buzz")

    if (one != "1"){
        t.Fail()
    }
    if (three != "3"){
        t.Fail()
    }
    if (four != "fizz"){
        t.Fail()
    }
    if (fifteen != "buzz"){
        t.Fail()
    }
    if (twenty != "fizzbuzz"){
        t.Fail()
    }
}

func TestCustomStr(t *testing.T) {
    res := FizzOrBuzz(15, 3, 5, "fish", "plant")

    if (res != "fishplant"){
        t.Fail()
    }
}
