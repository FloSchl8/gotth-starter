# Go HTMX Tailwind Starter Template

A starter template project that integrates Go, Templ, HTMX, and Tailwind CSS to kickstart your web development.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)

## Introduction
This project provides a foundational setup for building web applications using Go, Templ, HTMX, and Tailwind CSS. It aims to streamline the development process by offering a pre-configured environment with essential tools and libraries.

## Features
- **Go**: Powerful backend language.
- **Templ**: Efficient templating engine.
- **HTMX**: Modern HTML-centric front-end library.
- **Tailwind CSS**: Utility-first CSS framework.

## Requirements
- Go (version X.X.X or higher)
- Node.js (for Tailwind CSS)
- Templ (templating engine)
- HTMX (HTML-centric front-end library)

## Installation
To get started, clone the repository and navigate into the project directory:
```sh
git clone https://github.com/FloSchl8/gotth-starter.git
cd gotth-starter
```

## Usage
### Development
For development, use the following command to start the server with hot-reloading:

```sh
air
templ --watch --proxy=http://localhost:<port>
```

### Production
For production builds, use:

```sh
go build -tags=prod -o your_project_name
```

## Contributing
Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (git checkout -b feature-branch).
3. Make your changes.
4. Commit your changes (git commit -m 'Add new feature').
5. Push to the branch (git push origin feature-branch).
6. Open a Pull Request.
