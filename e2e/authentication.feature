Feature: Authentication

The PocketBase backend supports authentication out of the box both with email+password and several OAuth providers - e.g. Google, Facebook.
This feature description is mostly related to the web client implementation.

Scenario: Unauthenticated user loads a protected route
    Given an unauthenticated user
    When loading a protected route
    Then is being redirected to the route "/login-signup"

Scenario: Signup new user
    Given an unauthenticated user
    And is on route "/login-signup"
    When a signup form is submitted with data "{ name: '', email: '', password: '', passwordConfirm: '' }"
    Then a token is set (as cookie or local storage?)
    And the user has a reference to their newly created data in Brevo
    And is being redirected to the route "/home" (or "/dashboard")

Scenario: Login user
    Given an unauthenticated user
    And is on route "/login-signup"
    When a signup form is submitted with data "{ email: '', password: '' }"
    Then a token is set (as cookie or local storage?)
    And the user has a reference to their newly created data in Brevo
    And is being redirected to the route "/home" (or "/dashboard")

Scenario: Authenticated user loads a protected route
    Given an authenticated user loads a protected route
    Then the token is being refreshed
