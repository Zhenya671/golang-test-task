# recipe-service
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