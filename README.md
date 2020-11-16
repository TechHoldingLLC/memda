# MEMDA

AWS Lambda Memory
Check what lambda run out of memory

You can either scan all you lambdas or provide a function name to check if you ran out of memory.
Memda relies on a ~/.aws/credentials and / or ~/.aws/config

## Installation

```sh
brew install ctrlaltdev/tap/memda
```
or 
```sh
brew tap ctrlaltdev/tap
brew install memda
```

## Usage

Usage:
```sh
memda --region us-east-1 --all
```

Using a profile
```sh
memda --region us-east-1 --profile prod --all
```

For a specific lambda
```sh
memda --region us-east-1 --lambda my-function-name
```

And always:
```sh
memda -h
```
