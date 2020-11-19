CT="Content-Type: application/json"

PAST="@game_from_past.json"
FUTURE="@game_from_future.json"
FAR_FUTURE="@game_from_far_future.json"

## extSystemId
# 2285017c-f527-40d2-aead-b23cce159947 | generatedID
# custom-ext-system-id | customID

URL="http://localhost:8080/api/v1/game"

#   -d "$PAST" \
#   -d "$FUTURE" \
#   -d "$FAR_FUTURE" \

curl \
  -H "$CT" \
  -d "$FUTURE" \
  -X POST "$URL"