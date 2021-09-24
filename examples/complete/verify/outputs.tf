# We can use the verify stack to expose outputs only needed for testing purposes
output "bucket_id" {
  value = data.aws_s3_bucket.example.id
}
