import boto3

dynamodb = boto3.client("dynamodb", endpoint_url="http://localhost:8000")

dynamodb.create_table(
    AttributeDefinitions=[
        {
            "AttributeName": "Id",
            "AttributeType": "N"
        },
    ],
    TableName="test-table",
    KeySchema=[
        {
            "AttributeName": "Id",
            "KeyType": "HASH"
        },
    ],
    BillingMode="PAY_PER_REQUEST"
)
