# peng-api

## development

## 出现下载模块找不到执行 
```sh
go get github.com/gofrs/uuid
```

## swagger
### swagger使用文档
https://github.com/swaggo/swag/blob/master/README_zh-CN.md
### 生成 swagger 文档
```bash
swag init --parseDependency --parseInternal --parseDepth 1
```
### swagger地址
http://localhost:8080/swagger/index.html