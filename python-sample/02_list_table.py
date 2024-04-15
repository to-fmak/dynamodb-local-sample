import boto3

dynamodb = boto3.client(
    "dynamodb", region_name="ap-northeast-1", endpoint_url="http://localhost:8000"
)

print(dynamodb.list_tables()["TableNames"])
