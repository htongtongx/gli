# 0.0.1 (2020-12-02)
### 功能 / 优化
- **调用 JWT 中间件简化**. 现在使用AuthToken中间件只需传一个jwt conf对象参数就能完成调用.
- **Mysql Mongo 等配置优化**. 现在在 Mysql,Mongo 等配置中增加一个 enable，可手动应用配置.
- **更多 Git 支持**. 允许用户查看 Git Commits 记录，并 Checkout 到相应 Commit.
- **支持用 Hostname 作为节点注册类型**. 用户可以将 hostname 作为节点的唯一识别号.
- **RPC 支持**. 加入 RPC 支持来更好的管理节点通信.
- **是否在主节点运行开关**. 用户可以决定是否在主节点运行，如果为否，则所有任务将在工作节点上运行.
- **默认禁用教程**.
- **加入相关文档侧边栏**.
- **加载页面优化**.

### Bug 修复
- **重复节点**. [#391](https://github.com/crawlab-team/crawlab/issues/391)
- **重复上传爬虫**. [#603](https://github.com/crawlab-team/crawlab/issues/603)
- **节点第三方模块安装失败导致 节点安装第三方部分无法使用**. [#609](https://github.com/crawlab-team/crawlab/issues/609)
- **离线节点也会创建任务**. [#622](https://github.com/crawlab-team/crawlab/issues/622)