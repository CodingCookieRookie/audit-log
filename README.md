# Audit-Log
<h5> Service for user to get and post any event log
 
1. Git clone repository into audit-log folder
2. Open a new terminal in the audit-log folder directory and run cmd `./start.sh` (Make sure no other existing docker containers are running, else comment out docker kill and rm in start.sh file)
3. Curl endpoints depending on usage

<h2> Curl commands
    <h4> User get event 
    <p>
    curl localhost:3000/users/events\?event_type=bill\&event_timestamp_start=2023-05-02_17:04:04\&event_timestamp_end=2023-06-04_22:04:05\&gmt=*8 -H "email:alvinchee98@gmail.com" -H "token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFsdmluY2hlZTk4QGdtYWlsLmNvbSIsImV4cCI6MTY4ODExNDQyNH0.sXUYUdGRvIKOqtSEKWJkE9Q7CI2JK4R_0ZubciG8ZfE"

    curl localhost:3000/users/events\?event_type=bill\&event_order=DESC -H "email:alvinchee98@gmail.com" -H "token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFsdmluY2hlZTk4QGdtYWlsLmNvbSIsImV4cCI6MTY4ODExNDQyNH0.sXUYUdGRvIKOqtSEKWJkE9Q7CI2JK4R_0ZubciG8ZfE"
    </p>

    <h3> User post event
    curl -H "email:alvinchee98@gmail.com" -H "token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFsdmluY2hlZTk4QGdtYWlsLmNvbSIsImV4cCI6MTY4ODExNDQyNH0.sXUYUdGRvIKOqtSEKWJkE9Q7CI2JK4R_0ZubciG8ZfE" -d '{"event_type":"bill", "event_data":{"username":"alvin", "amount":311, "currency":"USD"}}' -X POST localhost:3000/users/events

    <h3> Staff send token to user via email<br>
    curl localhost:3000/api/token?email=alvinchee98@gmail.com -H "app-secret:Up0F9YrxSDruZlKAxgSiKfdZp7EB8D4XY5vWtbhElHw="