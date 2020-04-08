# Weathering App

## Notes

### Date
March 31, 2020

### Location of deployed application
Taken down: http://ac66454e95b994d5ea078ab72cef6dd7-1725874665.us-east-1.elb.amazonaws.com/docs
Taken down: http://ac66454e95b994d5ea078ab72cef6dd7-1725874665.us-east-1.elb.amazonaws.com/v1/weather?city=vancouver

### Time spent
4 hours

### Assumptions made
- Users only cares about temperature and condition (sunny, cloudy, etc.)
- Users can only query by a valid city.
- Users requested the data in Fahrenheit
- We have unlimited usage of openweathermap.org's API (I signed up for a free token, not sure how long it would last)
- Endpoint isn't fully RESTful standard. It should be `/weathers` (But I made it `/weather` to match the description)

### Shortcuts/Compromises made
- City validation relies on OpenWeather

### Stretch goals attempted
- I used openweathermap.org's API to query city-based weather in real-time. Their API is straightforward so it takes no time to integrate
- I deployed it live to AWS. I have an EKS cluster available, so deploying to Kubernetes using Helm takes no time at all
- I added a docs section just so that consumers of this API onboard themselves

### Instructions to run assignment locally
With Docker and Docker Compose installed (Docker >= `v19.03.5`. docker-compose >= `v1.24.1` recommended):
- `docker-compose up`
- `curl localhost:3031/v1/docs`

Without Docker (Make >= `3.81`. Go `1.13.4` recommended):
- `cd server`
- `go get`
- `make dev`
- `curl localhost:3031/v1/docs`

### What did you not include in your solution that you want us to know about?
I was experimenting with Devops, where I created a Kubernetes Cluster to deploy my projects.
I deployed it using EKS on AWS.

### Other information about your submission that you feel it's important that we know if applicable.
### Your feedback on this technical challenge
The task at hand is simple, it encourages developers to expand on this app in a way of their choosing/creativity
