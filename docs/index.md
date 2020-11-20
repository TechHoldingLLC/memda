---
layout: default
---

## Installation

```sh
brew install TechHoldingLLC/tap/memda
```
or 
```sh
brew tap TechHoldingLLC/tap
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