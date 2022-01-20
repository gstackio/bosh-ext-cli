package visualize

const linkConsumersDetailedTemplate string = `
<div id="linksConsumersDetailedLoadingSpinner">
 <div class="row">
  <div class="col-md-6 offset-md-3" style="text-align: center;">
   <i class="fa fa-circle-o-notch fa-spin" style="font-size: 7rem"></i>
  </div>
 </div>
</div>

<div id="linksConsumersDetailedContents" class="invisible">
<div class="row">
	<div class="col-md-4">
		<div class="card text-white bg-primary mb-3">
		  <div class="card-body">
			<h5 class="card-title">About this page:</h5>
			<p class="card-text">
              This page will list all the Link Consumers of the specified deployment. 
              It will also display the links established for these consumers, if any.
              Additionally, it will list the Link Providers for each consumer.   
            </p>
			 <div class="form-group" style="color: black;">
				<select class="form-control" id="deployment-select" data-placeholder="Choose a deployment">
				</select>
			 </div>
		  </div>
		</div>
	</div>
	<div class="col-md-4">
		<ul class="list-group">
		  <li class="list-group-item active">Statistics</li>
		  <li class="list-group-item">Current Deployment: <span id="stat-dep-name" class="badge badge-success badge-pill" style="font-size: 0.85rem"></span></li>
		  <li class="list-group-item">Total # of Consumers: <span id="stat-total" class="badge badge-success badge-pill" style="font-size: 0.85rem"></span></li>
		  <li class="list-group-item">Fulfilled consumers: <span id="stat-fulfilled" class="badge badge-success badge-pill" style="font-size: 0.85rem"></span></li>
		  <li class="list-group-item">External consumers: <span id="stat-external" class="badge badge-success badge-pill" style="font-size: 0.85rem"></span></li>
		  <li class="list-group-item">Consumers whose owner is a Job: <span id="stat-job" class="badge badge-success badge-pill" style="font-size: 0.85rem"></span></li>
		  <li class="list-group-item">Consumers whose owner is a Variable: <span id="stat-var" class="badge badge-success badge-pill" style="font-size: 0.85rem"></span></li>
		</ul>
	</div>
	<div class="col-md-4">
	</div>
</div>
<div class="row">
    <table id="linkConsumersDetailedTable" class="link-consumers-detailed-table table table-hover table-bordered" width="100%">
	</table>
</div>
</div>
`

