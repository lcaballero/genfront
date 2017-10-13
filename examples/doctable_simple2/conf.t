{{- with .data -}}
<table>
  {{range $field, $doc := .FieldDoc -}}
  <tr>
    <th>{{ $field }}</th>
    <td>{{ $doc }}</td>
  </tr>
  {{- end }}
</table>
{{- end }}
