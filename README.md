# Dynamic Download Experiment

A while back, I built a web application that had a file download facility. The
files being downloaded did not actually exist on the server. They were generated
by the back-end application based use input.

I was not happy with the way I built the link the user interacted with
or with how the download was initiated. This is a little project to figure out
a cleaner way of doing that.

## Running

If you have Docker installed, you can start this application by running `docker-compose up`.
This will build two containers and start them. It will use ports 3000, 8080, and 8081. If you 
open the address [http://localhost:3000](http://localhost:3000) in your browser, it will
open up a page with a couple of links that will initiate downloads. The URIs of the downloads
were dynamically generated.

Hit `^C` to shut the services down.

## Structure

This repository contains three separate projects to build deployable artifacts. 

[exp-download-api](exp-download-api) contains a simple REST service that emulates
generating a file for download.

[exp-download-ui](exp-download-ui) contains a simple web site that has a page with a
button to start the download. There is an Nginx configuration here that will allow
serve up the site and act as a reverse proxy to the API service. It is written assuming
you are using Docker Compose to run the application, but is easily modifiable for
other congigurations.

[exp-download-go](exp-download-go) contains a REST server witten in Go. I'm just beginning
to learn Go, and this was a simple enough project to let me make something.

This structure is overly-complicated for the scope of the test, but I chose to build it this
way to more closely resemble the architecture of the application I was originally
working with.
