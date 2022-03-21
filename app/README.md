# for Mac M1 (tutorial/updating-our-app/)
docker build -t getting-started --platform linux/arm64 .
docker run --platform linux/arm64  -dp 3000:3000 getting-started

## in play-with-docker
docker run --platform linux/arm64 -dp 3000:3000 tsepakkwan1990/getting-started