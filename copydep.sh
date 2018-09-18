echo "APi copydep..."

# Refresh "model"
rm -rf ./vendor/github.com/lagoon-platform/model/*.go
cp ../model/*.go  ./vendor/github.com/lagoon-platform/model/

# Refresh "engine"
rm -rf ./vendor/github.com/lagoon-platform/engine/*.go
cp ../engine/*.go  ./vendor/github.com/lagoon-platform/engine/

rm -rf ./vendor/github.com/lagoon-platform/engine/ssh/*.go
mkdir ./vendor/github.com/lagoon-platform/engine/ssh/
cp ../engine/ssh/*.go  ./vendor/github.com/lagoon-platform/engine/ssh/
