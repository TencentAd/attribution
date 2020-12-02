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
app.kubernetes.io/name: {{ include "attribution.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "attribution.serviceAccountName" -}}
{{- if .Values.serviceAccount.enabled }}
{{- default (include "attribution.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
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
{{- if .Values.ia.fullnameOverride -}}
{{- .Values.ia.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- printf "%s-%s" .Release.Name .Values.ia.name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s-%s" .Release.Name $name .Values.ia.name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Define the attribution.namespace template if set with forceNamespace or .Release.Namespace is set
*/}}
{{- define "attribution.namespace" -}}
namespace: {{ .Release.Namespace }}
{{- end -}}

