package fizzbuzz

import (
	"fmt"

	"github.com/DATA-DOG/godog"
)

var actual string

func iGot(num int) error {
	actual = fizzbuzz(num)
	return nil
}

func iCountWithFizzbuzz() error {
	return godog.ErrPending
}

func iShouldGet(expected string) error {
	if actual != expected {
		return fmt.Errorf("expected %s godogs to be remaining, but there is %s", expected, actual)
	}
	return nil
}

func nothing() error { return nil }

func FeatureContext(s *godog.Suite) {
	s.Step(`^I got (\d+)$`, iGot)

	s.Step(`^I count with fizzbuzz$`, nothing)
	s.Step(`^I should get "([^"]*)"$`, iShouldGet)

	s.BeforeScenario(func(interface{}) {
		actual = "" // clean the state before every scenario
	})
}
