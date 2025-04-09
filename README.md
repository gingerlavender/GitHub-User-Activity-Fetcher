# GitHub User Activity Fetcher 🚀

[![Go Version](https://img.shields.io/badge/Go-1.19%2B-blue)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![GitHub Stars](https://img.shields.io/github/stars/gingerlavender/GitHub-User-Activity-Fetcher?style=social)](https://github.com/gingerlavender/GitHub-User-Activity-Fetcher)

## 🔍Overview
A tool for analyzing GitHub user activity. Collects data on events (commits, pull requests, repository creation, etc.), displays statistics on them, or visualizes them in an interactive chart.


## 📦 Installing

### Steps:
1. **Clone repository**:
   ```bash
   git clone https://github.com/gingerlavender/GitHub-User-Activity-Fetcher.git
   cd GitHub-User-Activity-Fetcher
   ```
2.  **Install dependencies**:
    ```bash
    go mod tidy
    ```
3.  **Build binary file**:
    ```bash
    go build
    ```

## 🚀 Usage
**Command format:**

    gh-api [username] [flags]
  | Flags |Description |
  |--|--|
  | -d, -w, -m, -y| Collects events occured over given period: 1, 7, 31 or 365 day(s) respectively |
  |\-\-period|Specifies any other period  (requires days as int)|
  |\-\-eventType |Shows only particular event (requires event as string, specified as its full proper name, e.g. "PushEvent" or "IssuesEvent")|
  |\-\-plot  |Defines whether draw interactive chart of activity or not (if not, just omit this flag)|
  |\-\-token (-t)|Defines your GitHub API authorization token|
