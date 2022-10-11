FROM ubuntu:18.04

# Install EDK dependencies
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends build-essential uuid-dev iasl git python python3-distutils nano curl
