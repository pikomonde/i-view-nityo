
## Run Project Locally

You might want to run the service in your local machine. This backend service runs on port `:5001`.

1. Setup `config.yaml` file

    You need to create `config.yaml` file in order to run the program. You can copy the config file from `config.example.yaml`.

    there are 2 options for `app.db_type`:

    - IN_MEMORY

      The database is saved in memory, so the data is not preserved when the service is restarted.

    - MYSQL

      The database is saved in mysql, so the data is preserved when the service is restarted.

    ```
    app:
      env: "local"
      db_type: "IN_MEMORY"
      port: 5001
      read_timeout: 5000
      write_timeout: 5000
      jwt_secret: "some_jwt_secret"
    mysql:
      host: "localhost"
      port: 3306
      username: "root"
      password: "password"
      db_name: "database_name"
    ```
   
2. Init MySQL

    `mysql -h localhost -u root -p < ./setup/deploy_00.00.001_init_schema.sql`

3. Run service

    `make run`

    After run the service you can see something like this in the command line:

    ```
    ======================================
                  ADMIN INFO
    ======================================
      username: admin
      password: 9WQKPX4OBJEnIV74cTTH1LD3
    ======================================
    ```

    This is the generated admin password that can be used to access the dashboard (create & disabled invitation token).


## API Documentation

You can try the API below by [download the postman data](https://www.getpostman.com/collections/3138bc3900143d6ab9b7). Or run it online via postman.

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/6422855-aa398be5-a3dd-407c-8fa7-4d2ec070ba2a?action=collection%2Ffork&collection-url=entityId%3D6422855-aa398be5-a3dd-407c-8fa7-4d2ec070ba2a%26entityType%3Dcollection%26workspaceId%3Dfcd8bab3-5721-426b-bf71-175eaf5b8353)

There are 5 API in this service:

1. Login

    `POST localhost:5001/api/login`

    The username and password of admin is written in the command line when the first time the service running. It will be something like this:

    ```
    ======================================
                  ADMIN INFO
    ======================================
      username: admin
      password: 9WQKPX4OBJEnIV74cTTH1LD3
    ======================================
    ```
    This endpoint will return JWT token with user role `admin`, that can be used to access invitation create, list, and disable endpoint.

    - Request Body:

      ```
      {
          "username": "admin",
          "password": "9WQKPX4OBJEnIV74cTTH1LD3"
      }
      ```

    - Response:

      ```
      {
          "status": 200,
          "data": {
              "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoie1wiaWRcIjoxLFwiaW52aXRhdGlvbl90b2tlblwiOlwiXCIsXCJ1c2VybmFtZVwiOlwiYWRtaW5cIixcInBhc3N3b3JkXCI6XCJcIixcInJvbGVcIjpcImFkbWluXCIsXCJjcmVhdGVkX2F0XCI6MH0iLCJleHAiOjE2NTgxNjM0NzB9.FrS5B_NkUsg1ZK4iNSAUUO6ZOKDV-FRpYpqdub9GPzo"
          }
      }
      ```
    
    Another example of incorrect password

    - Request Body:

      ```
      {
          "username": "admin",
          "password": "9WQKPX4OBJEnIV74cTTH1LD3x"
      }
      ```

    - Response:

      ```
      {
          "status": 200,
          "data": "Invalid Username / Password"
      }
      ```

2. Login Invitation

    `POST localhost:5001/api/login-invitation`

    This endpoint accept `invitation_token` that generated by admin.
    
    It will return another JWT token with user role `invitation` that can be used to access an invitation page.

    This endpoint will prevent user to try multiple times under 30 seconds. The server only identify same user based on **IP Address** and **User Agent**.

    - Request Body:

      ```
      {
          "invitation_token": "XoA4KX"
      }
      ```

    - Response:

      ```
      {
          "status": 200,
          "data": {
              "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoie1wiaWRcIjowLFwiaW52aXRhdGlvbl90b2tlblwiOlwiWG9BNEtYXCIsXCJ1c2VybmFtZVwiOlwiXCIsXCJwYXNzd29yZFwiOlwiXCIsXCJyb2xlXCI6XCJpbnZpdGF0aW9uXCIsXCJjcmVhdGVkX2F0XCI6MH0iLCJleHAiOjE2NTgxNjU5NTV9.nhIVytNooemd100RK4q2y0BCslfGwHxsZL5hKzIs5tc"
          }
      }
      ```

    Example of wrong `invitation_token`:

    - Request Body:

      ```
      {
          "invitation_token": "XoA4KXxxx"
      }
      ```

    - Response:

      ```
      {
          "status": 200,
          "data": "Invalid Invitation Token"
      }
      ```

    Example of hitting the endpoint twice under 30 seconds:

    - Request Body:

      ```
      {
          "invitation_token": "XoA4KXxxx"
      }
      ```

    - Response:

      ```
      {
          "status": 200,
          "data": "Try Again In 30 seconds"
      }
      ```

3. Invitation Create

    `POST localhost:5001/api/invitation/create`

    The request's cookie should contain the necessary JWT token.

    This endpoint will return randomized 6-12 characters alphanumeric `invitation_token`

    - Response
      
      ```
      {
          "status": 200,
          "data": {
              "invitation_token": "XoA4KX"
          }
      }
      ```

4. Invitation List

    `GET localhost:5001/api/invitation/list`

    The request's cookie should contain the necessary JWT token.

    This endpoint will return list of invitation data, such as `invitation_token`, `invitation_status`, and `invitation_created_at`

    - Response
      
      ```
      {
          "status": 200,
          "data": {
              "invitations": [
                  {
                      "id": 1,
                      "token": "SDQYvtKiwC",
                      "status": "inactive",
                      "created_at": 1658079309237972584
                  },
                  {
                      "id": 2,
                      "token": "XoA4KX",
                      "status": "inactive",
                      "created_at": 1658079311587521741
                  }
              ]
          }
      }
      ```

5. Invitation Disable

    `POST localhost:5001/api/invitation/disable`

    The request's cookie should contain the necessary JWT token.

    This endpoint will disable (revoke) an `invitation_token`.

    - Request Body

      ```
      {
          "invitation_token": "SDQYvtKiwC"
      }
      ```

    - Response
      
      ```
      {
          "status": 200,
          "data": "success"
      }
      ```

    Another example of wrong `invitation_token`:

    - Request Body

      ```
      {
          "invitation_token": "SDQYvtKiwCxx"
      }
      ```

    - Response
      
      ```
      {
          "status": 200,
          "data": "Invalid Invitation Token"
      }
      ```

