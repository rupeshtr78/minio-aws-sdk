# Minio

This repository demonstrates how to interact with MinIO using AWS SDK for Go and AWS CLI, treating MinIO as an S3-compatible object storage service.

## Prerequisites

- Go 1.22.0 or later
- AWS CLI
- MinIO server can be configured using `docker-compose.yaml`

## Setup

1. Clone this repository
2. Install dependencies:



## Setup aws cli with Minio
To set up AWS S3 CLI to work with MinIO, follow these steps:

1. Install AWS CLI if you haven't already.

2. Configure AWS CLI with MinIO credentials:
```
aws configure --profile minio
```
Enter your MinIO access key, secret key, and region (e.g., us-east-1).

3. Enable AWS Signature Version 4:
```
aws configure set default.s3.signature_version s3v4
```

4. When using AWS CLI commands with MinIO,  One option is to include the `--endpoint-url` parameter pointing to your MinIO server:
```
aws --endpoint-url https://your-minio-server:9000 s3 <command>
```

For example, to list buckets:
```
aws --endpoint-url https://your-minio-server:9000 s3 ls
```

To create a bucket:
```
aws --endpoint-url https://your-minio-server:9000 s3 mb s3://mybucket
```

To upload a file:
```
aws --endpoint-url https://your-minio-server:9000 s3 cp file.txt s3://mybucket/
```

### Set Default Profile

```bash
aws configure --profile minio
# Enter access keyid and secret key

aws s3 ls --profile <profile_name>

export AWS_PROFILE=<profile_name>

aws s3 ls
```

### Aws Config File

```bash

~/.aws/config

[default]
region = us-west-1
s3 =
    signature_version = s3v4

[profile minio]
region = us-west-1
services = miniosrv

[services miniosrv]
s3 =
    endpoint_url = your-minio-server:9000

```
