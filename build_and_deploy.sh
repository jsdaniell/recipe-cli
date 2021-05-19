go build -o bin/recipe-cli -v .
git commit -am "update"
git tag -a v1.1.6 -m "Alpha Release" && git push origin v1.1.6
npm publish