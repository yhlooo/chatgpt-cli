# chatgpt-cli

通过命令行与 [ChatGPT](https://chat.openai.com/chat) 对话

## 使用

1. 安装 chatgpt-cli
   - 从源码构建

     ```shell
     go build github.com/keybrl/chatgpt-cli
     ```

   - 或从 [Releases](https://github.com/keybrl/chatgpt-cli/releases) 下载二进制

2. 从 OpenAI 获取 [API Key](https://platform.openai.com/account/api-keys) （形如 `sk-...` ）
3. 运行

   ```shell
   ./chatgpt-cli chat --secret-key '<Your-OpenAI-API-Key>'
   ```

   （需将命令中 `<Your-OpenAI-API-Key>` 替换为真实的 OpenAI API Key ）

更多使用方法参考：

```shell
./chatgpt-cli chat --help
```
