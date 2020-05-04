source .set
echo "Global Variables are set"
echo "Building Require Images..."
docker-compose build
echo "Images are Built..."
echo "Starting Containers..."
docker-compose up -d
echo "Containers are Started..."
source .unset
echo "Removing Global Variables"
