# 使い方
動作公開版はHerokuを用いております．URLは以下に記します．
https://helloworldrest.herokuapp.com/

以下のコマンドにより動作確認を行いました．
curl -XGET -H 'Content-Type:application/json' https://helloworldrest.herokuapp.com/

# ファイルの構成

- docker-compose.yml
- Dockerfile
- heroku.yml
- main.go