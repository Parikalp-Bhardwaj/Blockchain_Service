# Blockchain as a Service (BaaS)

A personal project that allows users to create and manage blockchain networks through simple API calls. The system leverages modern technologies like Kubernetes, Ansible, Geth, and Lighthouse to provide scalable and automated blockchain services.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)

## Overview

- **Execution Layer**: Utilizes [Geth](https://geth.ethereum.org/) for transaction execution.
- **Consensus Layer**: Employs [Lighthouse](https://lighthouse-book.sigmaprime.io/) for consensus and validation.
- **Infrastructure**: Deploys on a Kubernetes cluster set up using Ansible scripts.
- **Management**: Provides API endpoints to create, manage, and terminate blockchain networks.
- **Note**: This is a personal project for educational purposes and is not intended for production use **for now**.


## Features

- **Automated Deployment**: Use Ansible to create a Kubernetes cluster effortlessly.
- **API Control**: Manage your blockchain networks via simple API calls.
- **Customizable Setup**: Modify the configurations to suit your personal needs.

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
   git clone https://github.com/Parikalp-Bhardwaj/Blockchain_Service
   cd Blockchain_Service
