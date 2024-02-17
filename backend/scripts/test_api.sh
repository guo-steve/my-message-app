#!/bin/bash -e

email=${1:-demo@example.com}
password=${2:-1234}

set -x

curl -sS -f -XPOST localhost:8080/v1/register -d '{"email":"'$email'","password":"'$password'","full_name":"John Doe"}'
echo "Registered successfully"

out=$(curl -sS -f -XPOST localhost:8080/v1/login -d '{"email":"'$email'","password":"'$password'"}')
token=$(echo $out | jq -r .token)
echo "Logged in successfully! Token: $token"

out=$(curl -sS -f -XGET localhost:8080/v1/messages -H "Authorization: Bearer $token")
if [[ $(echo $out | jq 'length') -ne 0 ]]; then
  echo "Messages are not empty"
  exit 1
fi
echo "Messages: $out"

curl -sS -f -XPOST localhost:8080/v1/messages -H "Authorization: Bearer $token" -d '{"text":"Hello, World!"}'

out=$(curl -sS -f -XGET localhost:8080/v1/messages -H "Authorization: Bearer $token")
if [[ $(echo $out | jq 'length') -ne 1 ]]; then
  echo "Messages are empty"
  exit 1
fi
echo "Messages: $out"

curl -sS -f -XPOST localhost:8080/v1/logout -H "Authorization: Bearer $token"
echo "Logged out successfully"

out=$(curl -sS -XGET localhost:8080/v1/messages)
echo $out

set +x
