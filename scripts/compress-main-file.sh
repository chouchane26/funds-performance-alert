#!/bin/bash

mkdir ../deployment && cd ../deployment

cp ../main . && cp -r ../src/resources .

ls -la

zip -r main.zip .