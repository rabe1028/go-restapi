# 使い方
動作公開版はHerokuを用いております．URLは以下に記します．

https://resttesthelloworld.herokuapp.com/

以下のコマンドにより動作確認を行いました．
curl -XGET -H 'Content-Type:application/json' https://resttesthelloworld.herokuapp.com/

# ファイルの構成

- Dockerfile
    Goの環境を作成するためのDockerfile
- heroku.yml
    herokuでDockerを動かすためのFile
- main.go
    mainプログラム