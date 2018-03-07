REST client for the config service
----------------------------------

Converts http requests to grpc, allowing CRUD operation for configs.


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
How to start

Run to install dep
````````````````````````````
make install dep
````````````````````````````
Run to install application dependencies
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
To to run the application in docker container
````````````````````````````
make docker-build
````````````````````````````
