import boto3

def lambda_handler(event, context):
    path = event['path']
    s3_client = boto3.client('s3')
    bucket_name = 'iapolinario-bucket'

    if path == "/health":
        return {
            'statusCode': 200,
            'headers': {'Content-Type': 'text/plain'},
            'body': "Healthy"
        }
    elif path.startswith("/api/"):
        print("Entered the /api")
        file_name = path[5:] + '.json' 
        print("File name:", file_name)
        try:
            obj = s3_client.get_object(Bucket=bucket_name, Key=file_name)
            data = obj["Body"].read().decode('utf-8')
            print("Response:", data)
            return {
                'statusCode': 200,
                'headers': {'Content-Type': 'application/json'},
                'body': str(data).upper()
            }
        except Exception as e:
            print("Error", str(e))
            return {
                'statusCode': 500,
                'headers': {'Content-Type': 'text/plain'}, # Added headers
                'body': str(e)
            }

    return {
        'statusCode': 404,
        'headers': {'Content-Type': 'text/plain'}, # Added headers
        'body': "Not Found"
    }
