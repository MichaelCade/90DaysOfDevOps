# Set the base image to use
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port that the application will run on
EXPOSE 8080

# Define the command that will run when the container starts
CMD ["/app/main"]
