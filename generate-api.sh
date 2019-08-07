# generate cloudsearch api client
# echo "$GOPATH/bin/google-api-go-generator" -cache=false -api_json_file="$GOPATH/src/github.com/nilleb/google-api-go-client/cloudsearch-v1beta1.json" -install -api_pkg_base github.com/nilleb/google-api-go-client -output "$GOPATH/src/github.com/nilleb/google-api-go-client/cloudsearch/v1beta1/cloudsearch-gen.go"
# "$GOPATH/bin/google-api-go-generator" -cache=false -api_json_file="$GOPATH/src/github.com/nilleb/google-api-go-client/cloudsearch-v1beta1.json" -install -api_pkg_base github.com/nilleb/google-api-go-client -output "$GOPATH/src/github.com/nilleb/google-api-go-client/cloudsearch/v1beta1/cloudsearch-gen.go"

echo "$GOPATH/bin/google-api-go-generator" -cache=false -api_json_file="$GOPATH/src/github.com/nilleb/google-api-go-client/cloudsearch-v1.json" -install -api_pkg_base github.com/nilleb/google-api-go-client -output "$GOPATH/src/github.com/nilleb/google-api-go-client/cloudsearch/v1/cloudsearch-gen.go"
"$GOPATH/bin/google-api-go-generator" -cache=false -api_json_file="$GOPATH/src/github.com/nilleb/google-api-go-client/cloudsearch-v1.json" -install -api_pkg_base github.com/nilleb/google-api-go-client -output "$GOPATH/src/github.com/nilleb/google-api-go-client/cloudsearch/v1/cloudsearch-gen.go"

# generate enterprise grade groups api client v1alpha
# echo "$GOPATH/bin/google-api-go-generator" -cache=false -api_json_file="$GOPATH/src/github.com/nilleb/google-api-go-client/enterprisegradegroups-v1alpha.json" -install -api_pkg_base github.com/nilleb/google-api-go-client -output "$GOPATH/src/github.com/nilleb/google-api-go-client/groups/v1alpha/cloudidentity-gen.go"
# "$GOPATH/bin/google-api-go-generator" -cache=false -api_json_file="$GOPATH/src/github.com/nilleb/google-api-go-client/enterprisegradegroups-v1alpha.json" -install -api_pkg_base github.com/nilleb/google-api-go-client -output "$GOPATH/src/github.com/nilleb/google-api-go-client/groups/v1alpha/cloudidentity-gen.go"

# generate enterprise grade groups api client v1beta1
# echo "$GOPATH/bin/google-api-go-generator" -cache=false -api_json_file="$GOPATH/src/github.com/nilleb/google-api-go-client/enterprisegradegroups-v1beta1.json" -install -api_pkg_base github.com/nilleb/google-api-go-client -output "$GOPATH/src/github.com/nilleb/google-api-go-client/cloudidentity/v1beta1/cloudidentity-gen.go"
# "$GOPATH/bin/google-api-go-generator" -cache=false -api_json_file="$GOPATH/src/github.com/nilleb/google-api-go-client/enterprisegradegroups-v1beta1.json" -install -api_pkg_base github.com/nilleb/google-api-go-client -output "$GOPATH/src/github.com/nilleb/google-api-go-client/cloudidentity/v1beta1/cloudidentity-gen.go"

# generate enterprise grade groups api client v1
echo "$GOPATH/bin/google-api-go-generator" -cache=false -api_json_file="$GOPATH/src/github.com/nilleb/google-api-go-client/enterprisegradegroups-v1.json" -install -api_pkg_base github.com/nilleb/google-api-go-client -output "$GOPATH/src/github.com/nilleb/google-api-go-client/cloudidentity/v1/cloudidentity-gen.go"
"$GOPATH/bin/google-api-go-generator" -cache=false -api_json_file="$GOPATH/src/github.com/nilleb/google-api-go-client/enterprisegradegroups-v1.json" -install -api_pkg_base github.com/nilleb/google-api-go-client -output "$GOPATH/src/github.com/nilleb/google-api-go-client/cloudidentity/v1/cloudidentity-gen.go"
