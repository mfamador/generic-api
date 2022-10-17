Feature: Retrieve Foos
  In order to retrieve the foos
  As a Data engineer
  I need to have a gRPC API to retrieve those values

  Scenario: Foo Sunny day
    Given I have a clean Foos table
    When I insert 10 values into Foos table
    And I query the Foo service
    Then I should have 10 entities
    And I should have 1 page

  Scenario: Paginated Foo
    Given I have a clean Readings table
    When I insert 100 values into Foos table
    And I query the Foo service with pages of size 10
    Then I should have 100 entities
    And I should have 10 pages

