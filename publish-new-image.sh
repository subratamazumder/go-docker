##############################################
# usage e.g.; $publish-new-image.sh 5.0
##############################################
export IMAGE_TAG=$1
#replace .go source with SERVICE_VERSION="$IMAGE_TAG"
#TBC
# build executableß
GO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o eprescription-reg-service
# build & tag image
echo "Building image with tag as"+$IMAGE_TAG
docker build -t dockersubrata/eprescription-reg-service-image:$IMAGE_TAG .
# publish image to docker hub
echo "Publishing image with tag as"+$IMAGE_TAG
docker push dockersubrata/eprescription-reg-service-image:$IMAGE_TAG
