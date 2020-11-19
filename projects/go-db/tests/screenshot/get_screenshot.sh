CT="Content-Type: application/json"
GAME_ID="baabf15b-3a05-4592-9935-101637c12d67"
EXT_SYSTEM_ID="custom-ext-system-id"
USER_ID="user-id-1"

URL="http://localhost:1234/api/v1/game/"$GAME_ID"/screenshot?extSystemId="$EXT_SYSTEM_ID"&userId="$USER_ID""

curl \
  -H "$CT" \
  -X GET "$URL"
