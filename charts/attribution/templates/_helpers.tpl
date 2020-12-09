{{- define "attribution.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "attribution.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "attribution.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "attribution.labels" -}}
helm.sh/chart: {{ include "attribution.chart" . }}
{{ include "attribution.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "attribution.selectorLabels" -}}
app.kubernetes.io/name: {{ template "attribution.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create unified labels for components
*/}}
{{- define "attribution.common.matchLabels" -}}
app: {{ template "attribution.name" . }}
release: {{ .Release.Name }}
{{- end -}}

{{- define "attribution.common.metaLabels" -}}
chart: {{ template "attribution.chart" . }}
heritage: {{ .Release.Service }}
{{- end -}}


{{- define "attribution.ia.labels" -}}
{{ include "attribution.ia.matchLabels" . }}
{{ include "attribution.common.metaLabels" . }}
{{- end -}}

{{- define "attribution.ia.matchLabels" -}}
component: {{ .Values.ia.name | quote }}
{{ include "attribution.common.matchLabels" . }}
{{- end -}}

{{- define "attribution.ia.fullname" -}}
{{ template "attribution.fullname" . }}-{{ .Values.ia.name }}
{{- end -}}

{{- define "attribution.crypto.fullname" -}}
{{ template "attribution.fullname" . }}-{{ .Values.crypto.name }}
{{- end -}}
