# Microservices Test Project

This repository contains the configuration and setup for a microservices test project. This environment is designed to show the development, types of communications, and deployment of microservices-based applications.

## Overview

The microservices test development environment is built using [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/), providing an isolated and consistent environment for testing and development. It includes configurations for various tools commonly used in microservices development, such as service discovery, message brokers, databases, and logging.

## Setup

### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [GO](https://go.dev/doc/install)
- [MongoDB](https://www.mongodb.com/)
- [Postgres](https://www.postgresql.org/)

### Getting Started

1. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/microservices-test-environment.git

2. Navigate into the folder project inside the cloned repository:

   ```bash
   cd /project/

3. Using make file start all microservices

   ```bash
   make up_build

4. Using make file again start the front end

   ```bash
   make start
