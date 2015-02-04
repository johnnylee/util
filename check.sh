#!/bin/bash

go vet &&
errcheck github.com/johnnylee/util
