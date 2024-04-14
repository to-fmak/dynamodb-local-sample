import boto3

dynamodb = boto3.client("dynamodb", endpoint_url="http://localhost:8000")

dynamodb.put_item(TableName="test-table", Item={
    "Id": {"N": "1"},
    "Name": {"S": "to-fmak"}
})
