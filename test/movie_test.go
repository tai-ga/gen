package models

import (
	"testing"
)

type test struct {
	Exec     func() interface{}
	Expected interface{}
}

func getTests() map[string][]test {
	// the basic pattern for tests is zero/one/many for 'some' slice & sanity checks on 'none' slice
	tests := make(map[string][]test)

	tests["AggregateInt"] = []test{
		test{
			func() interface{} {
				return some.AggregateInt(sum_theaters)
			},
			6 + 9 + 5,
		},
		test{
			func() interface{} {
				return none.AggregateInt(sum_theaters)
			},
			0,
		},
	}

	tests["AggregateString"] = []test{
		test{
			func() interface{} {
				return some.AggregateString(concat_title)
			},
			"first" + "second" + "third",
		},
		test{
			func() interface{} {
				return none.AggregateString(concat_title)
			},
			"",
		},
	}

	tests["All"] = []test{
		test{
			func() interface{} {
				return some.All(is_dummy)
			},
			false,
		},
		test{
			func() interface{} {
				return some.All(is_first)
			},
			false,
		},
		test{
			func() interface{} {
				return some.All(is_first_or_second_or_third)
			},
			true,
		},
		test{
			func() interface{} {
				return none.All(is_false)
			},
			true,
		},
		test{
			func() interface{} {
				return none.All(is_true)
			},
			true,
		},
	}

	tests["Any"] = []test{
		test{
			func() interface{} {
				return some.Any(is_dummy)
			},
			false,
		},
		test{
			func() interface{} {
				return some.Any(is_first)
			},
			true,
		},
		test{
			func() interface{} {
				return some.Any(is_first_or_third)
			},
			true,
		},
		test{
			func() interface{} {
				return none.Any(is_false)
			},
			false,
		},
		test{
			func() interface{} {
				return none.Any(is_true)
			},
			false,
		},
	}

	tests["Count"] = []test{
		test{
			func() interface{} {
				return some.Count(is_dummy)
			},
			0,
		},
		test{
			func() interface{} {
				return some.Count(is_first)
			},
			1,
		},
		test{
			func() interface{} {
				return some.Count(is_first_or_third)
			},
			2,
		},
		test{
			func() interface{} {
				return some.Count(is_true)
			},
			len(some),
		},
		test{
			func() interface{} {
				return none.Count(is_false)
			},
			0,
		},
		test{
			func() interface{} {
				return none.Count(is_true)
			},
			0,
		},
	}

	tests["JoinString"] = []test{
		test{
			func() interface{} {
				return some.JoinString(get_title, ", ")
			},
			"first, second, third",
		},
		test{
			func() interface{} {
				return none.JoinString(get_title, ", ")
			},
			"",
		},
	}

	tests["SumInt"] = []test{
		test{
			func() interface{} {
				return some.SumInt(get_theaters)
			},
			6 + 9 + 5,
		},
		test{
			func() interface{} {
				return none.SumInt(get_theaters)
			},
			0,
		},
	}

	tests["Skip"] = []test{
		test{
			func() interface{} {
				return some.Skip(0)
			},
			some,
		},
		test{
			func() interface{} {
				return some.Skip(1)
			},
			Movies{_second, _third},
		},
		test{
			func() interface{} {
				return some.Skip(2)
			},
			Movies{_third},
		},
		test{
			func() interface{} {
				return some.Skip(3)
			},
			Movies{},
		},
		test{
			func() interface{} {
				return some.Skip(99)
			},
			Movies{},
		},
		test{
			func() interface{} {
				return none.Skip(1)
			},
			Movies{},
		},
	}

	tests["Take"] = []test{
		test{
			func() interface{} {
				return some.Take(0)
			},
			Movies{},
		},
		test{
			func() interface{} {
				return some.Take(1)
			},
			Movies{_first},
		},
		test{
			func() interface{} {
				return some.Take(2)
			},
			Movies{_first, _second},
		},
		test{
			func() interface{} {
				return some.Take(3)
			},
			some,
		},
		test{
			func() interface{} {
				return some.Take(4)
			},
			some,
		},
		test{
			func() interface{} {
				return none.Take(1)
			},
			Movies{},
		},
	}

	tests["Where"] = []test{
		test{
			func() interface{} {
				return some.Where(is_dummy)
			},
			Movies{},
		},
		test{
			func() interface{} {
				return some.Where(is_first)
			},
			Movies{_first},
		},
		test{
			func() interface{} {
				return some.Where(is_first_or_third)
			},
			Movies{_first, _third},
		},
		test{
			func() interface{} {
				return some.Where(is_true)
			},
			some,
		},
		test{
			func() interface{} {
				return none.Where(is_false)
			},
			Movies{},
		},
		test{
			func() interface{} {
				return none.Where(is_true)
			},
			Movies{},
		},
	}

	return tests
}

func TestAll(t *testing.T) {
	for _, tests := range getTests() {
		for _, test := range tests {
			switch test.Expected.(type) {
			default:
				got := test.Exec()
				if got != test.Expected {
					t.Errorf("Expected %v, got %v", test.Expected, got)
				}
			case Movies:
				got := test.Exec().(Movies)
				exp := test.Expected.(Movies)
				if len(got) != len(exp) {
					t.Errorf("Expected %v Movies, got %v", len(exp), len(got))
					break
				}
				for i := range got {
					if got[i] != exp[i] {
						t.Errorf("Expected %v, got %v", exp[i], got[i])
					}
				}
			}
		}
	}
}

var _first = &Movie{Title: "first", Theaters: 6}
var _second = &Movie{Title: "second", Theaters: 9}
var _third = &Movie{Title: "third", Theaters: 5}

var some = Movies{
	_first,
	_second,
	_third,
}

var none = Movies{}

var is_first = func(movie *Movie) bool {
	return movie.Title == "first"
}
var is_first_or_second_or_third = func(movie *Movie) bool {
	return movie.Title == "first" || movie.Title == "second" || movie.Title == "third"
}
var is_first_or_third = func(movie *Movie) bool {
	return movie.Title == "first" || movie.Title == "third"
}
var is_dummy = func(movie *Movie) bool {
	return movie.Title == "dummy"
}
var is_true = func(movie *Movie) bool {
	return true
}
var is_false = func(movie *Movie) bool {
	return false
}
var get_theaters = func(movie *Movie) int {
	return movie.Theaters
}
var sum_theaters = func(movie *Movie, acc int) int {
	return acc + movie.Theaters
}
var get_title = func(movie *Movie) string {
	return movie.Title
}
var concat_title = func(movie *Movie, acc string) string {
	return acc + movie.Title
}
