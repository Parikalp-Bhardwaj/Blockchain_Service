# Blockchain as a Service (BaaS)

A personal project that allows users to create and manage blockchain networks through simple API calls. The system leverages modern technologies like Kubernetes, Ansible, Geth, and Lighthouse to provide scalable and automated blockchain services.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
  - [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [License](#license)
- [Contact](#contact)

## Overview

- **Execution Layer**: Utilizes [Geth](https://geth.ethereum.org/) (Go implementation of Ethereum) for executing transactions.
- **Consensus Layer**: Uses [Lighthouse](https://lighthouse-book.sigmaprime.io/) (Rust implementation) for consensus and validation.
- **Infrastructure**: Deploys on a Kubernetes cluster created and managed with Ansible.
- **Management**: Provides API endpoints to create, manage, and terminate blockchain networks.

## Features

- **Easy Deployment**: Spin up new blockchain networks with a single API call.
- **Scalable Architecture**: Kubernetes ensures the system can handle increasing loads.
- **Automated Setup**: Ansible scripts automate the provisioning of infrastructure.
- **Comprehensive Management**: APIs handle all aspects of blockchain lifecycle management.

## Architecture

1. **Kubernetes Cluster**: Orchestrates containerized applications across multiple hosts.
2. **Ansible Automation**: Manages the deployment and configuration of the Kubernetes cluster.
3. **Geth Nodes**: Execute smart contracts and manage the Ethereum state.
4. **Lighthouse Validators**: Ensure consensus across the network.
5. **API Layer**: Facilitates interaction with the system using RESTful API calls.

## Getting Started

### Prerequisites

- **Ansible** installed on your local machine.
- Access to a **Kubernetes** cluster or the ability to create one.
- **Docker** installed on all nodes in the Kubernetes cluster.
- **Python 3.6+** for running the API server.

### Installation

1. **Clone the Repository**

   ```bash
   https://github.com/Parikalp-Bhardwaj/Blockchain_Service
   cd Blockchain_Service
