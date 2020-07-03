# 概要

tmpchat - _〜 エンジニアの儚いひととき 〜_

https://tmpchat.com/

## 特徴

- 匿名
- エンジニア向け

## チャット内容

- プログラミングなど IT 技術関連についての質問
- 転職相談
- 雑談
  - 勉強方法, 英語, リモートワークなど

## 機能

- トップページ
- ルーム検索
  - Sort
    - 人気順
    - 新着順
- チャットルーム
  - 全員匿名
  - 記法
    - Markdown ?
    - WYSIWYG ?
  - 自動 Room 削除
  - Room 名 != URL PATH
    - `room/XXXXXXXXXXXXXXXXXXX`
  - 自動 Room 生成
    - [GitHub Trending](https://github.com/trending) から取得

## 実装方法 / 技術選定

### Service

- Frontend
  - HTML, CSS, JavaScript(Elm)
- Backend
  - WebSocket(Room)
  - API(Search, Room 生成, Room 削除)
- Storage(DB)
  - Chat用
  - Room一覧
- Job
  - Room削除のトリガー
  - 自動Room生成

### Architecture

```
Create Room: Frontend -> API
Chat:        Frontend -> WebSocket
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
- ブランチ戦略
  - git-flow

### Communication

- Slack

### Issue Management

- GitHub Issue
