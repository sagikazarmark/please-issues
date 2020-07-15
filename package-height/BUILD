genrule(
    name = 'go.mod',
    outs = ['go.mod'],
    cmd = [
      """cat > $OUTS << EOF
module go.go.go.com/go

go 1.12
EOF
      """
    ],
    labels = [
      'link:plz-out/'
    ]
)
