aws dynamodb get-item \
  --endpoint-url http://localhost:8000 \
  --region ap-northeast-1 \
  --table-name test-table \
  --key '{"Id": {"N": "1"}}'
