# 0.0.1 (2020-12-02)
### 功能 / 优化
- **调用 JWT 中间件简化**. 现在使用AuthToken中间件只需传一个jwt conf对象参数就能完成调用.
- **Mysql Mongo 等配置优化**. 现在在 Mysql,Mongo 等配置中增加一个 enable，可手动应用配置.

### Bug 修复
- **重复节点**. [#391](https://github.com/crawlab-team/crawlab/issues/391)
- **重复上传爬虫**. [#603](https://github.com/crawlab-team/crawlab/issues/603)
- **节点第三方模块安装失败导致 节点安装第三方部分无法使用**. [#609](https://github.com/crawlab-team/crawlab/issues/609)
- **离线节点也会创建任务**. [#622](https://github.com/crawlab-team/crawlab/issues/622)