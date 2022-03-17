### Product-Search-Challenge ###
This project was created to perform a tech interview for Walmart Back End's positions. To this date (17/3/2022) consists of:
    * Go 1.17.8
    * MongoDB Database (found here), provided by Walmart to be used as an example
    * MongoDB Client for connection to the DB.
    * Docker + DockerCompose + make commands.
    
The main idea was to provide a RESTAPI that consumes the provided DB , using Solid, Clean Code concepts and TDD. Case definitions can be found below.
The layer separations are done resembling the ports+adapters architectural pattern. They follow this criteria:
    - Application layer:
        - Adapters implementations, both inbound (inhttp package) and outbound (mongo db adapter and it's mongoclient which embeds). 
        - In the inhttp folder, you may find example responses. In the future this can be changed to allow .golden files, for example (would love to, but time is short)
    - Infrastructure layer:
        - HTTP Server (the "listen and serve" method of the standard library) is stored here.
        - Configurations needed to run the project (see below) can also be found here. Environmental variables are retrieved from this layer (to allow a basis for an infrastructure-as-code management in the future, in case you need it)
    - Domain layer:
        - This layer stores business logic (services), ports definitions (using interfaces). 

All layers are coupled using Polimorphism and dependency injection in the main.go file.

All Go files in the packages are named after their intended use. Utils with their tests are provided, as it was designed using TDD as much as time allowed me. 


FAQ: 

    - Why no frameworks like Gin?
        TBH, because I wanted to see what I could do with the standard library. Although I knew I would face limitations as I use Go at work as a Proffesional, it was a way to get revenge on little things that I've had to bypass and not give attention to, as I work under the stress of time. MAINLY, Routers . And reusing Test Code.
    - Why no tests on the MongoDBAdapter?
        I just didn't have the time.
    - Are you planning to add Integration tests?
        Of course. Could not add them until now as I was rushing with the assignment, and I work 9hs a day. 


How to:

    * To run the project:
        - Clone the project.
        - Run "make start-test-env"
        
        Note: remember to have docker installed.
    
    * To build a DockerImage:
        - Clone the project.
        - Run "make svc-docker-build"
    
    * To run that image:
        - Clone the project.
        - After building the image, run "make start-mock-db". 
            This command will run the test DB first, and register a network for it to be able to be reached from the Microservice. 
        - Then run "make svc-docker-run"

    * Test a package: 
        - Clone the project.
        - Run "go test ${path of package to test} -cover 

    * Clean after evaluating:
        - Run "make clean"
            This will eliminate generated files, stop docker instances and more.

    * Test endpoints:
        - Use curl:
        To test retrieval of products via ID:
            curl -X GET \
            'http://localhost:8080/api/products/505' \
            --header 'Accept: */*' \
        To test retrieval of products via text search:
            curl -X GET \
            'http://localhost:8080/api/products/search?text=voees' \
            --header 'Accept: */*' \

        Note: Replace localhost and 8080 if needed. To change the /api/products, just change it in the AddRoutes method in main.go.


Endpoints available:
    * To look for an exact ID:
        - GET from HOST:PORT/api/products/{product id to look for}
    

    * To look for a string in title OR  description fields:
        - GET from HOST:PORT/api/products/search?text={text to look for}


If you need further data, feel free to contact me.
The Walmart repo can be found here:
    - https://github.com/walmartdigital/products-db



Useful resources:
If you want to get a grasp of some parts implemented here, read from:
    * https://benhoyt.com/writings/go-routing/ - AWESOME article on path parameters in Go and routing.
    * https://www.honeybadger.io/blog/golang-logging/  - Great article on diff logging strategies and   practices
    * https://www.youtube.com/watch?v=oL6JBUk6tj0 - Great talk about different approaches to manage package distribution in Golang. 
