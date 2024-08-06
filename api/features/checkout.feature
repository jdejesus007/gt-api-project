Feature: Checkout customer and verify order creation
  In order to know if order was created
  I need to be able to request checkout

  Scenario: Successfully create customer
    When I send "POST" request to "/customers/"
      """
      {
        "firstName": "John",
        "lastName": "john.doe@test.com",
        "email": "john.doe@test.com"
      }
      """
    Then the response code should be 201
    When I send "POST" request to "/checkout"
      """
      {
      "email": "{email}",
      "books": [
          {
            "UUID": "ba398055-8df8-497a-af1d-bc2fcf20b03d"
          },
          {
            "UUID": "662e342d-0929-4cd5-bd9a-6b9913d61b71"
          }
        ]
      }
      """
    Then the response code should be 201
    And the response should match json:
    """
    {
      "customerUUID": "{customerUUID}",
      "bookUUIDs": [
          "ba398055-8df8-497a-af1d-bc2fcf20b03d",
          "662e342d-0929-4cd5-bd9a-6b9913d61b71"
      ]
    }
    """
