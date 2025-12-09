#!/bin/bash

MODULES=("authentication" "storage" "tenant" "user")

echo "Start swagger generated."

for module in "${MODULES[@]}"; do
  echo "Cleaning docs for $module..."
  rm -rf services/$module/docs/*

  echo "Generating swagger for $module..."
  cd services/$module
  swag init -g ./module.go -d ./ -o ./docs --instanceName $module --parseDependency
  cd ../..
done

echo "All swagger docs generated successfully."
