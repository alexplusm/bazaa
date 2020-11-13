CT="Content-Type: application/json"

WITH_ID="@ext_system_with_id.json"
WITHOUT_ID="@ext_system_without_id.json"

URL="http://localhost:1234/api/v1/ext-system"

#  -d "$WITH_ID" \
#  -d "$WITHOUT_ID" \

curl \
  -H "$CT" \
  -d "$WITHOUT_ID" \
  -X POST "$URL"