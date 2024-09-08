# tree

tree is tool which convert json structure into directory structure. It's simialr to `tree` command in linux. 
For example:

```bash
.
$ tree
├── cmd
│   └── tree
│       ├── cli
│       │   ├── commands.go
│       │   ├── json
│       │   │   └── json.go
│       │   ├── json.go
│       │   └── options
│       │       ├── json.go
│       │       └── root.go
│       └── main.go
├── go.mod
├── go.sum
├── README.md
└── sbomqs-sca.json

5 directories, 10 files
```

## Why tree ?

Basically Json structure has parent-children kind of strucutre. In oder to visualize the relationship b/w parents-child, the tree comes into.

In the world of software and technology, every software has toms of dependencies and each dependecies has further dependencies. For a software supply chain enthusaist, it's a cusrious and required thing to visualize it's project depedencies. For now, this tool support json format.


## How to use it ??

```bash

$ tree json <sbomqs-sca.json>
└── 
    └── github.com/interlynk-io/sbomqs
        ├── github.com/CycloneDX/cyclonedx-go
        │   ├── github.com/bradleyjkemp/cupaloy/v2
        │   ├── github.com/terminalstatic/go-xsd-validate
        │   ├── github.com/xeipuuv/gojsonschema
        │   ├── github.com/xeipuuv/gojsonpointer
        │   └── github.com/xeipuuv/gojsonreference
        ├── github.com/DependencyTrack/client-go
        │   ├── github.com/google/go-cmp
        │   ├── github.com/google/uuid
        │   ├── github.com/jarcoal/httpmock
        │   └── github.com/stretchr/testify
        ├── github.com/Masterminds/semver/v3
        ├── github.com/ProtonMail/go-crypto
        │   ├── github.com/cloudflare/circl
        │   │   ├── golang.org/x/crypto
        │   │   │   ├── golang.org/x/net
        │   │   │   ├── golang.org/x/sys
        │   │   │   ├── golang.org/x/term
        │   │   │   └── golang.org/x/text
        │   │   │       ├── golang.org/x/tools
        │   │   │       │   ├── golang.org/x/net
        │   │   │       │   │   ├── golang.org/x/sys
        │   │   │       │   │   └── golang.org/x/term
        │   │   │       │   ├── golang.org/x/sync
        │   │   │       │   └── golang.org/x/text
        │   │   │       │       └── golang.org/x/tools
        │   │   │       ├── golang.org/x/mod
        │   │   │       │   ├── golang.org/x/crypto
        │   │   │       │   │   ├── golang.org/x/net
        │   │   │       │   │   │   └── golang.org/x/sys
        │   │   │       │   │   ├── golang.org/x/sys
        │   │   │       │   │   ├── golang.org/x/term
        │   │   │       │   │   └── golang.org/x/text
        │   │   │       │   └── golang.org/x/tools
        │   │   │       │       ├── golang.org/x/net
        │   │   │       │       │   ├── golang.org/x/crypto
        │   │   │       │       │   ├── golang.org/x/sys
        │   │   │       │       │   └── golang.org/x/text
        │   │   │       │       ├── golang.org/x/sync
        │   │   │       │       └── golang.org/x/xerrors
        │   │   │       └── golang.org/x/sys
        │   │   └── golang.org/x/sys
        │   └── golang.org/x/crypto
        │       ├── golang.org/x/net
        │       ├── golang.org/x/sys
        │       ├── golang.org/x/term
        │       └── golang.org/x/text
        │           ├── golang.org/x/tools
        │           │   ├── golang.org/x/net
        │           │   │   ├── golang.org/x/term
        │           │   │   └── golang.org/x/text
        │           │   └── golang.org/x/sync
        │           ├── golang.org/x/mod
        │           └── golang.org/x/sys
        ├── github.com/anchore/go-struct-converter
        ├── github.com/cloudflare/circl
        │   ├── github.com/bwesterb/go-ristretto
        │   ├── golang.org/x/crypto
        │   └── golang.org/x/sys
        ├── github.com/common-nighthawk/go-figure
        ├── github.com/davecgh/go-spew
        ├── github.com/github/go-spdx/v2
        │   └── github.com/stretchr/testify
        │       └── github.com/stretchr/objx
        │           └── github.com/stretchr/testify
        │               ├── github.com/davecgh/go-spew
        │               ├── github.com/stretchr/objx
        │               └── gopkg.in/yaml.v3
        ├── github.com/go-git/go-billy/v5
        │   ├── github.com/cyphar/filepath-securejoin
        │   ├── github.com/onsi/gomega
        │   ├── golang.org/x/sys
        │   ├── gopkg.in/check.v1
        │   ├── github.com/kr/pretty
        │   ├── github.com/kr/text
        │   ├── github.com/rogpeppe/go-internal
        │   ├── golang.org/x/net
        │   └── golang.org/x/text
        ├── github.com/google/go-cmp
        ├── github.com/google/go-github/v52
        │   ├── github.com/ProtonMail/go-crypto
        │   ├── golang.org/x/oauth2
        │   ├── github.com/cloudflare/circl
        │   ├── github.com/golang/protobuf
        │   ├── golang.org/x/net
        │   ├── golang.org/x/sys
        │   ├── google.golang.org/appengine
        │   └── google.golang.org/protobuf
        ├── github.com/google/go-querystring
        │   └── github.com/google/go-cmp
        │       └── golang.org/x/xerrors
        ├── github.com/google/uuid
        ├── github.com/inconshreveable/mousetrap
        ├── github.com/mattn/go-runewidth
        │   └── github.com/rivo/uniseg
        ├── github.com/maxbrunsfeld/counterfeiter/v6
        │   ├── github.com/onsi/gomega
        │   ├── github.com/sclevine/spec
        │   ├── golang.org/x/text
        │   ├── golang.org/x/tools
        │   ├── golang.org/x/mod
        │   └── golang.org/x/net
        ├── github.com/olekukonko/tablewriter
        │   └── github.com/mattn/go-runewidth
        ├── github.com/package-url/packageurl-go
        ├── github.com/pkg/errors
        ├── github.com/pmezard/go-difflib
        ├── github.com/rivo/uniseg
        ├── github.com/samber/lo
        │   └── golang.org/x/text
        ├── github.com/spdx/gordf
        ├── github.com/spdx/tools-golang
        │   ├── github.com/anchore/go-struct-converter
        │   └── github.com/spdx/gordf
        ├── github.com/spf13/afero
        │   ├── cloud.google.com/go/storage
        │   ├── github.com/googleapis/google-cloud-go-testing
        │   ├── github.com/pkg/sftp
        │   ├── golang.org/x/crypto
        │   ├── golang.org/x/oauth2
        │   ├── google.golang.org/api
        │   ├── cloud.google.com/go
        │   ├── cloud.google.com/go/compute
        │   ├── cloud.google.com/go/compute/metadata
        │   ├── cloud.google.com/go/iam
        │   ├── github.com/golang/groupcache
        │   ├── github.com/golang/protobuf
        │   ├── github.com/google/s2a-go
        │   ├── github.com/google/uuid
        │   ├── github.com/googleapis/enterprise-certificate-proxy
        │   ├── github.com/googleapis/gax-go/v2
        │   ├── github.com/kr/fs
        │   ├── go.opencensus.io
        │   ├── golang.org/x/net
        │   ├── golang.org/x/sync
        │   ├── golang.org/x/sys
        │   ├── golang.org/x/time
        │   ├── golang.org/x/xerrors
        │   ├── google.golang.org/genproto
        │   ├── google.golang.org/genproto/googleapis/api
        │   ├── google.golang.org/genproto/googleapis/rpc
        │   ├── google.golang.org/grpc
        │   └── google.golang.org/protobuf
        ├── github.com/spf13/cobra
        │   └── github.com/cpuguy83/go-md2man/v2
        │       └── github.com/russross/blackfriday/v2
        ├── github.com/spf13/pflag
        ├── github.com/stretchr/testify
        │   └── github.com/stretchr/objx
        │       └── github.com/stretchr/testify
        │           └── github.com/stretchr/objx
        ├── go
        │   └── toolchain
        ├── go.uber.org/multierr
        │   └── github.com/stretchr/testify
        ├── go.uber.org/zap
        │   ├── github.com/stretchr/testify
        │   ├── go.uber.org/goleak
        │   └── go.uber.org/multierr
        ├── golang.org/x/crypto
        │   ├── golang.org/x/net
        │   ├── golang.org/x/sys
        │   └── golang.org/x/term
        ├── golang.org/x/mod
        │   └── golang.org/x/tools
        ├── golang.org/x/oauth2
        │   └── cloud.google.com/go/compute/metadata
        ├── golang.org/x/sync
        ├── golang.org/x/sys
        ├── golang.org/x/text
        │   ├── golang.org/x/tools
        │   └── golang.org/x/mod
        ├── golang.org/x/tools
        │   ├── github.com/yuin/goldmark
        │   ├── golang.org/x/net
        │   └── golang.org/x/telemetry
        ├── gopkg.in/yaml.v2
        │   └── gopkg.in/check.v1
        ├── gopkg.in/yaml.v3
        ├── gotest.tools
        ├── sigs.k8s.io/release-utils
        │   ├── github.com/blang/semver/v4
        │   ├── github.com/moby/term
        │   ├── github.com/nozzle/throttler
        │   ├── github.com/sirupsen/logrus
        │   ├── github.com/uwu-tools/magex
        │   ├── k8s.io/utils
        │   ├── github.com/Azure/go-ansiterm
        │   ├── github.com/andybalholm/brotli
        │   ├── github.com/dsnet/compress
        │   ├── github.com/golang/snappy
        │   ├── github.com/klauspost/compress
        │   ├── github.com/klauspost/pgzip
        │   ├── github.com/magefile/mage
        │   ├── github.com/mholt/archiver/v3
        │   ├── github.com/nwaples/rardecode
        │   ├── github.com/pierrec/lz4/v4
        │   ├── github.com/ulikunitz/xz
        │   ├── github.com/xi2/xz
        │   ├── golang.org/x/sync
        │   ├── golang.org/x/sys
        │   ├── golang.org/x/text
        │   └── golang.org/x/tools
        ├── sigs.k8s.io/yaml
        └── toolchain
```

## Software Supply Chain Security usecase:

```bash
# generate list of componenets in json format from go modules
$ opensca-cli -path go.mod  -out tree-sca.json

# Now, convert that components into directory structure for easy visualization
$ go run cmd/tree/main.go json --out tree-output.json  tree-sca.json

```

**NOTE**:

- Install opensca-cli tool from [here](https://github.com/XmirrorSecurity/OpenSCA-cli/releases).
