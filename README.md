# Go Alfresco

## What is go alfresco ?

Go Alfresco is a library that allow to interact with alfresco REST API.
This library use *openapi generator cli* to generate golang files that allow to interact with alfresco based on the *alfresco openapi spec* (replacement for goalfresco).

## Code

All code is generated using [openapi generator](https://openapi-generator.tech/) based on [alfresco spec](https://api-explorer.alfresco.com/api-explorer/definitions/alfresco-core.yaml) which is hosted in the repository as [api-spec.yaml](https://github.com/ahuret/goalfresco/blob/master/api-spec.yaml).
Note that alfresco specification use swagger 2.0 format which is outdated.

### Time package

You can see there's a time package, we use it as replacement to time.Time package in all generated code since time.Unmarshal return error while parsing alfresco time (format issue).
Goal is to remove this folder at some point, it will be probably possible when alfresco upgrade its specification to upstream one.

### Documentation

You can see detailed documentation for every api and model in [openapi](https://github.com/ahuret/goalfresco/tree/master/openapi#readme).

## Contribute

### Update generated code

```bash
# remove old generated code
rm -rf openapi/*

# generate new code based on api-spec.yaml
docker run -u $UID:$GID --rm \
    -v $PWD:/local openapitools/openapi-generator-cli generate \
    -i /local/api-spec.yaml \
    -g go \
    -o /local/openapi

# change the go module path
grep -r "github.com/GIT_USER_ID/GIT_REPO_ID" openapi -l | xargs sed -i 's/GIT_USER_ID\/GIT_REPO_ID/ahuret\/goalfresco\/openapi/g'

# force to use custom time package
grep -r '"time"' openapi -l | xargs sed -i 's/"time"/"github.com\/ahuret\/goalfresco\/time"/g'
```
