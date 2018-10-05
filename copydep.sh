echo "APi copydep..."

# Refresh "model"
rm -rf ./vendor/github.com/lagoon-platform/model/*.go
cp ../model/*.go  ./vendor/github.com/lagoon-platform/model/

# Refresh "engine"
# Refresh "engine"
rm -rf ./vendor/github.com/lagoon-platform/engine/*.go
cp ../engine/*.go  ./vendor/github.com/lagoon-platform/engine/

rm -rf ./vendor/github.com/lagoon-platform/engine/ansible/*.go
mkdir ./vendor/github.com/lagoon-platform/engine/ansible/
cp ../engine/ansible/*.go  ./vendor/github.com/lagoon-platform/engine/ansible/

rm -rf ./vendor/github.com/lagoon-platform/engine/component/*.go
mkdir ./vendor/github.com/lagoon-platform/engine/component/
cp ../engine/component/*.go  ./vendor/github.com/lagoon-platform/engine/component/

rm -rf ./vendor/github.com/lagoon-platform/engine/ssh/*.go
mkdir ./vendor/github.com/lagoon-platform/engine/ssh/
cp ../engine/ssh/*.go  ./vendor/github.com/lagoon-platform/engine/ssh/

rm -rf ./vendor/github.com/lagoon-platform/engine/util/*.go
mkdir ./vendor/github.com/lagoon-platform/engine/util/
cp ../engine/util/*.go  ./vendor/github.com/lagoon-platform/engine/util/