const linkConsumersDetailedJS string = `
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
}

function formatDetailsRow ( d ) {
    var consumerStr = JSON.stringify(d["consumer"], undefined, 4);
    var linkStr = d["link"] ? JSON.stringify(d["link"], undefined, 4) : "Not Established";
    var providerStr = d["provider"] ? JSON.stringify(d["provider"], undefined, 4) : "Not Resolved";
    var consumerBlock = "<div class=\"card\"><div class=\"card-header text-white bg-primary mb-3\">Consumer</div><div class=\"card-body\"><pre>" + syntaxHighlight(consumerStr) +"</pre></div></div>";
    var linkBlock = "<div class=\"card\"><div class=\"card-header text-white bg-success mb-3\">Link</div><div class=\"card-body\"><pre>" + syntaxHighlight(linkStr) +"</pre></div></div>";
    var providerBlock = "<div class=\"card\"><div class=\"card-header text-white bg-info mb-3\">Provider</div><div class=\"card-body\"><pre>" + syntaxHighlight(providerStr) +"</pre></div></div>";
    return "<div class=\"row\"><div class=\"col-sm-4\">"+consumerBlock+"</div><div class=\"col-sm-4\">"+linkBlock+"</div><div class=\"col-sm-4\">"+providerBlock+"</div></div>";
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

function createdRowCallback( row, data, dataIndex){
    var consumerType = data["consumer"]["owner_object"]["type"];
    if( consumerType === "variable" ){
        $(row).addClass('list-group-item-success');
    } else if(consumerType === "external" ) {
        $(row).addClass('list-group-item-danger');
    }
}

function populateLinkConsumersDetailedTable ( consumersList ) {
    var consumersTable = $('#linkConsumersDetailedTable').DataTable( {
        data: consumersList,
        columns: [
            {
                "class":          "details-control",
                "orderable":      false,
                "data":           null,
                "defaultContent": ""
            },
            { data: "consumer.id", title: "ID" },
            {
                data: {name : "consumer.name"},
                title: "Name/Type",
                mRender : function(data, type, full) {
                    return data["consumer"]["name"] + " / " +data["consumer"]["link_consumer_definition"]["type"];
                }
            },
            {
                data: {name : "consumer.name"},
                title: "Consumer Owner:",
                mRender : function(data, type, full) {
                    if (data["consumer"]["owner_object"]["type"] === "job") {
                        var instanceGroup = data["consumer"]["owner_object"]["info"]["instance_group"];
                        var job = data["consumer"]["owner_object"]["name"];
                        return "<b>Job: </b>\"" + job + "\",<br/>" + "<b>Instance Group: </b>\"" + instanceGroup + "\"";
                    } else {
                        var consumerOwnerName = data["consumer"]["owner_object"]["name"];
                        var consumerOwnerType = data["consumer"]["owner_object"]["type"];
                        return "<b>"+consumerOwnerType+": </b>\"" + consumerOwnerName;
                    }
                }
            },
            {
                data: {name : "consumer.name"},
                title: "Fulfilled:",
                mRender : function(data, type, full) {
                    return data["link"] ? "Yes" : "No";
                }
            },
            {
                data: {name : "consumer.name"},
                title: "Fulfilled By:",
                mRender : function(data, type, full) {
                    if (data["provider"]) {
                        var deployment = data["provider"]["deployment"];
                        if (data["provider"]["owner_object"]["type"] === "job") {
                            var instanceGroup = data["provider"]["owner_object"]["info"]["instance_group"];
                            var job = data["provider"]["owner_object"]["name"];
                            return  "<b>Job: </b> \"" + job + "\",<br/><b>In Instance Group: </b>\"" + instanceGroup + "\",<br/><b>In Deployment: </b>\"" + deployment + "\"";
                        } else {
                            var providerName = data["provider"]["owner_object"]["name"];
                            var providerType = data["provider"]["owner_object"]["type"];
                            return "<b>"+providerType+": </b>\"" + providerName + "\",<br/><b>In Deployment: </b>\"" + deployment + "\"";
                        }
                    } else {
                        return "No link established";
                    }
                }
            }
        ],
        lengthChange: false,
        pageLength: 100,
        ordering: true,
        buttons: [ 'pageLength'],
        createdRow: createdRowCallback
    } );

    var detailRows = [];
    $('#linkConsumersDetailedTable tbody').on( 'click', 'tr td.details-control', function () {
        var tr = $(this).closest('tr');
        var row = consumersTable.row( tr );
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
    consumersTable.on( 'draw', function () {
        $.each( detailRows, function ( i, id ) {
            $('#'+id+' consumersTable.details-control').trigger( 'click' );
        } );
    } );
}

function populateDeploymentsDropDown (deploymentsList) {
    var deploymentListDropDown = $("#deployment-select");
    deploymentListDropDown.find('option').remove();
    deploymentListDropDown.append($("<option />"));
    $.each(deploymentsList, function() {
        deploymentListDropDown.append($("<option />").val(this.name).text(this.name));
    });
    deploymentListDropDown.chosen();
}

function populateStatistics ( consumersList, currentDeploymentName ) {
    var totalNumberOfConsumers = consumersList.length;
    var totalNumberOfFulfilledConsumers = 0;
    var totalNumberOfExternalConsumers = 0;
    var totalNumberOfConsumersOwnedByAJob = 0;
    var totalNumberOfConsumersOwnedByAVariable = 0;
    $(consumersList).each(function () {
        if (this["link"]) {
            totalNumberOfFulfilledConsumers++;
        }

        if (this["consumer"]["owner_object"]["type"] === "external") {
            totalNumberOfExternalConsumers++;
        }

        if (this["consumer"]["owner_object"]["type"] === "job") {
            totalNumberOfConsumersOwnedByAJob++;
        }

        if (this["consumer"]["owner_object"]["type"] === "variable") {
            totalNumberOfConsumersOwnedByAVariable++;
        }
    });

    $("#stat-dep-name").text(currentDeploymentName);
    $("#stat-total").text(totalNumberOfConsumers);
    $("#stat-fulfilled").text(totalNumberOfFulfilledConsumers);
    $("#stat-external").text(totalNumberOfExternalConsumers);
    $("#stat-job").text(totalNumberOfConsumersOwnedByAJob);
    $("#stat-var").text(totalNumberOfConsumersOwnedByAVariable);
}

$(function() {
    var linkConsumersReq = {"command":"curl", "arguments": [{"name": "path", "value": "/deployments"}]};
    $.post("/api/command", JSON.stringify(linkConsumersReq))
        .done(function(deploymentsList) {
            populateDeploymentsDropDown(deploymentsList);

            var currentDeploymentName = getUrlParameter('deployment');
            if (currentDeploymentName) {
                // Get all the consumers of this deployment
                // For each deployment, get the link providers
                // when that is done, get the links of the deployment
                // match provider for each link consumer

                // Get all link consumers of this deployment
                var linkConsumersReq = {"command":"curl", "arguments": [{"name": "path", "value": "/link_consumers?deployment="+currentDeploymentName}]};
                $.post("/api/command", JSON.stringify(linkConsumersReq))
                    .done(function(data) {
                        var consumersList = data;

                        // get all the links of this deployment
                        var linksReq = {"command":"curl", "arguments": [{"name": "path", "value": "/links?deployment="+currentDeploymentName}]};
                        $.post("/api/command", JSON.stringify(linksReq))
                            .done(function(data) {
                                var linksList = data;
                                var hashedLinksList = linksList.reduce(function(map, obj) {
                                    map[obj["link_consumer_id"]] = obj;
                                    return map;
                                }, {});

                                var numberOfTotalRequests = deploymentsList.length; // number of total requests
                                var aggregatedProvidersList = [];
                                var consumerProviderLinkList = [];

                                $(deploymentsList).each(function() {
                                    var deploymentName = this.name;
                                    var providersReq = {"command":"curl", "arguments": [{"name": "path", "value": "/link_providers?deployment="+deploymentName}]};
                                    $.post("/api/command", JSON.stringify(providersReq))
                                        .done(function(data) {
                                            $.merge( aggregatedProvidersList, data );
                                            numberOfTotalRequests -= 1;
                                            if(numberOfTotalRequests === 0) {
                                                var hashedProvidersList = aggregatedProvidersList.reduce(function(map, obj) {
                                                    map[obj.id] = obj;
                                                    return map;
                                                }, {});

                                                $.each(consumersList, function (i, value) {
                                                    var tableBlob = {"consumer": value};
                                                    var consumerID = value.id;
                                                    var foundLink = hashedLinksList[consumerID];
                                                    if (foundLink) {
                                                        // there is a link for the consumer
                                                        tableBlob["link"] = foundLink;
                                                        var linkProviderId = foundLink["link_provider_id"];
                                                        if (hashedProvidersList[linkProviderId]) {
                                                            tableBlob["provider"] = hashedProvidersList[linkProviderId];
                                                        }
                                                    }
                                                    consumerProviderLinkList.push(tableBlob);
                                                });

                                                populateLinkConsumersDetailedTable(consumerProviderLinkList);
                                                populateStatistics(consumerProviderLinkList, currentDeploymentName);
                                                $("#linksConsumersDetailedLoadingSpinner").remove();
                                                $("#linksConsumersDetailedContents").removeClass("invisible").addClass("visible");
                                            }
                                        });
                                });
                            });
                    });
            } else {
                $("#linksConsumersDetailedLoadingSpinner").remove();
                $("#linksConsumersDetailedContents").removeClass("invisible").addClass("visible");
            }
        });

    $("#deployment-select").change(function () {
        window.location.href = "/link-consumers-detailed?deployment=" + $("#deployment-select").find(":selected").text();
    });
});

</script>
`
