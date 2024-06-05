Feature: Public Dashboard

As a user of the app I can create trading-accounts, upload my trades via a txt-file exported from Sierra Charts and make parts of my stats publicly accessible on a public dashboard.
As a user I want granular control over what categories of data can be viewed by the public - i.e. unauthenticated website visitors. The initial version of the public dashboard will look similar to /:lang/overview

Implementation planning:
- A database table "public_dashboard_permissions" that holds a user_id and a list of permitted public resources of the user. Public resources can be read (HTTP GET), but not mutated by anybody (except for the owner of the data itself, of course).
- Pocketbase permissions for `list` and `get` need to follow a logic: allow `list` of `get`, if a) the request is issued from the public url (e.g. /:lang/public/:userId/:tradingAccountId/overview) OR b) the user is the owner of the resource AND c) the permission is granted through database table "public_dashboard_permissions" --> (a || b) && c

Scenario:
    Given a user's `trade` records are public according to permissions table
    When an anonymous user visits /en/public/:userId/:accountId/dashboard
    Then a list and a chart of the `trade` records are displayed

Scenario:
    Given a user's `trade` records are NOT public according to permissions table
    When an anonymous user visits /en/public/:userId/:accountId/dashboard
    Then no `trade` records are displayed

Scenario:
    Given an authenticated user with an existing trading accountId
    When the go to the settings for their public dashboard permissions
    Then the current permissions for each resource type (e.g. `trades`) are shown
    And the permissions can be updated