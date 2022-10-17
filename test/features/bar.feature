Feature: Retrieve Bars
  In order to retrieve the bars
  As a Data engineer
  I need to have a gRPC API to retrieve those values

  Scenario: Bar Sunny day
    Given I have a clean Foos table
    When I insert 10 values into Foos table
    And I query the Bar service
    Then I should have 10 entities
    And I should have 1 page

  Scenario: Paginated Bar
    Given I have a clean Readings table
    When I insert 100 values into Bars table
    And I query the Bar service with pages of size 10
    Then I should have 100 entities
    And I should have 10 pages

