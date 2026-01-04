from google.cloud import storage

# Credentials automatic
client = storage.Client(project='PROJECT_ID')
bucket = client.bucket('my-bucket')
blob = bucket.blob('data.txt')
blob.download_to_filename('data.txt')