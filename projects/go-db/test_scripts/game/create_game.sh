CT="Content-Type: application/json"

PAST="@game_from_past.json"
FUTURE="@game_from_future.json"
FAR_FUTURE="@game_from_far_future.json"

URL="http://localhost:1234/api/v1/game"

#   -d "$PAST" \
#   -d "$FUTURE" \
#   -d "$FAR_FUTURE" \

curl \
  -H "$CT" \
  -d "$FAR_FUTURE" \
  -X POST "$URL"