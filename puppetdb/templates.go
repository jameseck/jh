package puppetdb

const V2Query = `
[ "and"
, [ "or"
,   [ "~", "certname", "{{ .Host }}" ]
  ]
, [ "or"

{{- range .Ff.Filters }}
,   [ "=", "name", "{{ .Name }}" ]
{{- end }}
  ]
, [ "and"

{{- range .Ff.Filters }}
,   [ "in", "certname"
,     [ "extract", "certname", [ "select-facts"
,       [ "and"
,         [ "=", "name", "{{ .Name }}" ]
,         [ "{{ .Operator }}", "value", "{{ .Value }}" ]
        ]
      ]]
    ]
{{- end }}
  ]
]
`
