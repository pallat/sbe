Feature: fizzbuzz
  A counting game for children
    three divisible is Fizz
    five divisible is Buzz
    three and five divisible is FizzBuzz

  Scenario: Count 1
    Given I got 1
    When I count with fizzbuzz
    Then I should get "1"

  Scenario: Count 3
    Given I got 3
    When I count with fizzbuzz
    Then I should get "Fizz"