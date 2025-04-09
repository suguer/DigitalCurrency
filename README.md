# Digital Currency

这是一个数字货币相关的Go项目。

## 项目结构

```
.
├── api/            # API定义和文档
├── cmd/            # 主程序入口
├── docs/           # 项目文档
├── etc/            # 配置文件
│   └── config.yaml # 主配置文件
├── internal/       # 内部包
│   ├── router/     # API路由处理
│   ├── config/     # 配置管理
│   ├── util/       # 工具类
│   ├── blockchain/ # 区块链相关逻辑
│   └── model/      # 数据模型
├── storage/        # 存储目录
│   ├── logs/       # 日志文件
│   └── data/       # 其他资源文件
├── test/           # 测试文件
├── go.mod          # Go模块文件
└── README.md       # 项目说明文档
```

## 目录说明

- `api/`: 存放API接口定义和相关文档
- `cmd/`: 存放项目的主程序入口
- `docs/`: 项目文档，包含设计文档、使用说明等
- `etc/`: 配置文件目录
- `internal/`: 内部包
  - `router/`: 提供API接口路由管理
  - `config/`: 配置文件管理
  - `util/`: 通用工具类
  - `blockchain/`: 区块链核心逻辑
  - `model/`: 数据模型定义
- `storage/`: 文件存储目录
  - `logs/`: 日志文件
  - `data/`: 其他资源文件
- `test/`: 单元测试和集成测试