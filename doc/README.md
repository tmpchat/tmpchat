# tmpchat

A fleeting moment of Engineer.

https://tmpchat.com/

## Feature

- Anonymity
- For Engineer

## Chat topics

- Programming Language Questions
- About career change
- Question regarding IT technology and other
- Light talk
  - How to study, English, About Remote Work

## Component

- Top page
- Search Room
  - Sort Room
    - Popular
    - New arrivals
- Char Room
  - Anonymity
  - Code Syntax
    - Markdown ?
    - WYSIWYG ?
  - Auto Room deletion
  - Room name is not URL Path
    - `room/XXXXXXXXXXXXXXXXXXX`
  - Auto Room creation
    - From [GitHub Trending](https://github.com/trending)

## Architecture

### Service

- Frontend
  - HTML, CSS, JavaScript(Elm)
- Backend(Golang)
  - WebSocket(Room)
  - API(Search, Create Room, Delete Room)
- Storage(DB)
  - For Room Chat
  - Room list
- Job
  - Room deletion trigger
  - Auto Room creation

```
Create Room: Frontend -> API
Chat: Frontend -> WebSocket
Delete Room: Job      -> API
```

### Infrastructure

- Compute
  - GKE
- Storage
  - Spanner
- Dockerfile
  - Docker Hub

### Git Repository

- GitHub
- Monorepo
  - doc
  - frontend
  - api
  - websocket
  - job
  - terraoform
  - kubernetes
- Branch strategy
  - git-flow

### Communication

- Slack

### Issue Management

- GitHub Issue
