# weather-api
This is a weather api in relation to coding exercise

## Project Submission - Weather Service Assignment

### Shortcomings

- Unit Testing: Currently, the project needs more unit tests, especially utilizing the go mock package. The goal is to achieve at least 80% test coverage.
- Dockerization: A Docker file should be added to allow easy deployment and containerization of the service.
- Automation: A Makefile is necessary to simplify the build and test process, making it easier to run, build, and test the application with a single command.
- Integrate swagger

## Steps to set up the service

1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-repo/weather-api.git
   ```

2. **Install dependencies**:
    1. **Clone the repository**:
   ```bash
   go mod tidy
   ```


3. **Cd to cmd/svr**:
   ```bash
   cd cmd/svr
   ```

4. **Build the service**:
   ```bash
   go build .
   ```


5. **Run the service**:
   ```bash
   go run .
   ```




