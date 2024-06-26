# Audit-Log
##### Service for user to get and post any event log

## Run with docker
1. Git clone repository in any desired folder
2. Enter the audit-log folder with `cd audit-log`
3. Inside the audit-log folder, run `./docker-start.sh`
4. Curl endpoints depending on usage

## Run locally (with go installed on machine)
1. Git clone repository in any desired folder
2. Enter the audit-log folder with `cd audit-log`
3. Create a new .env file and copy-paste the content from the env.txt file in the email to the .env file
4. Open a new terminal in the audit-log folder directory and run cmd `./go-start.sh` (Make sure no other existing docker containers are running, else comment out docker kill and rm in start.sh file)
5. Curl endpoints depending on usage

## Run locally (without go installed on machine)
1. Git clone repository in any desired folder
2. Enter the audit-log folder with `cd audit-log`
3. Create a new .env file and copy-paste the content from the env.txt file in the email to the .env file
4. Open a new terminal in the audit-log folder directory and run cmd `./ngo-start.sh` (Make sure no other existing docker containers are running, else comment out docker kill and rm in start.sh file)
5. Curl endpoints depending on usage

### Curl command examples
#### User get event 
<p>

    curl localhost:3000/users/events\?event_type=bill\&event_timestamp_start=2023-05-02_17:04:04\&event_timestamp_end=2023-06-04_22:04:05\&gmt=*8 -H "email:alvinchee98@gmail.com" -H "token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFsdmluY2hlZTk4QGdtYWlsLmNvbSIsImV4cCI6MTY4ODExNDQyNH0.sXUYUdGRvIKOqtSEKWJkE9Q7CI2JK4R_0ZubciG8ZfE"
</p>

<p>

    curl localhost:3000/users/events\?event_type=bill\&event_order=DESC -H "email:alvinchee98@gmail.com" -H "token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFsdmluY2hlZTk4QGdtYWlsLmNvbSIsImV4cCI6MTY4ODExNDQyNH0.sXUYUdGRvIKOqtSEKWJkE9Q7CI2JK4R_0ZubciG8ZfE"
</p>

#### User post event
<p>

    curl -H "email:alvinchee98@gmail.com" -H "token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFsdmluY2hlZTk4QGdtYWlsLmNvbSIsImV4cCI6MTY4ODExNDQyNH0.sXUYUdGRvIKOqtSEKWJkE9Q7CI2JK4R_0ZubciG8ZfE" -d '{"event_type":"bill", "event_data":{"username":"alvin", "amount":311, "currency":"USD"}}' -X POST localhost:3000/users/events
</p>

#### Staff send token to user via email
<p>

    curl localhost:3000/api/token?email=alvinchee98@gmail.com -H "app-secret:Up0F9YrxSDruZlKAxgSiKfdZp7EB8D4XY5vWtbhElHw="
</p>

###### Link to design document: https://docs.google.com/document/d/1z6fQMKwaZNspzKK10ndP_L6Ay_qa2suqLalyVCAioWA/edit?usp=sharing
