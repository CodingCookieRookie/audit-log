# audit-log

Audit log for user to get and post any event
 
1. Start docker daemon using cmd `sudo dockerd` in a terminal
2. Git clone repository into audit-log folder
3. Open a new terminal in the audit-log folder directory and run cmd `./start.sh`
4. Curl endpoints depending on usage

<h3> Curl commands
curl localhost:3000/users/events\?event_type=exchange\&event_timestamp_start=2023-05-02_17:04:04\&event_timestamp_end=2023-06-01_15:04:05\&gmt=*8 -H "email:alvinchee98@gmail.com" -H "token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFsdmluY2hlZTk4QGdtYWlsLmNvbSIsImV4cCI6MTY4ODExNDQyNH0.sXUYUdGRvIKOqtSEKWJkE9Q7CI2JK4R_0ZubciG8ZfE"