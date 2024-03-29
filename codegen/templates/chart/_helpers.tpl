[[/* Custom Helpers */]]
[[- if .HelpersTpl ]]
[[- .HelpersTpl ]]
[[- end ]]

{{/* Below are library functions provided by skv2 */}}
[[/* "skv2.utils."merge is forked from github.solo.io/gloo/blob/master/install/helm/gloo/templates/_helpers.tpl */]]
{{- /*

"skv2.utils.merge" takes an array of three values:
- the top context
- the yaml block that will be merged in (override)
- the name of the base template (source)

note: the source must be a named template (helm partial). This is necessary for the merging logic.

The behaviour is as follows, to align with already existing helm behaviour:
- If no source is found (template is empty), the merged output will be empty
- If no overrides are specified, the source is rendered as is
- If overrides are specified and source is not empty, overrides will be merged in to the source.

Overrides can replace / add to deeply nested dictionaries, but will completely replace lists.
Examples:

┌─────────────────────┬───────────────────────┬────────────────────────┐
│ Source (template)   │       Overrides       │        Result          │
├─────────────────────┼───────────────────────┼────────────────────────┤
│ metadata:           │ metadata:             │ metadata:              │
│   labels:           │   labels:             │   labels:              │
│     app: gloo       │    app: gloo1         │     app: gloo1         │
│     cluster: useast │    author: infra-team │     author: infra-team │
│                     │                       │     cluster: useast    │
├─────────────────────┼───────────────────────┼────────────────────────┤
│ lists:              │ lists:                │ lists:                 │
│   groceries:        │  groceries:           │   groceries:           │
│   - apple           │   - grapes            │   - grapes             │
│   - banana          │                       │                        │
└─────────────────────┴───────────────────────┴────────────────────────┘

skv2.utils.merge is a fork of a helm library chart function (https://github.com/helm/charts/blob/master/incubator/common/templates/_util.tpl).
This includes some optimizations to speed up chart rendering time, and merges in a value (overrides) with a named template, unlike the upstream
version, which merges two named templates.

*/ -}}
{{- define "skv2.utils.merge" -}}
{{- $top := first . -}}
{{- $overrides := (index . 1) -}}
{{- $tpl := fromYaml (include (index . 2) $top) -}}
{{- if or (empty $overrides) (empty $tpl) -}}
{{ include (index . 2) $top }} {{/* render source as is */}}
{{- else -}}
{{- $merged := merge $overrides $tpl -}}
{{- toYaml $merged -}} {{/* render source with overrides as YAML */}}
{{- end -}}
{{- end -}}

[[- range $operator := $.Operators ]]
    [[- if $operator.NamespaceRbac ]]

{{- define "[[ (lower_camel $operator.Name) ]].namespacesForResource" }}
{{- $resourcesToNamespaces := dict }}
{{- range $entry := [[ (opVar $operator) ]].namespacedRbac }}
  {{- range $resource := $entry.resources }}
    {{- $_ := set $resourcesToNamespaces $resource (concat $entry.namespaces (get $resourcesToNamespaces $resource | default list) | mustUniq) }}
  {{- end }}
{{- end }}
{{- get $resourcesToNamespaces  .Resource | join "," }}
{{- end }}
    [[- end ]]
[[- end ]]