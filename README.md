# go's test task
##
The problem

Imagine that you are studying to be a programmer at university and you are constantly being approached by people in your class who don't understand algorithms, asking you to help them with a solution for money. One day you get the idea that you should create an application that does all the work for you, and all you have to do is collect money from the people who ask you for it.

The task is

To implement a service that provides the following functionality:

Creation of a new student using his full name and group number, for example, Stirlitz Ivan Vasilievich IP-394. For convenience of further queries, each user can be assigned a login, which will act as a shortcut to that student.
Changing the student's debt (decrease and increase). The maximum debt is 1000 roubles. When the maximum debt is reached, the application should not give answers to new requests for solving algorithmic problems.
Getting an answer for the algorithmic task. Each task has a different cost for getting an answer, so you need to change the debt of the student for whom you want to get an answer (don't forget about the maximum debt).
The only algorithmic task that can be answered and must be implemented:

Given an unsorted array of N numbers from 1 to N,
some numbers from the range [1, N] are missing,
and some are present twice.

Find all the missing numbers.
Service requirements

The language is Golang. Any frameworks and libraries can be used.
Use PostgreSQL as a database.
All code must be uploaded to Github or Gitlab with a Readme file containing launch instructions and example queries/responses.
If there are any questions about the TOR, leave the decision to the candidate (in this case, the readme file to the project should list the issues the candidate encountered and how he/she solved them).

Development of the interface in a browser is NOT required. For testing you can use any tool you like. For example: in the terminal via curl or Postman.

Will be a plus

If you additionally do something from this list, you'll noticeably stand out from the other candidates!

Covering the code with tests.
Swagger for the API.
Using docker and docker-compose to bring up and deploy a working environment.
Primitive CI/CD (only builds and auto-tests are sufficient).
Makefile.

###
- copy paste in your postman to load prepared endpoints
 [softweather-test-task.postman_collection.json](..%2FDownloads%2Fsoftweather-test-task.postman_collection.json)  (don't forget to pass auth token as well as set your host (see logs))

- make env files and fill them up
- ### Start in docker with port specified in the env vars:
    ~~~bash
         make start
- ### Stop:
    ~~~bash
         make stop
- ### see app's logs:
    ~~~bash
         make log_app
- ### see db's logs:
    ~~~bash
         make log_db
- ### update swagger documentation:
    ~~~bash
         make swagger