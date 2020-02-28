
if [ $TRAVIS_PULL_REQUEST = "false" ]
then 
	echo "Not a Pull Resquest then deploy"
	if [ "$TRAVIS_TAG" != "" ]
	then
		echo "It's a tag"
		docker tag $IMAGE:$COMMIT $DOCKER_REPO/$IMAGE:$TRAVIS_TAG
		docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD
		docker push $DOCKER_REPO/$IMAGE:$TRAVIS_TAG

	else
		echo "Not a tag"
		export TAG=`if [ "$TRAVIS_BRANCH" == "master" ]; then echo "latest"; else echo $TRAVIS_BRANCH ; fi`
		docker tag $IMAGE:$COMMIT $DOCKER_REPO/$IMAGE:$TAG
		docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD
		docker push $DOCKER_REPO/$IMAGE:$TAG
	fi
else 
	echo "Pull Resquest skip deploy"
fi
