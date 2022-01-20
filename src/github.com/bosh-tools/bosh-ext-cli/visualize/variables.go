package visualize

const variablesTemplate string = `

<div id="variablesLoadingSpinner">
 <div class="row">
  <div class="col-md-6 offset-md-3" style="text-align: center;">
   <i class="fa fa-circle-o-notch fa-spin" style="font-size: 7rem"></i>
  </div>
 </div>
</div>

<div class="row">
	<div class="col-md-4">
	 <div class="form-group">
	    <select class="form-control" id="deployment-select" data-placeholder="Choose a deployment">
	    </select>
	 </div>
	</div>
</div>
<table id="variablesTable" class="configs-table table table-striped table-hover table-bordered" width="100%">
</table>
`

const variablesJS string = `
<script src="https://cdnjs.cloudflare.com/ajax/libs/chosen/1.8.7/chosen.jquery.min.js" type="text/javascript" charset="utf-8"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/ace/1.3.3/ace.js" type="text/javascript" charset="utf-8"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/ace/1.3.3/theme-monokai.js" type="text/javascript" charset="utf-8"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/ace/1.3.3/mode-javascript.js" type="text/javascript" charset="utf-8"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/js-yaml/3.12.0/js-yaml.min.js" type="text/javascript" charset="utf-8"></script>

<script type="text/javascript">
function formatDetailsRow ( d ) {
    return '<div style="position: relative; height: 400px;" id="variable-editor-'+d.id+'">Loading, please wait... </br>(If this takes too long, make sure you have configured your credhub env variables correctly)</div>';
}

function fetchAndPopulateDeployments () {
    var linkProvidersReq = {"command":"curl", "arguments": [{"name": "path", "value": "/deployments"}]};
    $.post("/api/command", JSON.stringify(linkProvidersReq))
        .done(function(deploymentsList) {
            var deploymentListDropDown = $("#deployment-select");
            deploymentListDropDown.find('option').remove();
            deploymentListDropDown.append($("<option />"));
            $.each(deploymentsList, function() {
                deploymentListDropDown.append($("<option />").val(this.name).text(this.name));

            });
            deploymentListDropDown.chosen();
        });
}

function getUrlParameter(sParam) {
    var sPageURL = decodeURIComponent(window.location.search.substring(1)),
        sURLVariables = sPageURL.split('&'),
        sParameterName,
        i;

    for (i = 0; i < sURLVariables.length; i++) {
        sParameterName = sURLVariables[i].split('=');

        if (sParameterName[0] === sParam) {
            return sParameterName[1] === undefined ? true : sParameterName[1];
        }
    }
}

$(function() {
    fetchAndPopulateDeployments();

    $("#deployment-select").change(function () {
        window.location.href = "/variables?deployment=" + $("#deployment-select").find(":selected").text();
    });

    var currentDeploymentName = getUrlParameter('deployment');
    if (currentDeploymentName) {
        var variablesReq = {"command":"curl","arguments": [{"name": "path", "value": "/deployments/"+currentDeploymentName+"/variables"}]};
        $.post("/api/command", JSON.stringify(variablesReq))
            .done(function(data) {
                $("#variablesLoadingSpinner").remove();
                var table = $('#variablesTable').DataTable( {
                    data: data,
                    columns: [
                        {
                            "class":          "details-control",
                            "orderable":      false,
                            "data":           null,
                            "defaultContent": ""
                        },
                        { data: "name", title: "Name" },
                        { data: "id", title: "ID" }
                    ],
                    lengthChange: false,
                    paging: false,
                    ordering: true,
                    buttons: [ 'pageLength']
                } );

                var detailRows = [];

                $('#variablesTable tbody').on( 'click', 'tr td.details-control', function () {
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

                        var credhubReq = {"command":"credhub","arguments": [{"name": "id", "value": row.data().id}, {"name":"output-json"}]};
                        $.post("/api/command", JSON.stringify(credhubReq)).done(function (data) {
                            $("#variable-editor-"+row.data().id).text(jsyaml.safeDump(data));
                            var editor = ace.edit("variable-editor-"+row.data().id);
                            editor.setTheme("ace/theme/terminal");
                            editor.session.setMode("ace/mode/yaml");
                            editor.setFontSize(15);
                            editor.setReadOnly(true);
                            console.log(data);
                        });
                    }
                } );

                // On each draw, loop over the detailRows array and show any child rows
                table.on( 'draw', function () {
                    $.each( detailRows, function ( i, id ) {
                        $('#'+id+' table.details-control').trigger( 'click' );
                    } );
                } );
            });
    } else {
        $("#variablesLoadingSpinner").remove();
    }
});
  </script>
`
