
# Goのインストール

```bash
$ brew install go
```

# ソースコードのクローン

```bash
$ git clone git@github.com:horie1024/AndroidPushTestTool.git
```

# GOPATHのセット

```bash
$ cd AndroidPushTestTool
$ export GOPATH=`pwd` 
```

# ライブラリのインストール

```bash
$ go get github.com/BurntSushi/toml
```

# APIKey、RegistrationIdの設定

`src/main`以下のconfig.tomlに記載します。

```
[setting]
api_key = "YOUR_API_KEY"
regist_id = "YOUR_REGISTRATION_ID"
gcm_server = "GCM_SERVER"
```

# 実行

```bash
$ cd src/main
$ go run main.go -msg="test" -key="KEY" -value="value"
```

# 実行ファイルの生成

```
$ go install
```
実行すると`src/main`以下に実行ファイルが生成されるので、PATHに追加すればコマンドラインから簡単に実行できます。
