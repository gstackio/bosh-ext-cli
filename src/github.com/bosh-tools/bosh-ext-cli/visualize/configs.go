package visualize

const configsTemplate string = `

<div id="configsLoadingSpinner">
 <div class="row">
  <div class="col-md-6 offset-md-3" style="text-align: center;">
   <i class="fa fa-circle-o-notch fa-spin" style="font-size: 7rem"></i>
  </div>
 </div>
</div>

<table id="configsTable" class="configs-table table table-striped table-hover table-bordered" width="100%">
</table>
`

const configsJS string = `
<script src="https://cdnjs.cloudflare.com/ajax/libs/ace/1.3.3/ace.js" type="text/javascript" charset="utf-8"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/ace/1.3.3/theme-monokai.js" type="text/javascript" charset="utf-8"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/ace/1.3.3/mode-javascript.js" type="text/javascript" charset="utf-8"></script>

<script type="text/javascript">
function formatDetailsRow ( d ) {
    return '<div style="position: relative; height: 1200px;" id="config-editor-'+d.id+'"">'+d.content+'</div>';
}

$(function() {
    var configsReq = {"command":"curl","arguments": [{"name": "path", "value": "/configs"}]};
    $.post("/api/command", JSON.stringify(configsReq))
        .done(function(data) {
            $("#configsLoadingSpinner").remove();
            var table = $('#configsTable').DataTable( {
                data: data,
                columns: [
                    {
                        "class":          "details-control",
                        "orderable":      false,
                        "data":           null,
                        "defaultContent": ""
                    },
                    { data: "id", title: "ID" },
                    { data: "type", title: "Type" },
                    { data: "name", title: "Name" },
                    { data: "team", title: "Team" },
                    { data: "created_at", title: "Created at" }
                ],
                lengthChange: false,
                paging: false,
                ordering: true,
                buttons: [ 'pageLength']
            } );

            var detailRows = [];

            $('#configsTable tbody').on( 'click', 'tr td.details-control', function () {
                var tr = $(this).closest('tr');
                var row = table.row( tr );
                var idx = $.inArray( tr.attr('id'), detailRows );

                if ( row.child.isShown() ) {
                    tr.removeClass( 'details' );
                    row.child.hide();

                    // Remove from the 'open' array
                    detailRows.splice( idx, 1 );
                }
                else {
                    tr.addClass( 'details' );
                    row.child( formatDetailsRow( row.data() ) ).show();

                    // Add to the 'open' array
                    if ( idx === -1 ) {
                        detailRows.push( tr.attr('id') );
                    }

                    var editor = ace.edit("config-editor-"+row.data().id);
                    editor.setTheme("ace/theme/terminal");
                    editor.session.setMode("ace/mode/yaml");
                    editor.setFontSize(15);
                    editor.setReadOnly(true);
                }
            } );

            // On each draw, loop over the detailRows array and show any child rows
            table.on( 'draw', function () {
                $.each( detailRows, function ( i, id ) {
                    $('#'+id+' table.details-control').trigger( 'click' );
                } );
            } );
        });
});
  </script>
`
