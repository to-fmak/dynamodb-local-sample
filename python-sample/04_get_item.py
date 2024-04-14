import boto3

dynamodb = boto3.client("dynamodb", endpoint_url="http://localhost:8000")

item = dynamodb.get_item(
    TableName="test-table", Key={"Id": {"N": "1"}}
)
print(item["Item"])
