
# yamlのチェック
openapi-generator-cli validate -i openapi.yml

# goコードの生成
openapi-generator-cli generate -i api/openapi.yml -g go-server -o ./server
