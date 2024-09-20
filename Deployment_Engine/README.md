# Blockchain Management Platform

A personal project that provides a platform to create, manage, and monitor blockchain networks through API endpoints. The system consists of a Go backend server and a React frontend, orchestrated using a Makefile for easy setup and management. This guide includes instructions tailored for Visual Studio Code (VS Code) users.

## Table of Contents

- [Introduction](#introduction)
- [Project Structure](#project-structure)
- [Setup in VS Code](#setup-in-vs-code)
  - [Prerequisites](#prerequisites)
  - [Extensions](#extensions)
  - [Workspace Settings](#workspace-settings)
- [Makefile Commands](#makefile-commands)
- [API Endpoints](#api-endpoints)
- [Getting Started](#getting-started)
  - [Installation](#installation)
- [Usage](#usage)
- [License](#license)
- [Contact](#contact)

## Introduction

This project aims to simplify the process of deploying and managing blockchain networks. By utilizing a Go backend and a React frontend, users can interact with the system through a user-friendly interface or API calls. The Makefile included in the project provides convenient commands for installing dependencies, running servers, and cleaning up the environment.

## Project Structure

```plaintext
.
├── Backend/          # Go backend server code
├── Frontend/         # React frontend application
├── hosts.ini         # Ansible inventory file (auto-generated)
├── Makefile          # Makefile with predefined commands
└── README.md         # Project documentation


