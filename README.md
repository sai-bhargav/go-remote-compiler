# Remote Code Compiler
Remote Compiler service can be used to compile run and fetch the output of any code in a sandbox environment.

# About
Remote Compiler service relies on Docker to run the code in a sandbox environment. A new Docker container is spawned on every request and the request payload is mounted on the docker container. After the code is executed the container is removed.

# Supported Languages:
  - Ruby


