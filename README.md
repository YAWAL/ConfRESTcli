REST client for the config service
----------------------------------

This is a REST client for the Config service. It converts http requests to grpc, allowing CRUD operation for configs. 


To find a config using its **type** and **name** make a GET request
````````````````````````````
http://host:port/getConfig/type/name
````````````````````````````
To find all configs of a given **type** make a GET request
````````````````````````````
http://host:port/getConfig/type
````````````````````````````
To add a config make a POST request with a new config JSON-object in body
````````````````````````````
http://host:port/createConfig/type
````````````````````````````
To delete a config using its **type** and **name** make a DELETE request
````````````````````````````
http://host:port/deleteConfig/type/name
````````````````````````````
To update a config make a PUT request with a config JSON-object in body
````````````````````````````
http://host:port/updateConfig/type
````````````````````````````
To check if the application is running make a GET request
````````````````````````````
http://host:port/info
````````````````````````````
How to start

To install dep  dependency management tool run 
````````````````````````````
make install dep
````````````````````````````
To install application dependencies run
````````````````````````````
make dependencies
````````````````````````````

To to build the application
````````````````````````````
make build
````````````````````````````
To to run the application
````````````````````````````
make run
````````````````````````````
To run tests

``````````````````
make tests
``````````````````
To to run the application in docker container
````````````````````````````
make docker-build
````````````````````````````
