<table class="iam-formtable valign-middle">
<tbody>
  <tr>
    <td width="200px">Access Key</td>
    <td class="iam-monofont">
      {[=it.access_key]}
    </td>
  </tr>
  <tr>
    <td>Secret Key</td>
    <td class="iam-monofont">
      {[=it.secret_key]}
    </td>
  </tr>
  <tr>
    <td>Description</td>
    <td>
      {[=it.desc]}
    </td>
  </tr>
  <tr>
    <td>Action</td>
    <td>
      {[~it._actionls :v]}
      {[if (it.action == v.action) {]}{[=v.title]}{[}]}
      {[~]}
    </td>
  </tr>
  <tr>
    <td>Bounds</td>
    <td class="iam-monofont">
      {[~it.bounds :bv]}
      <div style="padding-bottom:5px;">{[=bv.name]}</div>
      {[~]}
    </td>
  </tr>
</tbody>
</table>

