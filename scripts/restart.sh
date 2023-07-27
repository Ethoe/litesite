#!/bin/bash

# Step 1: Build React app
cd web/app
npm install
npm run build

# Step 2: Move back to the project root directory
cd ../../

# Step 3: Run docker-compose
docker-compose up -d --force-recreate --build