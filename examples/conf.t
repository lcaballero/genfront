{{ range $struct := .data }}<table>{{ range $field, $doc := $struct.FieldDoc }}
<tr>
  <th>{{ $field }}</th>
  <td>{{ $doc }}</td>
</tr>{{ end }}
</table>
{{ end }}
