ARCHIVE_LOW1="archives=@/Users/a.mogilevskiy/work/5gen/clean-city/archives/low1.zip"
ARCHIVE_LOW2="archives=@/Users/a.mogilevskiy/work/5gen/clean-city/archives/low2.zip"
ARCHIVE_MEDIUM="archives=@/Users/a.mogilevskiy/work/5gen/clean-city/archives/medium.zip"

GAME_ID1="29ceff68-1bc0-4099-b0f2-2205f8b1f12d" # past
GAME_ID2="bd255325-e7d1-44bd-8f76-1ff4796e71a2" # future

URL="http://localhost:1234/api/v1/game/${GAME_ID2}"

#  -F "$ARCHIVE_LOW1" \
#  -F "$ARCHIVE_MEDIUM" \

curl \
  -F "$ARCHIVE_LOW1" \
  -X PUT "$URL"
