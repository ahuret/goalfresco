docker run -u $UID:$GID --rm \
    -v $PWD:/local openapitools/openapi-generator-cli generate \
    -i /local/api-spec.yaml \
    -g go \
    -o /local/openapi

grep -r "github.com/GIT_USER_ID/GIT_REPO_ID" openapi -l | xargs sed -i 's/GIT_USER_ID\/GIT_REPO_ID/ahuret\/goalfresco\/openapi/g'
grep -r '"time"' openapi -l | xargs sed -i 's/"time"/"github.com\/ahuret\/goalfresco\/time"/g'
