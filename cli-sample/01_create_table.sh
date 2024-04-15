aws dynamodb \
  --endpoint-url http://localhost:8000 \
  --region ap-northeast-1 \
    create-table \
  --table-name test-table \
  --attribute-definitions \
    AttributeName=Id,AttributeType=N \
  --key-schema \
    AttributeName=Id,KeyType=HASH \
  --billing-mode PAY_PER_REQUEST
