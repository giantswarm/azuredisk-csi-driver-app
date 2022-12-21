{{/* vim: set filetype=mustache: */}}

{{- define "chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/* Expand the name of the chart.*/}}
{{- define "azuredisk.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/* labels for helm resources */}}
{{- define "azuredisk.labels" -}}
app.kubernetes.io/instance: "{{ .Release.Name }}"
app.kubernetes.io/managed-by: "{{ .Release.Service }}"
app.kubernetes.io/name: "{{ template "azuredisk.name" . }}"
app.kubernetes.io/version: "{{ .Chart.AppVersion }}"
helm.sh/chart: {{ include "chart" . | quote }}
{{- end -}}

{{- define "annotations.CRDInstall" -}}
"helm.sh/hook": "pre-install,pre-upgrade"
"helm.sh/hook-delete-policy": "before-hook-creation,hook-succeeded,hook-failed"
{{- end -}}

{{- define "azuredisk.name.crdInstall" -}}
{{- printf "%s-%s" ( include "azuredisk.name" . ) "crd-install" | replace "+" "_" | trimSuffix "-" -}}
{{- end -}}

{{/* Create a label which can be used to select any orphaned crd-install hook resources */}}
{{- define "azuredisk.CRDInstallSelector" -}}
{{- printf "%s" "crd-install-hook" -}}
{{- end -}}
