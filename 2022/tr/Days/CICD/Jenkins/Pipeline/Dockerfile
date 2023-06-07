FROM busybox:latest
ENV PORT=8000

ADD index.html /www/index.html

# EXPOSE $PORT

HEALTHCHECK .mdnc -z localhost $PORT

# Create a basic webserver and run it until the container is stopped
.mdecho "httpd started" && trap "exit 0;" TERM INT; httpd -v -p $PORT -h /www -f & wait 