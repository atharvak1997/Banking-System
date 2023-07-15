# Banking-System
Banking system using Go, Docker, Postgres and gRPC

<div align="center">
<h1 align="center">
<img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="100" />
<br>Banking-System
</h1>
<h3>◦ </h3>
<h3>◦ Developed with the software and tools listed below.</h3>

<p align="center">
<img src="https://img.shields.io/badge/Go-00ADD8.svg?style&logo=Go&logoColor=white" alt="Go" />
<img src="https://img.shields.io/badge/Markdown-000000.svg?style&logo=Markdown&logoColor=white" alt="Markdown" />
</p>
<img src="https://img.shields.io/github/languages/top/atharvak1997/Banking-System?style&color=5D6D7E" alt="GitHub top language" />
<img src="https://img.shields.io/github/languages/code-size/atharvak1997/Banking-System?style&color=5D6D7E" alt="GitHub code size in bytes" />
<img src="https://img.shields.io/github/commit-activity/m/atharvak1997/Banking-System?style&color=5D6D7E" alt="GitHub commit activity" />
<img src="https://img.shields.io/github/license/atharvak1997/Banking-System?style&color=5D6D7E" alt="GitHub license" />
</div>

---

## 📒 Table of Contents
- [📒 Table of Contents](#-table-of-contents)
- [📍 Overview](#-overview)
- [⚙️ Features](#-features)
- [🧩 Modules](#modules)
- [🚀 Getting Started](#-getting-started)
- [🤝 Contributing](#-contributing)
- [📄 License](#-license)
- [👏 Acknowledgments](#-acknowledgments)

---


## 📍 Overview

1. You will gain a thorough understanding of database design, including generating reliable and consistent code for interacting with the database using transactions. You will also learn about database isolation levels and how to correctly utilize them in a production environment. In addition to databases, you will learn how to utilize Docker for local development, effectively use Git for code management, and leverage GitHub Actions for automated unit testing.

2. You will acquire the skills to construct a collection of RESTful HTTP APIs using Gin, a widely adopted Golang framework for creating web services. This encompasses various aspects such as configuring the application, implementing robust unit testing with a simulated database, managing errors, verifying user identities, and ensuring API security through JWT and PASETO access tokens.

3. You will gain knowledge on building your application using Docker and deploying it to a production Kubernetes cluster on AWS. The lectures provide comprehensive instructions with a detailed, step-by-step guide. You will learn essential topics such as building a minimal Docker image, setting up an AWS account, creating a production database, securely managing production secrets, establishing a Kubernetes cluster using EKS, automating image building and deployment to the EKS cluster using GitHub Actions, acquiring a domain name and routing traffic to your service, ensuring secure connections with HTTPS, and automatically renewing TLS certificates from Let's Encrypt.

4. We will explore various advanced backend topics. These include the management of user sessions, the creation of gRPC APIs, the utilization of gRPC gateway to handle both gRPC and HTTP requests simultaneously, the integration of Swagger documentation into the backend service, the ability to update records partially by leveraging optional parameters, and the development of structured logger HTTP middlewares and gRPC interceptors.

5. Asynchronous processing in Golang by implementing background workers and utilizing Redis as a message queue. Additionally, we will discuss the importance of gracefully shutting down the server to safeguard processing resources. It's worth noting that this section is continually being updated, and we will regularly upload new videos covering additional topics such as sending emails, gracefully shutting down servers, CORS (Cross-Origin Resource Sharing), bulk inserts, and more. Therefore, I encourage you to revisit this platform periodically to explore and access the latest content.
---

## ⚙️ Features


1. Establishing and overseeing bank accounts.
2. Documenting all modifications in the balances of individual accounts.
3. Executing a transaction to transfer funds between two accounts.

---

## 🧩 Modules

<details closed><summary>Root</summary>

| File                                                                                                                             | Summary                                                                                                                                                                                                                                                   |
| ---                                                                                                                              | ---                                                                                                                                                                                                                                                       |
| [go.mod](https://github.com/atharvak1997/Banking-System/blob/main/go.mod)                                                        | The code snippet is a Go module that manages a simple bank application. It requires external packages for database connectivity, testing, and handling YAML files.                                                                                        |
| [000001_init_schema.down.sql](https://github.com/atharvak1997/Banking-System/blob/main/db\migration\000001_init_schema.down.sql) | The code snippet is responsible for dropping three database tables: entries, transfers, and accounts.                                                                                                                                                     |
| [000001_init_schema.up.sql](https://github.com/atharvak1997/Banking-System/blob/main/db\migration\000001_init_schema.up.sql)     | This code snippet creates three tables: "accounts", "entries", and "transfers" with their respective columns. It establishes foreign key relationships between the tables, adds indexes on relevant columns, and includes comments on column constraints. |
| [000001_init_schema.up.sql](https://github.com/atharvak1997/Banking-System/blob/main/db\migration\000001_init_schema.up.sql)     | This code snippet creates three tables: "accounts", "entries", and "transfers" with their respective columns. It establishes foreign key relationships between the tables, adds indexes on relevant columns, and includes comments on column constraints. |

| [account.sql](https://github.com/atharvak1997/Banking-System/blob/main/db\query\account.sql)     | This code snippet has functions like CreateAccount, GetAccount, UpdateAccount, UpdateAccountBalance, DeleteAccount|

| [entry.sql](https://github.com/atharvak1997/Banking-System/blob/main/db\query\entry.sql)     | This code snippet is responsible for creating enrtries of numbers in respective accounts|

</details>

---

## 🚀 Getting Started

### ✔️ Prerequisites

Before you begin, ensure that you have the following prerequisites installed:
> - `ℹ️ Golang`
> - `ℹ️ SQLC`
> - `ℹ️ SQL`
> - `ℹ️ Docker`
> - `ℹ️ Kubernetes`

### 📦 Installation

1. Clone the Banking-System repository:
```sh
git clone https://github.com/atharvak1997/Banking-System
```

2. Change to the project directory:
```sh
cd Banking-System
```

3. Install the dependencies:


### 🎮 Using Banking-System


### 🧪 Running Tests

---


## 🤝 Contributing

Contributions are always welcome! Please follow these steps:
1. Fork the project repository. This creates a copy of the project on your account that you can modify without affecting the original project.
2. Clone the forked repository to your local machine using a Git client like Git or GitHub Desktop.
3. Create a new branch with a descriptive name (e.g., `new-feature-branch` or `bugfix-issue-123`).
```sh
git checkout -b new-feature-branch
```
4. Make changes to the project's codebase.
5. Commit your changes to your local branch with a clear commit message that explains the changes you've made.
```sh
git commit -m 'Implemented new feature.'
```
6. Push your changes to your forked repository on GitHub using the following command
```sh
git push origin new-feature-branch
```
7. Create a new pull request to the original project repository. In the pull request, describe the changes you've made and why they're necessary.
The project maintainers will review your changes and provide feedback or merge them into the main branch.

---

## 📄 License

This project is licensed under the `ℹ️  INSERT-LICENSE-TYPE` License. See the [LICENSE](https://docs.github.com/en/communities/setting-up-your-project-for-healthy-contributions/adding-a-license-to-a-repository) file for additional info.

---

## 👏 Acknowledgments

> - `ℹ️  List any resources, contributors, inspiration, etc.`

---
