aws dynamodb put-item \
  --endpoint-url http://localhost:8000 \
  --region ap-northeast-1 \
  --table-name test-table \
  --item '{"Id": {"N": "1"}, "Name": {"S": "to-fmak"}}'
