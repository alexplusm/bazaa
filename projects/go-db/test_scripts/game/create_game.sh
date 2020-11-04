CT="Content-Type: application/json"
URL="http://localhost:1234/api/v1/game"

curl -H "$CT" -d @create_game.json -X POST "$URL"