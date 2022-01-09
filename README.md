# ddd-sample

DDD(ドメイン駆動設計)のサンプル

## ユースケース図


<img src="https://user-images.githubusercontent.com/61341861/148669646-139845b4-fe21-4fb9-b267-c2a074c3737a.png" width="800" />"

## ドメインモデル図/オブジェクト図

<img src="https://user-images.githubusercontent.com/61341861/148669648-c7cedd70-e3db-4991-9351-e4e8521edeac.png" width="800" />"

## 使用技術
- Go 1.17
- オニオンアーキテクチャ
- google/wire ※Dependency Injection
- oklog/ulid ※ID生成
- gomock

#### クリーンアーキテクチャではなくオニオンアーキテクチャを採用している理由

- オニオンアーキテクチャのほうが登場する概念が少なくシンプルであるため
- クリーンアーキテクチャの中心は「Entities層」となっているが、これが表すのはDDDのエンティティとは異なるため(DDDのエンティティはドメイン層の一部)

## ディレクトリ構成

```shell
% tree -d
.
├── domain ※ドメイン層
│   ├── shared
│   │   ├── error
│   │   └── page
│   ├── task
│   │   └── mock_task
│   └── user
│       └── mock_user
├── infra ※インフラ層
│   ├── in_memory
│   │   ├── task
│   │   ├── transaction
│   │   └── user
│   └── rdb
│       └── transaction
├── presentation ※プレゼンテーション層
│   ├── shared
│   │   └── error
│   ├── task
│   └── user
└── usecase ※ユースケース層
    ├── shared
    │   ├── error
    │   └── transaction
    ├── task
    └── user
```

## 参考資料
- [ドメイン駆動設計 モデリング/実装ガイド](https://little-hands.booth.pm/items/1835632)
- [ドメイン駆動設計 サンプルコード&FAQ](https://little-hands.booth.pm/items/3363104)
- [ドメイン駆動設計入門 ボトムアップでわかる! ドメイン駆動設計の基本](https://www.amazon.co.jp/%E3%83%89%E3%83%A1%E3%82%A4%E3%83%B3%E9%A7%86%E5%8B%95%E8%A8%AD%E8%A8%88%E5%85%A5%E9%96%80-%E3%83%9C%E3%83%88%E3%83%A0%E3%82%A2%E3%83%83%E3%83%97%E3%81%A7%E3%82%8F%E3%81%8B%E3%82%8B-%E3%83%89%E3%83%A1%E3%82%A4%E3%83%B3%E9%A7%86%E5%8B%95%E8%A8%AD%E8%A8%88%E3%81%AE%E5%9F%BA%E6%9C%AC-%E6%88%90%E7%80%AC-%E5%85%81%E5%AE%A3/dp/479815072X/ref=tmm_pap_swatch_0?_encoding=UTF8&qid=&sr=)
