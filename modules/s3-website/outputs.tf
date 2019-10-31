output "url" {
  value = "http://${aws_s3_bucket.example.website_endpoint}"
}

output "arn" {
  value = aws_s3_bucket.example.arn
}