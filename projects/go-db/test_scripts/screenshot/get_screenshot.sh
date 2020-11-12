CT="Content-Type: application/json"
GAME_ID="warcraft3"
USER_ID="user-id-1"
EXT_SYSTEM_ID="exID-2"

URL="http://localhost:1234/api/v1/game/"$GAME_ID"/screenshot?extSystemId="$EXT_SYSTEM_ID"&userId="$USER_ID""

curl \
  -H "$CT" \
  -X GET "$URL"
