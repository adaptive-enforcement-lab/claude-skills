from google.cloud import secretmanager

client = secretmanager.SecretManagerServiceClient()
secret_name = f"projects/PROJECT_ID/secrets/api-key/versions/latest"
response = client.access_secret_version(request={"name": secret_name})
api_key = response.payload.data.decode('UTF-8')