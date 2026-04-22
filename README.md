# Bureaucracy sort

First and foremost - this is a joke project.

The way it's intended to work is:

1. Client - there is a program that sends requests to sort items via the API. I can't be asked to build this so I'll ask a clanker to do it. I'm using this project just to refresh my golang skills
2. Gateway + Load balancer - I'll implement this from scratch in golang
3. Sorting microservice - implemented in golang, will containerize this and write a k8s deployment around it to explore how autoscaling and other things like that work. This will take in the input from the client and issue swap request events to a message queue
4. Queue - RabbitMQ will work fine here
5. Approver microservice - can't be asked to have a separate client to emulate an approval process, so I'll just have a service that approves the request at a random time
6. PostgreSQL - Outbox for the queue + place to store the response for a given request for later retrieval
7. Notification microservice - this will get polled by the client and receive the requests as it goes

Chaos ensues lmao
