# ottoTTS_server
使用 http 请求获取电棍音频

## 使用方法
- api：`/speak`
- 方法：`POST`
- 请求体：
```json
{
  "message": "<内容>"
}
```
- 返回：`audio/wav`

## 配置文件说明
```toml
expression_override = true
Debug = false
Port = 8080
```
`expression_override` 为 true 时，会进行短语匹配，否则只匹配单个字符。`Debug` 为 true 时，会在控制台打印调试信息。`port` 为端口号。

## 部署

### 拉取仓库
```bash
git clone https://github.com/gujial/ottoTTS_server.git
```

### 运行
```bash
cd ottoTTS_server/
go mod tidy
go run main.go
```

### 构建二进制文件
```bash
make
```
构建的二进制文件在 `./build` 目录下，并且资源文件会自动复制到该文件夹中。

## 修改资源文件
参考[ottoTTS](https://github.com/gujial/ottoTTS?tab=readme-ov-file#%E4%BF%AE%E6%94%B9%E8%B5%84%E6%BA%90%E6%96%87%E4%BB%B6)

## 使用到的开源库
- [ottoTTS](https://github.com/gujial/ottoTTS_server)