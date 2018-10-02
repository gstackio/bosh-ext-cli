package visualize

const linkProvidersJS string = `
<script src="https://cdnjs.cloudflare.com/ajax/libs/chosen/1.8.7/chosen.jquery.min.js" type="text/javascript" charset="utf-8"></script>

<style type="text/css">
pre {outline: 1px solid #ccc; padding: 5px; margin: 5px; }
.string { color: green; }
.number { color: darkorange; }
.boolean { color: blue; }
.null { color: magenta; }
.key { color: red; }
</style>

<script type="text/javascript">

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
};

function formatDetailsRow ( d ) {
    var str = JSON.stringify(d, undefined, 4);
    return "<pre>"+syntaxHighlight(str)+"</pre>";
}

// Source: https://stackoverflow.com/questions/4810841/how-can-i-pretty-print-json-using-javascript
function syntaxHighlight(json) {
    json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
    return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function (match) {
        var cls = 'number';
        if (/^"/.test(match)) {
            if (/:$/.test(match)) {
                cls = 'key';
            } else {
                cls = 'string';
            }
        } else if (/true|false/.test(match)) {
            cls = 'boolean';
        } else if (/null/.test(match)) {
            cls = 'null';
        }
        return '<span class="' + cls + '">' + match + '</span>';
    });
}

function populateLinkProvidersTable ( providersList ) {
    var providersTable = $('#linkProvidersTable').DataTable( {
        data: providersList,
        columns: [
            {
                "class":          "details-control",
                "orderable":      false,
                "data":           null,
                "defaultContent": ""
            },
            { data: "name", title: "Name" },
            {
                data: {name : "name"},
                title: "Type",
                mRender : function(data, type, full) {
                    return data["link_provider_definition"]["type"];
                }
            },
            { data: "shared", title: "Shared" },
            {
                data: {name : "name", shared: "shared"},
                title: "Provider Type:",
                mRender : function(data, type, full) {
                    return data["owner_object"]["type"];
                }
            },
            {
                data: {name : "name", shared: "shared"},
                title: "Provided by:",
                mRender : function(data, type, full) {
                    if (data["owner_object"]["type"] == "job") {
                        var instanceGroup = data["owner_object"]["info"]["instance_group"];
                        var job = data["owner_object"]["name"];
                        return "<b>Instance Group:</b> " + instanceGroup + "<br />" + "<b>Job:</b> " + job;
                    } else {
                        return data["owner_object"]["name"]
                    }
                }
            }
        ],
        lengthChange: false,
        pageLength: 100,
        ordering: true,
        buttons: [ 'pageLength']
    } );

    var detailRows = [];
    $('#linkProvidersTable tbody').on( 'click', 'tr td.details-control', function () {
        var tr = $(this).closest('tr');
        var row = providersTable.row( tr );
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
        }
    });

    // On each draw, loop over the detailRows array and show any child rows
    providersTable.on( 'draw', function () {
        $.each( detailRows, function ( i, id ) {
            $('#'+id+' providersTable.details-control').trigger( 'click' );
        } );
    } );
}

function fetchAndPopulateLinkProvidersForDeployment ( deploymentName ) {
    var linkProvidersReq = {"command":"curl", "arguments": [{"name": "path", "value": "/link_providers?deployment="+deploymentName}]};
    $.post("/api/command", JSON.stringify(linkProvidersReq))
        .done(function(data) {
            populateLinkProvidersTable(data);
        });
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

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

$(function() {
    fetchAndPopulateDeployments();

    $("#deployment-select").change(function () {
        window.location.href = "/link-providers?deployment=" + $("#deployment-select").find(":selected").text();
    });

    var currentDeploymentName = getUrlParameter('deployment');
    if (currentDeploymentName) {
        fetchAndPopulateLinkProvidersForDeployment(currentDeploymentName);
        $("#title-description").removeClass("invisible").addClass("visible").text(" : " + currentDeploymentName);
    }

    $("#linksProvidersLoadingSpinner").remove();
    $("#linksProvidersContents").removeClass("invisible").addClass("visible");

});
</script>
`
