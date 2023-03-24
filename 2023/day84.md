# Writing an API - What is an API? 

The acronym API stands for “application programming interface”. What does this really mean though? It’s a way of
controlling an application programmatically. So when you use a website that displays some data to you (like Twitter)
there will be an action taken by the interface to get data or send data to the application (the twitter backend in this
example) - this is done programmatically in the background by code running in the user interface.

In the example given above we looked at an example of a public API, however the vast majority of APIs are private, one
request to the public twitter API will likely cause a cascade of interactions between programs in the backend. These
could be to save the tweet text into a datastore, to update the number of likes or views a tweet has or to take an image
that has been uploaded and resize it for a better viewing experience.

We build programs with APIs that other people can call so that we can expose program logic to other developers, teams
and our customers or suppliers. They are a predefined way of sharing information. For example, we can define an API
using [openapi specification](https://swagger.io/resources/open-api/) which is used
for [RESTful](https://en.wikipedia.org/wiki/Representational_state_transfer) API design. This api specification forms a
contract that we can fulfil. For example, If you make an API request to me and pass a specific set of content, such as a
date range, I will respond with a specific set of data. Therefore you can reliably expect to receive data of a certain
type when calling my API.

We are going to build up an example set of applications that communicate using an API for this section of the learning
journey to illustrate the topics and give you a hands-on look at how things can be done.

Design:
2 programs that communicate bi-directionally, every minute or so one application will request a random string from the
other, once one is received it will store this number in a database for future use

The Random Number generator will generate a random string when requested and save this into a database, the application
will then ask the first program for confirmation that it received the string, and store this information against the
string in the database

The applications will be called:
generator
requestor

This may sound like a silly example but it allows us to quickly look into the various tasks involved with building,
deploying, monitoring and owning a service that runs in production. There are bi-directional failure modes as each
application needs something from the other to complete successfully and things we can monitor such as API call rates -
We can see if one application stops running.

We need to now decide what our API Interfaces should look like. We have 2 API calls that will be used to communicate
here. Firstly, the `requestor` will call the `generator` and ask for a string. This is likely going to be an API call
without any additional content other than making a request for a string. Secondly, the `generator` will start to ask
the `requestor` for confirmation that it received and stored the string, in this case we need to pass a parameter to
the `requestor` which will be the string we are interested in knowing about.

The `generator` will use a URL path of `/new` to serve up random strings
The `requestor` is going to use URL paths to receive string information from the `generator` to check the status, so we
will setup a URL of `/strings/<STRING>` where <STRING> is the string of interest.


## Building the API
There is a folder on the Github repository  under the 2023 section called `day2-ops-code` and we will be using this 
folder to store our code for this, and future, section of the learning journey.

We are using Golang's built in HTTP server to serve up our endpoints and asynchronous goroutines to handle the 
checks. Every 60 seconds we will look into the generators database and get all the strings which we dont have 
conformation for and then calling the requesters endpoint to check if the string is there.