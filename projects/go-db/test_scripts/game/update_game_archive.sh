ARCHIVE_LOW1="archives=@/Users/a.mogilevskiy/work/5gen/clean-city/archives/low1.zip"
ARCHIVE_LOW2="archives=@/Users/a.mogilevskiy/work/5gen/clean-city/archives/low2.zip"
ARCHIVE_MEDIUM="archives=@/Users/a.mogilevskiy/work/5gen/clean-city/archives/medium.zip"

GAME_ID1="29ceff68-1bc0-4099-b0f2-2205f8b1f12d" # past
GAME_ID2="52d9cfed-1d03-46f5-9921-fca5cb9c116e" # future

URL="http://localhost:1234/api/v1/game/${GAME_ID2}"

#  -F "$ARCHIVE_LOW1" \
#  -F "$ARCHIVE_MEDIUM" \

curl \
  -F "$ARCHIVE_MEDIUM" \
  -X PUT "$URL"
