[![CircleCI](https://circleci.com/gh/ProjectReferral/Get-me-in/tree/master.svg?style=svg&circle-token=632ab80f9b534a6dab955b1f27f267b00b700ac4)](https://circleci.com/gh/ProjectReferral/Get-me-in/tree/master)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)[![Version](https://badge.fury.io/gh/tterb%2FHyde.svg)](https://badge.fury.io/gh/tterb%2FHyde)
[![GitHub last commit](https://img.shields.io/github/last-commit/google/skia.svg?style=flat)]()
[![Open Source](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://opensource.org/)

# Referral Marketing System

## Overview
TESTING 123
TBC

## Technical Overview

All the current microservices are built using GO with a mix of request-driven and event-driven architecture. For event-driven, we using RabbitMQ to broadcast messages.

#### Current services:
- [Under dev]Authentication Service(auth-service) - handles the lifecycle of JSON Web Tokens(JWT).
- [Under dev]Account Service(account-service) - handles all the CRUD operations to do with users.
- [Under dev]Marketing Service(marketing-service) - handles all the CRUD operations to do with job adverts.

#### New services under development/analysis:
- Customer Service(customer-service) - handles email confirmations, reset passwords and any other communications between the consumer and producer.
- Messaging Service(msg-service) - handles instant messaging between users.

#### Front-end:
Front end will be designed using React and Redux.

#### Deploy process:
We using CircleCI to manage our build pipeline. To manage our infrastructure, we are using Docker and AWS.

#### Future work:
- Service orchestration using K8s or Docker Swarms
- Setup ELB(Elastic Load Balancer)
- Terraform to manage AWS infrastructure
- Setup Grafana


## High level overview
![High-level Architecture](Q-split-6.png)


