package fizzbuzz

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFizzBuzz(test *testing.T) {
	Convey("Fizzzbuzz 1 should return 1", test, func() {
		So(FizzBuzz(1), ShouldEqual, "1")
	})

	Convey("Fizzbuzz 2 should return 2", test, func() {
		So(FizzBuzz(2), ShouldEqual, "2")
	})

	Convey("Fizzbuzz 3 should return Fizz", test, func() {
		So(FizzBuzz(3), ShouldEqual, "Fizz")
	})

	Convey("Fizzbuzz 5 should return Buzz", test, func() {
		So(FizzBuzz(5), ShouldEqual, "Buzz")
	})

	Convey("Fizzbuzz 15 should return FizzBuzz", test, func() {
		So(FizzBuzz(15), ShouldEqual, "FizzBuzz")
	})
}